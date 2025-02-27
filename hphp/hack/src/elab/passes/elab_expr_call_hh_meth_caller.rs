// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.
// Copyright (c) Meta Platforms, Inc. and affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.
use std::mem::take;
use std::ops::ControlFlow;

use naming_special_names_rust as sn;
use oxidized::aast_defs::ClassId;
use oxidized::aast_defs::ClassId_;
use oxidized::aast_defs::Expr;
use oxidized::aast_defs::Expr_;
use oxidized::ast_defs::Id;
use oxidized::ast_defs::ParamKind;
use oxidized::naming_error::NamingError;
use oxidized::naming_phase_error::NamingPhaseError;

use crate::config::Config;
use crate::elab_utils;
use crate::Pass;

#[derive(Clone, Copy, Default)]
pub struct ElabExprCallHhMethCallerPass;

impl Pass for ElabExprCallHhMethCallerPass {
    fn on_ty_expr__bottom_up<Ex: Default, En>(
        &mut self,
        elem: &mut Expr_<Ex, En>,
        _cfg: &Config,
        errs: &mut Vec<NamingPhaseError>,
    ) -> ControlFlow<(), ()> {
        let invalid = |expr_: &mut Expr_<_, _>| {
            let inner_expr_ = std::mem::replace(expr_, Expr_::Null);
            let inner_expr = elab_utils::expr::from_expr_(inner_expr_);
            *expr_ = Expr_::Invalid(Box::new(Some(inner_expr)));
            ControlFlow::Break(())
        };

        match elem {
            Expr_::Call(box (
                Expr(_, fn_expr_pos, Expr_::Id(box id)),
                _,
                fn_param_exprs,
                fn_variadic_param_opt,
            )) if id.name() == sn::autoimported_functions::METH_CALLER => {
                // Raise an error if we have a variadic arg
                if let Some(Expr(_, pos, _)) = fn_variadic_param_opt {
                    errs.push(NamingPhaseError::Naming(NamingError::TooFewArguments(
                        pos.clone(),
                    )))
                }

                match fn_param_exprs.as_mut_slice() {
                    [_, _, _, ..] => {
                        errs.push(NamingPhaseError::Naming(NamingError::TooManyArguments(
                            fn_expr_pos.clone(),
                        )));
                        invalid(elem)
                    }
                    [
                        (ParamKind::Pnormal, Expr(_, rcvr_pos, rcvr)),
                        (ParamKind::Pnormal, Expr(_, meth_pos, Expr_::String(meth))),
                    ] => match rcvr {
                        Expr_::String(rcvr) => {
                            *elem = Expr_::MethodCaller(Box::new((
                                Id(take(rcvr_pos), rcvr.to_string()),
                                (take(meth_pos), meth.to_string()),
                            )));
                            ControlFlow::Continue(())
                        }
                        Expr_::ClassConst(box (ClassId(_, _, ClassId_::CI(id)), (_, mem)))
                            if mem == sn::members::M_CLASS =>
                        {
                            *elem = Expr_::MethodCaller(Box::new((
                                take(id),
                                (take(meth_pos), meth.to_string()),
                            )));
                            ControlFlow::Continue(())
                        }
                        _ => {
                            errs.push(NamingPhaseError::Naming(NamingError::IllegalMethCaller(
                                fn_expr_pos.clone(),
                            )));
                            invalid(elem)
                        }
                    },

                    // We expect a string literal as the second argument and neither param
                    // can be inout; raise an error and invalidate
                    [_, _] => {
                        errs.push(NamingPhaseError::Naming(NamingError::IllegalMethCaller(
                            fn_expr_pos.clone(),
                        )));
                        invalid(elem)
                    }

                    // We are expecting exactly two args
                    [] | [_] => {
                        errs.push(NamingPhaseError::Naming(NamingError::TooFewArguments(
                            fn_expr_pos.clone(),
                        )));
                        invalid(elem)
                    }
                }
            }
            _ => ControlFlow::Continue(()),
        }
    }
}

#[cfg(test)]
mod tests {

    use super::*;
    use crate::elab_utils;
    use crate::Transform;

    // -- Valid cases resulting in elaboration to `MethodCaller` ---------------

