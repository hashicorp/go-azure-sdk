package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityDiscoverySource string

const (
	ImportedAppleDeviceIdentityDiscoverySourceadminImport             ImportedAppleDeviceIdentityDiscoverySource = "AdminImport"
	ImportedAppleDeviceIdentityDiscoverySourcedeviceEnrollmentProgram ImportedAppleDeviceIdentityDiscoverySource = "DeviceEnrollmentProgram"
	ImportedAppleDeviceIdentityDiscoverySourceunknown                 ImportedAppleDeviceIdentityDiscoverySource = "Unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityDiscoverySource() []string {
	return []string{
		string(ImportedAppleDeviceIdentityDiscoverySourceadminImport),
		string(ImportedAppleDeviceIdentityDiscoverySourcedeviceEnrollmentProgram),
		string(ImportedAppleDeviceIdentityDiscoverySourceunknown),
	}
}

func parseImportedAppleDeviceIdentityDiscoverySource(input string) (*ImportedAppleDeviceIdentityDiscoverySource, error) {
	vals := map[string]ImportedAppleDeviceIdentityDiscoverySource{
		"adminimport":             ImportedAppleDeviceIdentityDiscoverySourceadminImport,
		"deviceenrollmentprogram": ImportedAppleDeviceIdentityDiscoverySourcedeviceEnrollmentProgram,
		"unknown":                 ImportedAppleDeviceIdentityDiscoverySourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityDiscoverySource(input)
	return &out, nil
}
