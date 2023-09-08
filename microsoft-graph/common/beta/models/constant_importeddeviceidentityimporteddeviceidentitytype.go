package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityImportedDeviceIdentityTypeimei         ImportedDeviceIdentityImportedDeviceIdentityType = "Imei"
	ImportedDeviceIdentityImportedDeviceIdentityTypeserialNumber ImportedDeviceIdentityImportedDeviceIdentityType = "SerialNumber"
	ImportedDeviceIdentityImportedDeviceIdentityTypeunknown      ImportedDeviceIdentityImportedDeviceIdentityType = "Unknown"
)

func PossibleValuesForImportedDeviceIdentityImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeimei),
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeserialNumber),
		string(ImportedDeviceIdentityImportedDeviceIdentityTypeunknown),
	}
}

func parseImportedDeviceIdentityImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityImportedDeviceIdentityType{
		"imei":         ImportedDeviceIdentityImportedDeviceIdentityTypeimei,
		"serialnumber": ImportedDeviceIdentityImportedDeviceIdentityTypeserialNumber,
		"unknown":      ImportedDeviceIdentityImportedDeviceIdentityTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityImportedDeviceIdentityType(input)
	return &out, nil
}
