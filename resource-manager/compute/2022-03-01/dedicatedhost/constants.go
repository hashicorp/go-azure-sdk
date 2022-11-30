package dedicatedhost

import "strings"

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

func parseDedicatedHostLicenseTypes(input string) (*DedicatedHostLicenseTypes, error) {
	vals := map[string]DedicatedHostLicenseTypes{
		"none":                     DedicatedHostLicenseTypesNone,
		"windows_server_hybrid":    DedicatedHostLicenseTypesWindowsServerHybrid,
		"windows_server_perpetual": DedicatedHostLicenseTypesWindowsServerPerpetual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DedicatedHostLicenseTypes(input)
	return &out, nil
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

func parseStatusLevelTypes(input string) (*StatusLevelTypes, error) {
	vals := map[string]StatusLevelTypes{
		"error":   StatusLevelTypesError,
		"info":    StatusLevelTypesInfo,
		"warning": StatusLevelTypesWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusLevelTypes(input)
	return &out, nil
}
