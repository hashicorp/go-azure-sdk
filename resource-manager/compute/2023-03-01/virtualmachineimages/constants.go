package virtualmachineimages

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlternativeType string

const (
	AlternativeTypeNone  AlternativeType = "None"
	AlternativeTypeOffer AlternativeType = "Offer"
	AlternativeTypePlan  AlternativeType = "Plan"
)

func PossibleValuesForAlternativeType() []string {
	return []string{
		string(AlternativeTypeNone),
		string(AlternativeTypeOffer),
		string(AlternativeTypePlan),
	}
}

func parseAlternativeType(input string) (*AlternativeType, error) {
	vals := map[string]AlternativeType{
		"none":  AlternativeTypeNone,
		"offer": AlternativeTypeOffer,
		"plan":  AlternativeTypePlan,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlternativeType(input)
	return &out, nil
}

type ArchitectureTypes string

const (
	ArchitectureTypesArmSixFour ArchitectureTypes = "Arm64"
	ArchitectureTypesXSixFour   ArchitectureTypes = "x64"
)

func PossibleValuesForArchitectureTypes() []string {
	return []string{
		string(ArchitectureTypesArmSixFour),
		string(ArchitectureTypesXSixFour),
	}
}

func parseArchitectureTypes(input string) (*ArchitectureTypes, error) {
	vals := map[string]ArchitectureTypes{
		"arm64": ArchitectureTypesArmSixFour,
		"x64":   ArchitectureTypesXSixFour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ArchitectureTypes(input)
	return &out, nil
}

type HyperVGenerationTypes string

const (
	HyperVGenerationTypesVOne HyperVGenerationTypes = "V1"
	HyperVGenerationTypesVTwo HyperVGenerationTypes = "V2"
)

func PossibleValuesForHyperVGenerationTypes() []string {
	return []string{
		string(HyperVGenerationTypesVOne),
		string(HyperVGenerationTypesVTwo),
	}
}

func parseHyperVGenerationTypes(input string) (*HyperVGenerationTypes, error) {
	vals := map[string]HyperVGenerationTypes{
		"v1": HyperVGenerationTypesVOne,
		"v2": HyperVGenerationTypesVTwo,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HyperVGenerationTypes(input)
	return &out, nil
}

type ImageState string

const (
	ImageStateActive                  ImageState = "Active"
	ImageStateDeprecated              ImageState = "Deprecated"
	ImageStateScheduledForDeprecation ImageState = "ScheduledForDeprecation"
)

func PossibleValuesForImageState() []string {
	return []string{
		string(ImageStateActive),
		string(ImageStateDeprecated),
		string(ImageStateScheduledForDeprecation),
	}
}

func parseImageState(input string) (*ImageState, error) {
	vals := map[string]ImageState{
		"active":                  ImageStateActive,
		"deprecated":              ImageStateDeprecated,
		"scheduledfordeprecation": ImageStateScheduledForDeprecation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImageState(input)
	return &out, nil
}

type OperatingSystemTypes string

const (
	OperatingSystemTypesLinux   OperatingSystemTypes = "Linux"
	OperatingSystemTypesWindows OperatingSystemTypes = "Windows"
)

func PossibleValuesForOperatingSystemTypes() []string {
	return []string{
		string(OperatingSystemTypesLinux),
		string(OperatingSystemTypesWindows),
	}
}

func parseOperatingSystemTypes(input string) (*OperatingSystemTypes, error) {
	vals := map[string]OperatingSystemTypes{
		"linux":   OperatingSystemTypesLinux,
		"windows": OperatingSystemTypesWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperatingSystemTypes(input)
	return &out, nil
}

type VMDiskTypes string

const (
	VMDiskTypesNone      VMDiskTypes = "None"
	VMDiskTypesUnmanaged VMDiskTypes = "Unmanaged"
)

func PossibleValuesForVMDiskTypes() []string {
	return []string{
		string(VMDiskTypesNone),
		string(VMDiskTypesUnmanaged),
	}
}

func parseVMDiskTypes(input string) (*VMDiskTypes, error) {
	vals := map[string]VMDiskTypes{
		"none":      VMDiskTypesNone,
		"unmanaged": VMDiskTypesUnmanaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VMDiskTypes(input)
	return &out, nil
}