    #[test]
    fn test_valid_two_string_args() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(rcvr_name.into())),
                ),
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
            ],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        // Expect no errors
        assert!(errs.is_empty());

        // Expect our `Expr_` to elaborate to a `MethodCaller`
        assert!(match elem {
            Expr_::MethodCaller(meth_caller) => {
                let (Id(_, x), (_, y)) = *meth_caller;
                x == rcvr_name && y == meth_name
            }
            _ => false,
        })
    }

    #[test]
    fn test_valid_class_const_string_args() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr(
                        (),
                        elab_utils::pos::null(),
                        Expr_::ClassConst(Box::new((
                            ClassId(
                                (),
                                elab_utils::pos::null(),
                                ClassId_::CI(Id(elab_utils::pos::null(), rcvr_name.into())),
                            ),
                            (elab_utils::pos::null(), sn::members::M_CLASS.to_string()),
                        ))),
                    ),
                ),
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
            ],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        // Expect no errors
        assert!(errs.is_empty());

        // Expect our `Expr_` to elaborate to a `MethodCaller`
        assert!(match elem {
            Expr_::MethodCaller(meth_caller) => {
                let (Id(_, x), (_, y)) = *meth_caller;
                x == rcvr_name && y == meth_name
            }
            _ => false,
        })
    }

    #[test]
    fn test_valid_with_variadic_arg() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(rcvr_name.into())),
                ),
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
            ],
            Some(elab_utils::expr::null()),
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        // Expect `TooFewArgs` error from variadic param
        let too_few_args_err_opt = errs.pop();
        assert!(matches!(
            too_few_args_err_opt,
            Some(NamingPhaseError::Naming(NamingError::TooFewArguments(_)))
        ));

        // Expect our `Expr_` to elaborate to a `MethodCaller`
        assert!(match elem {
            Expr_::MethodCaller(meth_caller) => {
                let (Id(_, x), (_, y)) = *meth_caller;
                x == rcvr_name && y == meth_name
            }
            _ => false,
        })
    }
    // -- Invalid cases resulting in elaboration to `Invalid(Some(Call(..))` ---

    #[test]
    fn test_invalid_arg_type() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::Null),
                ),
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
            ],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        let illegal_err_opt = errs.pop();
        assert!(matches!(
            illegal_err_opt,
            Some(NamingPhaseError::Naming(NamingError::IllegalMethCaller(..)))
        ));

        // Expect our original expression to be wrapped in `Invalid`
        assert!(match elem {
            Expr_::Invalid(expr) => {
                if let Some(Expr(_, _, Expr_::Call(cc))) = *expr {
                    if let (Expr(_, _, Expr_::Id(id)), _, _, _) = *cc {
                        let Id(_, nm) = *id;
                        nm == sn::autoimported_functions::METH_CALLER
                    } else {
                        false
                    }
                } else {
                    false
                }
            }
            _ => false,
        })
    }

    #[test]
    fn test_invalid_param_kind() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(rcvr_name.into())),
                ),
                (
                    ParamKind::Pinout(elab_utils::pos::null()),
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
            ],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        let illegal_err_opt = errs.pop();
        assert!(matches!(
            illegal_err_opt,
            Some(NamingPhaseError::Naming(NamingError::IllegalMethCaller(..)))
        ));

        // Expect our original expression to be wrapped in `Invalid`
        assert!(match elem {
            Expr_::Invalid(expr) => {
                if let Some(Expr(_, _, Expr_::Call(cc))) = *expr {
                    if let (Expr(_, _, Expr_::Id(id)), _, _, _) = *cc {
                        let Id(_, nm) = *id;
                        nm == sn::autoimported_functions::METH_CALLER
                    } else {
                        false
                    }
                } else {
                    false
                }
            }
            _ => false,
        })
    }

    #[test]
    fn test_too_few_args() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![(
                ParamKind::Pnormal,
                Expr((), elab_utils::pos::null(), Expr_::String(rcvr_name.into())),
            )],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        let too_few_args_err_opt = errs.pop();
        assert!(matches!(
            too_few_args_err_opt,
            Some(NamingPhaseError::Naming(NamingError::TooFewArguments(_)))
        ));

        // Expect our original expression to be wrapped in `Invalid`
        assert!(match elem {
            Expr_::Invalid(expr) => {
                if let Some(Expr(_, _, Expr_::Call(cc))) = *expr {
                    if let (Expr(_, _, Expr_::Id(id)), _, _, _) = *cc {
                        let Id(_, nm) = *id;
                        nm == sn::autoimported_functions::METH_CALLER
                    } else {
                        false
                    }
                } else {
                    false
                }
            }
            _ => false,
        })
    }

    #[test]
    fn test_too_many_args() {
        let cfg = Config::default();
        let mut errs = Vec::default();
        let mut pass = ElabExprCallHhMethCallerPass;
        let rcvr_name = "wut";
        let meth_name = "foo";
        let mut elem: Expr_<(), ()> = Expr_::Call(Box::new((
            Expr(
                (),
                elab_utils::pos::null(),
                Expr_::Id(Box::new(Id(
                    elab_utils::pos::null(),
                    sn::autoimported_functions::METH_CALLER.to_string(),
                ))),
            ),
            vec![],
            vec![
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(rcvr_name.into())),
                ),
                (
                    ParamKind::Pnormal,
                    Expr((), elab_utils::pos::null(), Expr_::String(meth_name.into())),
                ),
                (ParamKind::Pnormal, elab_utils::expr::null()),
            ],
            None,
        )));
        elem.transform(&cfg, &mut errs, &mut pass);

        let too_many_args_err_opt = errs.pop();
        assert!(matches!(
            too_many_args_err_opt,
            Some(NamingPhaseError::Naming(NamingError::TooManyArguments(_)))
        ));

        // Expect our original expression to be wrapped in `Invalid`
        assert!(match elem {
            Expr_::Invalid(expr) => {
                if let Some(Expr(_, _, Expr_::Call(cc))) = *expr {
                    if let (Expr(_, _, Expr_::Id(id)), _, _, _) = *cc {
                        let Id(_, nm) = *id;
                        nm == sn::autoimported_functions::METH_CALLER
                    } else {
                        false
                    }
                } else {
                    false
                }
            }
            _ => false,
        })
    }
}
