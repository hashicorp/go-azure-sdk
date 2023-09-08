package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultDiscoverySource string

const (
	ImportedAppleDeviceIdentityResultDiscoverySourceadminImport             ImportedAppleDeviceIdentityResultDiscoverySource = "AdminImport"
	ImportedAppleDeviceIdentityResultDiscoverySourcedeviceEnrollmentProgram ImportedAppleDeviceIdentityResultDiscoverySource = "DeviceEnrollmentProgram"
	ImportedAppleDeviceIdentityResultDiscoverySourceunknown                 ImportedAppleDeviceIdentityResultDiscoverySource = "Unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityResultDiscoverySource() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultDiscoverySourceadminImport),
		string(ImportedAppleDeviceIdentityResultDiscoverySourcedeviceEnrollmentProgram),
		string(ImportedAppleDeviceIdentityResultDiscoverySourceunknown),
	}
}

func parseImportedAppleDeviceIdentityResultDiscoverySource(input string) (*ImportedAppleDeviceIdentityResultDiscoverySource, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultDiscoverySource{
		"adminimport":             ImportedAppleDeviceIdentityResultDiscoverySourceadminImport,
		"deviceenrollmentprogram": ImportedAppleDeviceIdentityResultDiscoverySourcedeviceEnrollmentProgram,
		"unknown":                 ImportedAppleDeviceIdentityResultDiscoverySourceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultDiscoverySource(input)
	return &out, nil
}
