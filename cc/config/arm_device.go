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
	"strings"

	"android/soong/android"
)

var (
	armToolchainCflags = []string{
		"-msoft-float",
	}

	armCflags = []string{
		"-fomit-frame-pointer",
		// Revert this after b/322359235 is fixed
		"-mllvm", "-enable-shrink-wrap=false",
	}

	armCppflags = []string{
		// Revert this after b/322359235 is fixed
		"-mllvm", "-enable-shrink-wrap=false",
	}

	armLdflags = []string{
		"-Wl,-m,armelf",
		// Revert this after b/322359235 is fixed
		"-Wl,-mllvm", "-Wl,-enable-shrink-wrap=false",
	}

	armFixCortexA8LdFlags = []string{"-Wl,--fix-cortex-a8"}

	armNoFixCortexA8LdFlags = []string{"-Wl,--no-fix-cortex-a8"}

	armArmCflags = []string{}

	armThumbCflags = []string{
		"-mthumb",
		"-Os",
	}

	armArchVariantCflags = map[string][]string{
		"armv7-a": []string{
			"-march=armv7-a",
			"-mfloat-abi=softfp",
			"-mfpu=vfpv3-d16",
		},
		"armv7-a-neon": []string{
			"-march=armv7-a",
			"-mfloat-abi=softfp",
			"-mfpu=neon",
		},
		"armv8-a": []string{
			"-march=armv8-a",
			"-mfloat-abi=softfp",
			"-mfpu=neon-fp-armv8",
		},
		"armv8-2a": []string{
			"-march=armv8.2-a",
			"-mfloat-abi=softfp",
			"-mfpu=neon-fp-armv8",
		},
	}

	armCpuVariantCflags = map[string][]string{
		"cortex-a7": []string{
			"-mtune=cortex-a7",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a8": []string{
			"-mtune=cortex-a8",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a9": []string{
			"-mtune=cortex-a9",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a15": []string{
			"-mtune=cortex-a15",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a32": []string{
			"-mtune=cortex-a32",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a53": []string{
			"-mtune=cortex-a53",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a55": []string{
			"-mtune=cortex-a55",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a72": []string{
			"-mtune=cortex-a72",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a73": []string{
			"-mtune=cortex-a73",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a75": []string{
			"-mtune=cortex-a75",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"cortex-a76": []string{
			"-mtune=cortex-a76",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"krait": []string{
			"-mtune=krait",
			"-mfpu=neon-vfpv4",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
		"kryo": []string{
			"-mtune=kryo",
			"-mfpu=neon-fp-armv8",
			// Fake an ARM compiler flag as these processors support LPAE which clang
			// don't advertise.
			// TODO This is a hack and we need to add it for each processor that supports LPAE until some
			// better solution comes around. See Bug 27340895
			"-D__ARM_FEATURE_LPAE=1",
		},
	}
)

const (
	name        = "arm"
	ndkTriple   = "arm-linux-androideabi"
	clangTriple = "armv7a-linux-androideabi"
)

func init() {
	pctx.StaticVariable("ArmLdflags", strings.Join(armLdflags, " "))

	pctx.StaticVariable("ArmFixCortexA8LdFlags", strings.Join(armFixCortexA8LdFlags, " "))
	pctx.StaticVariable("ArmNoFixCortexA8LdFlags", strings.Join(armNoFixCortexA8LdFlags, " "))

	// Clang cflags
	pctx.StaticVariable("ArmToolchainCflags", strings.Join(armToolchainCflags, " "))
	pctx.StaticVariable("ArmCflags", strings.Join(armCflags, " "))
	pctx.StaticVariable("ArmCppflags", strings.Join(armCppflags, " "))

	// Clang ARM vs. Thumb instruction set cflags
	pctx.StaticVariable("ArmArmCflags", strings.Join(armArmCflags, " "))
	pctx.StaticVariable("ArmThumbCflags", strings.Join(armThumbCflags, " "))

	// Clang arch variant cflags
	pctx.StaticVariable("ArmArmv7ACflags", strings.Join(armArchVariantCflags["armv7-a"], " "))
	pctx.StaticVariable("ArmArmv7ANeonCflags", strings.Join(armArchVariantCflags["armv7-a-neon"], " "))
	pctx.StaticVariable("ArmArmv8ACflags", strings.Join(armArchVariantCflags["armv8-a"], " "))
	pctx.StaticVariable("ArmArmv82ACflags", strings.Join(armArchVariantCflags["armv8-2a"], " "))

	// Clang cpu variant cflags
	pctx.StaticVariable("ArmGenericCflags", strings.Join(armCpuVariantCflags[""], " "))
	pctx.StaticVariable("ArmCortexA7Cflags", strings.Join(armCpuVariantCflags["cortex-a7"], " "))
	pctx.StaticVariable("ArmCortexA8Cflags", strings.Join(armCpuVariantCflags["cortex-a8"], " "))
	pctx.StaticVariable("ArmCortexA9Cflags", strings.Join(armCpuVariantCflags["cortex-a9"], " "))
	pctx.StaticVariable("ArmCortexA15Cflags", strings.Join(armCpuVariantCflags["cortex-a15"], " "))
	pctx.StaticVariable("ArmCortexA32Cflags", strings.Join(armCpuVariantCflags["cortex-a32"], " "))
	pctx.StaticVariable("ArmCortexA53Cflags", strings.Join(armCpuVariantCflags["cortex-a53"], " "))
	pctx.StaticVariable("ArmCortexA55Cflags", strings.Join(armCpuVariantCflags["cortex-a55"], " "))
	pctx.StaticVariable("ArmCortexA72Cflags", strings.Join(armCpuVariantCflags["cortex-a72"], " "))
	pctx.StaticVariable("ArmCortexA73Cflags", strings.Join(armCpuVariantCflags["cortex-a73"], " "))
	pctx.StaticVariable("ArmCortexA75Cflags", strings.Join(armCpuVariantCflags["cortex-a75"], " "))
	pctx.StaticVariable("ArmCortexA76Cflags", strings.Join(armCpuVariantCflags["cortex-a76"], " "))
	pctx.StaticVariable("ArmKraitCflags", strings.Join(armCpuVariantCflags["krait"], " "))
	pctx.StaticVariable("ArmKryoCflags", strings.Join(armCpuVariantCflags["kryo"], " "))
}

var (
	armArchVariantCflagsVar = map[string]string{
		"armv7-a":      "${config.ArmArmv7ACflags}",
		"armv7-a-neon": "${config.ArmArmv7ANeonCflags}",
		"armv8-a":      "${config.ArmArmv8ACflags}",
		"armv8-2a":     "${config.ArmArmv82ACflags}",
	}

	armCpuVariantCflagsVar = map[string]string{
		"":               "${config.ArmGenericCflags}",
		"cortex-a7":      "${config.ArmCortexA7Cflags}",
		"cortex-a8":      "${config.ArmCortexA8Cflags}",
		"cortex-a9":      "${config.ArmCortexA9Cflags}",
		"cortex-a15":     "${config.ArmCortexA15Cflags}",
		"cortex-a32":     "${config.ArmCortexA32Cflags}",
		"cortex-a53":     "${config.ArmCortexA53Cflags}",
		"cortex-a55":     "${config.ArmCortexA55Cflags}",
		"cortex-a72":     "${config.ArmCortexA72Cflags}",
		"cortex-a73":     "${config.ArmCortexA73Cflags}",
		"cortex-a75":     "${config.ArmCortexA75Cflags}",
		"cortex-a76":     "${config.ArmCortexA76Cflags}",
		"krait":          "${config.ArmKraitCflags}",
		"kryo":           "${config.ArmKryoCflags}",
	}
)

type toolchainArm struct {
	toolchainBionic
	toolchain32Bit
	ldflags         string
	toolchainCflags string
}

func (t *toolchainArm) Name() string {
	return name
}

func (t *toolchainArm) IncludeFlags() string {
	return ""
}

func (t *toolchainArm) ClangTriple() string {
	// http://b/72619014 work around llvm LTO bug.
	return clangTriple
}

func (t *toolchainArm) ndkTriple() string {
	// Use current NDK include path, while ClangTriple is changed.
	return ndkTriple
}

func (t *toolchainArm) ToolchainCflags() string {
	return t.toolchainCflags
}

func (t *toolchainArm) Cflags() string {
	return "${config.ArmCflags}"
}

func (t *toolchainArm) Cppflags() string {
	return "${config.ArmCppflags}"
}

func (t *toolchainArm) Ldflags() string {
	return t.ldflags // TODO: handle V8 cases
}

func (t *toolchainArm) InstructionSetFlags(isa string) (string, error) {
	switch isa {
	case "arm":
		return "${config.ArmArmCflags}", nil
	case "thumb", "":
		return "${config.ArmThumbCflags}", nil
	default:
		return t.toolchainBase.InstructionSetFlags(isa)
	}
}

func (toolchainArm) LibclangRuntimeLibraryArch() string {
	return name
}

func armToolchainFactory(arch android.Arch) Toolchain {
	toolchainCflags := make([]string, 2, 3)

	toolchainCflags[0] = "${config.ArmToolchainCflags}"
	toolchainCflags[1] = armArchVariantCflagsVar[arch.ArchVariant]

	toolchainCflags = append(toolchainCflags,
		variantOrDefault(armCpuVariantCflagsVar, arch.CpuVariant))

	return &toolchainArm{
		ldflags:         "${config.ArmLdflags}",
		toolchainCflags: strings.Join(toolchainCflags, " "),
	}
}

func init() {
	registerToolchainFactory(android.Android, android.Arm, armToolchainFactory)
}
