/**
 * Copyright (c) Meta, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree. An additional
 * directory.
 *
 */
use cxx::CxxString;

#[cxx::bridge(namespace = "HPHP::package")]
mod ffi {
    struct PackageInfo {
        packages: Vec<PackageMapEntry>,
        deployments: Vec<DeploymentMapEntry>,
    }
    struct PackageMapEntry {
        name: String,
        package: Package,
    }
    struct Package {
        uses: Vec<String>,
        includes: Vec<String>,
    }
    struct DeploymentMapEntry {
        name: String,
        deployment: Deployment,
    }
    struct Deployment {
        packages: Vec<String>,
        domain: String,
    }
    extern "Rust" {
        pub fn package_info_cpp_ffi(source_text: &CxxString) -> PackageInfo;
    }
}

pub fn package_info_cpp_ffi(source_text: &CxxString) -> ffi::PackageInfo {
    let info = package::PackageInfo::from_text(&source_text.to_string()).unwrap();
    let convert = |v: Option<&Vec<toml::Spanned<String>>>| {
        v.map(|v| v.iter().map(|v| v.get_ref().clone()).collect())
            .unwrap_or_default()
    };
    let packages = info
        .packages()
        .iter()
        .map(|(name, package)| {
            let package_ffi = ffi::Package {
                uses: convert(package.uses.as_ref()),
                includes: convert(package.includes.as_ref()),
            };
            ffi::PackageMapEntry {
                name: name.to_string(),
                package: package_ffi,
            }
        })
        .collect();
    let deployments = info
        .deployments()
        .map(|deployments_unwrapped| {
            deployments_unwrapped
                .iter()
                .map(|(name, deployment)| {
                    let deployment_ffi = ffi::Deployment {
                        packages: convert(deployment.packages.as_ref()),
                        domain: deployment
                            .domain
                            .as_ref()
                            .map_or("", |domain_unwrapped| domain_unwrapped.get_ref())
                            .to_string(),
                    };
                    ffi::DeploymentMapEntry {
                        name: name.to_string(),
                        deployment: deployment_ffi,
                    }
                })
                .collect()
        })
        .unwrap_or_default();
    ffi::PackageInfo {
        packages,
        deployments,
    }
}
