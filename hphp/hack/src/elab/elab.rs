// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

#![feature(box_patterns)]
#![feature(is_some_and)]
#![feature(let_chains)]

/// Used to combine multiple types implementing `Pass` into nested `Passes` types
/// without requiring them to hand write it so :
/// `passes![p1, p2, p3]` => `Passes(p1, Passes(p2, p3))`
macro_rules! passes {
    ( $p:expr $(,$ps:expr)+ $(,)? ) => {
        $crate::pass::Passes { fst: $p, snd: passes!($($ps),*) }
    };
    ( $p:expr $(,)? ) => {
        $p
    };
}

pub mod config;
mod elab_utils;
mod pass;
mod passes;
mod transform;

use config::ProgramSpecificOptions;
use oxidized::ast;
use oxidized::naming_phase_error::NamingPhaseError;
use oxidized::typechecker_options::TypecheckerOptions;
use pass::Pass;
use transform::Transform;

pub fn elaborate_program(
    program: &mut ast::Program,
    tco: &TypecheckerOptions,
) -> Vec<NamingPhaseError> {
    let pos = program.first_pos();
    let filename = pos.map(|pos| pos.filename().path());
    let dir = filename.and_then(|file| file.parent());

    let is_hhi = pos.map_or(false, |pos| {
        pos.filename().prefix() == relative_path::Prefix::Hhi
    });

    let allow_type_constant_in_enum_class = tco
        .tco_allow_all_locations_for_type_constant_in_enum_class
        || dir.map_or(false, |dir| {
            tco.tco_allowed_locations_for_type_constant_in_enum_class
                .iter()
                .any(|prefix| dir.starts_with(prefix))
        });

    let allow_module_declarations = tco.tco_allow_all_files_for_module_declarations
        || tco
            .tco_allowed_files_for_module_declarations
            .iter()
            .any(|spec| {
                !spec.is_empty()
                    && filename.map_or(false, |filename| {
                        spec.ends_with('*') && filename.starts_with(&spec[..spec.len() - 1])
                            || filename.to_str().unwrap() == spec
                    })
            });

    elaborate(
        program,
        tco,
        &ProgramSpecificOptions {
            is_hhi,
            allow_type_constant_in_enum_class,
            allow_module_declarations,
        },
    )
}

