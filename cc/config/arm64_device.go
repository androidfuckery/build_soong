// Copyright 2015 Google Inc. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
	"fmt"
	"strings"

	"android/soong/android"
)

var (
	arm64Cflags = []string{
		// Help catch common 32/64-bit errors.
		"-Werror=implicit-function-declaration",
	}

	arm64ArchVariantCflags = map[string][]string{
		"armv8-a":            {"-march=armv8-a"},
		"armv8-a-branchprot": {"-march=armv8-a"},
		"armv8-2a":           {"-march=armv8.2-a"},
		"armv8-2a-dotprod":   {"-march=armv8.2-a+dotprod"},
		"armv8-5a":           {"-march=armv8.5-a"},
		"armv8-7a":           {"-march=armv8.7-a"},
		"armv9-a":            {"-march=armv9-a"},
		"armv9-2a":           {"-march=armv9.2-a"},
		"armv9-3a":           {"-march=armv9.3-a"},
		"armv9-4a":           {"-march=armv9.4-a"},
	}

	arm64ArchFeatureCflags = map[string][]string{
		// When Pointer Authentication Codes (PAC) are available, -fstack-protector is unnecessary.
		"branchprot": {
			"-mbranch-protection=standard",
			"-fno-stack-protector",
		},
	}

	arm64Ldflags = []string{
		"-Wl,-z,separate-code",
		"-Wl,-z,separate-loadable-segments",
	}

	arm64Cppflags = []string{}

	arm64CpuVariantCflags = map[string][]string{
		"cortex-a53": []string{
			"-mtune=cortex-a53",
		},
		"cortex-a55": []string{
			"-mtune=cortex-a55",
		},
		"cortex-a72": []string{
			"-mtune=cortex-a72",
		},
		"cortex-a73": []string{
			"-mtune=cortex-a73",
		},
		"cortex-a75": []string{
			"-mtune=cortex-a75",
		},
		"cortex-a76": []string{
			"-mtune=cortex-a76",
		},
		"cortex-a720": []string{
			"-mtune=cortex-a720",
		},
		"kryo": []string{
			"-mtune=kryo",
		},
		"oryon": []string{
			"-mtune=oryon-1",
		},
	}
)

func init() {
	pctx.VariableFunc("Arm64Ldflags", func(ctx android.PackageVarContext) string {
		maxPageSizeFlag := "-Wl,-z,max-page-size=" + ctx.Config().MaxPageSizeSupported()
		flags := append(arm64Ldflags, maxPageSizeFlag)
		return strings.Join(flags, " ")
	})

	pctx.VariableFunc("Arm64Cflags", func(ctx android.PackageVarContext) string {
		flags := arm64Cflags
		if ctx.Config().NoBionicPageSizeMacro() {
			flags = append(flags, "-D__BIONIC_NO_PAGE_SIZE_MACRO")
		} else {
			flags = append(flags, "-D__BIONIC_DEPRECATED_PAGE_SIZE_MACRO")
		}
		return strings.Join(flags, " ")
	})

	pctx.StaticVariable("Arm64Cppflags", strings.Join(arm64Cppflags, " "))

	for variant, cflags := range arm64ArchVariantCflags {
		pctx.StaticVariable("Arm64"+variant+"VariantCflags", strings.Join(cflags, " "))
	}

	pctx.StaticVariable("Arm64CortexA53Cflags", strings.Join(arm64CpuVariantCflags["cortex-a53"], " "))
	pctx.StaticVariable("Arm64CortexA55Cflags", strings.Join(arm64CpuVariantCflags["cortex-a55"], " "))
	pctx.StaticVariable("Arm64CortexA72Cflags", strings.Join(arm64CpuVariantCflags["cortex-a72"], " "))
	pctx.StaticVariable("Arm64CortexA73Cflags", strings.Join(arm64CpuVariantCflags["cortex-a73"], " "))
	pctx.StaticVariable("Arm64CortexA75Cflags", strings.Join(arm64CpuVariantCflags["cortex-a75"], " "))
	pctx.StaticVariable("Arm64CortexA76Cflags", strings.Join(arm64CpuVariantCflags["cortex-a76"], " "))
	pctx.StaticVariable("Arm64CortexA720Cflags", strings.Join(arm64CpuVariantCflags["cortex-a720"], " "))
	pctx.StaticVariable("Arm64KryoCflags", strings.Join(arm64CpuVariantCflags["kryo"], " "))
	pctx.StaticVariable("Arm64OryonCflags", strings.Join(arm64CpuVariantCflags["oryon"], " "))

	pctx.StaticVariable("Arm64FixCortexA53Ldflags", "-Wl,--fix-cortex-a53-843419")
}

