package virtualmachineimages

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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
