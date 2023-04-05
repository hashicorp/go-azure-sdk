package dedicatedhosts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedicatedHostLicenseTypes string

const (
	DedicatedHostLicenseTypesNone                   DedicatedHostLicenseTypes = "None"
	DedicatedHostLicenseTypesWindowsServerHybrid    DedicatedHostLicenseTypes = "Windows_Server_Hybrid"
	DedicatedHostLicenseTypesWindowsServerPerpetual DedicatedHostLicenseTypes = "Windows_Server_Perpetual"
)

func PossibleValuesForDedicatedHostLicenseTypes() []string {
	return []string{
		string(DedicatedHostLicenseTypesNone),
		string(DedicatedHostLicenseTypesWindowsServerHybrid),
		string(DedicatedHostLicenseTypesWindowsServerPerpetual),
	}
}

type InstanceViewTypes string

const (
	InstanceViewTypesInstanceView InstanceViewTypes = "instanceView"
	InstanceViewTypesUserData     InstanceViewTypes = "userData"
)

func PossibleValuesForInstanceViewTypes() []string {
	return []string{
		string(InstanceViewTypesInstanceView),
		string(InstanceViewTypesUserData),
	}
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}