var (
	arm64CpuVariantCflagsVar = map[string]string{
		"cortex-a53": "${config.Arm64CortexA53Cflags}",
		"cortex-a55": "${config.Arm64CortexA55Cflags}",
		"cortex-a72": "${config.Arm64CortexA72Cflags}",
		"cortex-a73": "${config.Arm64CortexA73Cflags}",
		"cortex-a75": "${config.Arm64CortexA75Cflags}",
		"cortex-a76": "${config.Arm64CortexA76Cflags}",
		"cortex-a720": "${config.Arm64CortexA720Cflags}",
		"kryo":       "${config.Arm64KryoCflags}",
		"oryon":      "${config.Arm64OryonCflags}",
	}

	arm64CpuVariantLdflags = map[string]string{
		"cortex-a53": "${config.Arm64FixCortexA53Ldflags}",
		"cortex-a72": "${config.Arm64FixCortexA53Ldflags}",
		"cortex-a73": "${config.Arm64FixCortexA53Ldflags}",
		"kryo":       "${config.Arm64FixCortexA53Ldflags}",
	}
)

type toolchainArm64 struct {
	toolchainBionic
	toolchain64Bit

	ldflags         string
	toolchainCflags string
}

func (t *toolchainArm64) Name() string {
	return "arm64"
}

func (t *toolchainArm64) IncludeFlags() string {
	return ""
}

func (t *toolchainArm64) ClangTriple() string {
	return "aarch64-linux-android"
}

func (t *toolchainArm64) Cflags() string {
	return "${config.Arm64Cflags}"
}

func (t *toolchainArm64) Cppflags() string {
	return "${config.Arm64Cppflags}"
}

func (t *toolchainArm64) Ldflags() string {
	return t.ldflags
}

func (t *toolchainArm64) ToolchainCflags() string {
	return t.toolchainCflags
}

func (toolchainArm64) LibclangRuntimeLibraryArch() string {
	return "aarch64"
}

func arm64ToolchainFactory(arch android.Arch) Toolchain {
	// Error now rather than having a confusing Ninja error
	if _, ok := arm64ArchVariantCflags[arch.ArchVariant]; !ok {
		panic(fmt.Sprintf("Unknown ARM64 architecture version: %q", arch.ArchVariant))
	}

	toolchainCflags := []string{"${config.Arm64" + arch.ArchVariant + "VariantCflags}"}
	toolchainCflags = append(toolchainCflags,
		variantOrDefault(arm64CpuVariantCflagsVar, arch.CpuVariant))
	for _, feature := range arch.ArchFeatures {
		toolchainCflags = append(toolchainCflags, arm64ArchFeatureCflags[feature]...)
	}

	extraLdflags := variantOrDefault(arm64CpuVariantLdflags, arch.CpuVariant)
	return &toolchainArm64{
		ldflags: strings.Join([]string{
			"${config.Arm64Ldflags}",
			extraLdflags,
		}, " "),
		toolchainCflags: strings.Join(toolchainCflags, " "),
	}
}

func init() {
	registerToolchainFactory(android.Android, android.Arm64, arm64ToolchainFactory)
}
