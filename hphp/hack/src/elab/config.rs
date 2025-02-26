// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

use bitflags::bitflags;
use oxidized::typechecker_options::TypecheckerOptions;

#[derive(Debug, Clone, Default)]
pub struct ProgramSpecificOptions {
    pub is_hhi: bool,
    pub allow_type_constant_in_enum_class: bool,
    pub allow_module_declarations: bool,
}

bitflags! {
    struct Flags: u16 {
        const SOFT_AS_LIKE = 1 << 0;
        const HKT_ENABLED = 1 << 1;
        const IS_HHI = 1 << 2;
        const IS_SYSTEMLIB = 1 << 3;
        const LIKE_TYPE_HINTS_ENABLED = 1 << 4;
        const CONST_ATTRIBUTE = 1 << 5;
        const CONST_STATIC_PROPS = 1 << 6;
        const ALLOW_TYPE_CONSTANT_IN_ENUM_CLASS = 1 << 7;
        const ALLOW_MODULE_DECLARATIONS = 1 << 8;
    }
}

impl Flags {
    pub fn new(tco: &TypecheckerOptions, pso: &ProgramSpecificOptions) -> Self {
        let mut flags: Self = Flags::empty();

        flags.set(
            Self::SOFT_AS_LIKE,
            tco.po_interpret_soft_types_as_like_types,
        );
        flags.set(Self::HKT_ENABLED, tco.tco_higher_kinded_types);
        flags.set(Self::IS_SYSTEMLIB, tco.tco_is_systemlib);
        flags.set(Self::LIKE_TYPE_HINTS_ENABLED, tco.tco_like_type_hints);
        flags.set(Self::CONST_ATTRIBUTE, tco.tco_const_attribute);
        flags.set(Self::CONST_STATIC_PROPS, tco.tco_const_static_props);

        flags.set(Self::IS_HHI, pso.is_hhi);
        flags.set(
            Self::ALLOW_TYPE_CONSTANT_IN_ENUM_CLASS,
            pso.allow_type_constant_in_enum_class,
        );
        flags.set(
            Self::ALLOW_MODULE_DECLARATIONS,
            pso.allow_module_declarations,
        );

        flags
    }
}

#[derive(Debug, Clone)]
pub struct Config {
    flags: Flags,
    pub consistent_ctor_level: isize,
}

impl Default for Config {
    fn default() -> Self {
        Self::new(
            &TypecheckerOptions::default(),
            &ProgramSpecificOptions::default(),
        )
    }
}

impl Config {
    pub fn new(tco: &TypecheckerOptions, pso: &ProgramSpecificOptions) -> Self {
        Self {
            flags: Flags::new(tco, pso),
            consistent_ctor_level: tco.tco_explicit_consistent_constructors,
        }
    }

    pub fn soft_as_like(&self) -> bool {
        self.flags.contains(Flags::SOFT_AS_LIKE)
    }

    pub fn allow_type_constant_in_enum_class(&self) -> bool {
        self.flags
            .contains(Flags::ALLOW_TYPE_CONSTANT_IN_ENUM_CLASS)
    }

    pub fn allow_module_declarations(&self) -> bool {
        self.flags.contains(Flags::ALLOW_MODULE_DECLARATIONS)
    }

    pub fn hkt_enabled(&self) -> bool {
        self.flags.contains(Flags::HKT_ENABLED)
    }

    pub fn is_systemlib(&self) -> bool {
        self.flags.contains(Flags::IS_SYSTEMLIB)
    }

    pub fn like_type_hints_enabled(&self) -> bool {
        self.flags.contains(Flags::LIKE_TYPE_HINTS_ENABLED)
    }

    pub fn is_hhi(&self) -> bool {
        self.flags.contains(Flags::IS_HHI)
    }

    pub fn const_attribute(&self) -> bool {
        self.flags.contains(Flags::CONST_ATTRIBUTE)
    }

    pub fn const_static_props(&self) -> bool {
        self.flags.contains(Flags::CONST_STATIC_PROPS)
    }
}