fn elaborate<T: Transform>(
    node: &mut T,
    tco: &TypecheckerOptions,
    pso: &ProgramSpecificOptions,
) -> Vec<NamingPhaseError> {
    let cfg = config::Config::new(tco, pso);

    #[rustfmt::skip]
    let mut passes = passes![
        // Stop on `Invalid` expressions
        passes::guard_invalid::GuardInvalidPass::default(),

        // -- Canonicalization passes -----------------------------------------

        // Remove top-level file attributes, noop and markup statements
        passes::elab_defs::ElabDefsPass::default(),

        // Remove function bodies when in hhi mode
        passes::elab_func_body::ElabFuncBodyPass::default(),

        // Flatten `Block` statements
        passes::elab_block::ElabBlockPass::default(),

        // Strip `Hsoft` hints or replace with `Hlike`
        passes::elab_hint_hsoft::ElabHintHsoftPass::default(),

        // Elaborate `Happly` to canonical representation, if any
        passes::elab_hint_happly::ElabHintHapplyPass::default(),

        // Elaborate class identifier expressions (`CIexpr`) to canonical
        // representation: `CIparent`, `CIself`, `CIstatic`, `CI` _or_
        // `CIexpr (_,_, Lvar _ | This )`
        passes::elab_class_id::ElabClassIdPass::default(),

        // Strip type parameters from type parameters when HKTs are not enabled
        passes::elab_hkt::ElabHktPass::default(),

        // Elaborate `Collection` to `ValCollection` or `KeyValCollection`
        passes::elab_expr_collection::ElabExprCollectionPass::default(),

        // Deduplicate user attributes
        passes::elab_user_attributes::ElabUserAttributesPass::default(),

        // Replace import expressions with invalid expression marker
        passes::elab_expr_import::ElabExprImportPass::default(),

        // Elaborate local variables to canonical representation
        passes::elab_expr_lvar::ElabExprLvarPass::default(),

        // Warn of explicit use of builtin enum classes; make subtyping of
        // enum classes explicit
        passes::elab_enum_class::ElabEnumClassPass::default(),

        // Elaborate class members & xhp attributes
        passes::elab_class_vars::ElabClassVarsPass::default(),

        // Elaborate special function calls to canonical representation, if any
        passes::validate_expr_call_echo::ValidateExprCallEchoPass::default(),
        passes::elab_expr_call_call_user_func::ElabExprCallCallUserFuncPass::default(),
        passes::elab_expr_call_hh_meth_caller::ElabExprCallHhMethCallerPass::default(),

        // Elaborate invariant calls to canonical representation
        passes::elab_expr_call_hh_invariant::ElabExprCallHhInvariantPass::default(),

        // -- Mark invalid hints and expressions & miscellaneous validation ---

        // Replace invalid uses of `void` and `noreturn` with `Herr`
        passes::elab_hint_retonly::ElabHintRetonlyPass::default(),

        // Replace invalid uses of wildcard hints with `Herr`
        passes::elab_hint_wildcard::ElabHintWildcardPass::default(),

        // Replace uses to `self` in shape field names with referenced class
        passes::elab_shape_field_name::ElabShapeFieldNamePass::default(),

        // Replace invalid uses of `this` hints with `Herr`
        passes::elab_hint_this::ElabHintThisPass::default(),

        // Replace invalid `Haccess` root hints with `Herr`
        passes::elab_hint_haccess::ElabHintHaccessPass::default(),

        // Replace empty `Tuple`s with invalid expression marker
        passes::elab_expr_tuple::ElabExprTuplePass::default(),

        // Validate / replace invalid uses of dynamic classes in `New` and `Class_get`
        // expressions
        passes::elab_dynamic_class_name::ElabDynamicClassNamePass::default(),

        // Replace non-constant class or global constant with invalid expression marker
        passes::elab_const_expr::ElabConstExprPass::default(),

        // Replace malformed key / value bindings in as expressions with invalid
        // local var markers
        passes::elab_as_expr::ElabAsExprPass::default(),

        // Validate hints used in `Cast` expressions
        passes::validate_expr_cast::ValidateExprCastPass::default(),

        // Check for duplicate function parameter names
        passes::validate_fun_params::ValidateFunParamsPass::default(),

        // Validate use of `require implements`, `require extends` and
        // `require class` declarations for traits, interfaces and classes
        passes::validate_class_req::ValidateClassReqPass::default(),

        // Validation dealing with common xhp naming errors
        passes::validate_xhp_name::ValidateXhpNamePass::default(),

        // -- Elaboration & validation under typechecker options --------------

        // Add `supportdyn` and `Like` wrappers everywhere - under `everything-sdt`
        // typechecker option
        // passes::elab_everything_sdt::ElabEverythingSdtPass::default(),

        // Validate use of `Hlike` hints - depends on `enable-like-type-hints`
        // and `everything_sdt` typechecker options
        // passes::validate_like_hint::ValidateLikeHintPass::default(),

        // Validate constructors under
        // `consistent-explicit_consistent_constructors` typechecker option
        passes::validate_class_consistent_construct::ValidateClassConsistentConstructPass::default(),

        // Validate  use of `SupportDyn` class - depends on `enable-supportdyn`
        // and `everything_sdt` typechecker options
        // passes::validate_supportdyn::ValidateSupportdynPass::default(),

        // Validate uses of enum class type constants - depends on:
        // - `allow_all_locations_for_type_constant_in_enum_class`
        // - `allowed_locations_for_type_constant_in_enum_class`
        // typecheck options
        passes::validate_enum_class_typeconst::ValidateEnumClassTypeconstPass::default(),

        // Validate use of module definitions - depends on:
        // - `allow_all_files_for_module_declarations`
        // - `allowed_files_for_module_declarations`
        // typechecker options
        passes::validate_module::ValidateModulePass::default(),

        // -- Old 'NAST checks' ------------------------------------------------

        // Validate use of the `__Const` attribute on classes - depends on
        // `const_attribute` typechecker option
        passes::validate_class_user_attribute_const::ValidateClassUserAttributeConstPass::default(),

        // Validate use of the `__Const` attribute on static class vars - depends
        // on the `const_static_props` typechecker option
        passes::validate_class_var_user_attribute_const::ValidateClassVarUserAttributeConstPass::default(),
        passes::validate_class_var_user_attribute_lsb::ValidateClassVarUserAttributeLsbPass::default(),

        // Validate `inout` `FunParam`s ensuring they are not used in functions with
        // special semantics or in memoized functions
        passes::validate_fun_param_inout::ValidateFunParamInoutPass::default(),

        // Validate use of `Await` in sync functions and return in generators
        passes::validate_coroutine::ValidateCoroutinePass::default(),

        // Checks for the presence of a function body in methods, use of traits
        // and instance and static member variables in an interface definition
        passes::validate_interface::ValidateInterfacePass::default(),


        // Checks for use of reserved names in functions, methods, class identifiers
        // and class constants
        passes::validate_illegal_name::ValidateIllegalNamePass::default(),


        passes::validate_control_context::ValidateControlContextPass::default(),

        passes::validate_class_tparams::ValidateClassTparamsPass::default(),

    ];

    let mut errs = Vec::default();
    node.transform(&cfg, &mut errs, &mut passes);
    errs
}
