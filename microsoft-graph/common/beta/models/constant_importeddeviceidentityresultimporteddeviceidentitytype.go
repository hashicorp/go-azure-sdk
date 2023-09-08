package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityResultImportedDeviceIdentityTypeimei         ImportedDeviceIdentityResultImportedDeviceIdentityType = "Imei"
	ImportedDeviceIdentityResultImportedDeviceIdentityTypeserialNumber ImportedDeviceIdentityResultImportedDeviceIdentityType = "SerialNumber"
	ImportedDeviceIdentityResultImportedDeviceIdentityTypeunknown      ImportedDeviceIdentityResultImportedDeviceIdentityType = "Unknown"
)

func PossibleValuesForImportedDeviceIdentityResultImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityResultImportedDeviceIdentityTypeimei),
		string(ImportedDeviceIdentityResultImportedDeviceIdentityTypeserialNumber),
		string(ImportedDeviceIdentityResultImportedDeviceIdentityTypeunknown),
	}
}

func parseImportedDeviceIdentityResultImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityResultImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityResultImportedDeviceIdentityType{
		"imei":         ImportedDeviceIdentityResultImportedDeviceIdentityTypeimei,
		"serialnumber": ImportedDeviceIdentityResultImportedDeviceIdentityTypeserialNumber,
		"unknown":      ImportedDeviceIdentityResultImportedDeviceIdentityTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultImportedDeviceIdentityType(input)
	return &out, nil
}
