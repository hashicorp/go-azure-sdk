package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityResultImportedDeviceIdentityType_Imei         ImportedDeviceIdentityResultImportedDeviceIdentityType = "imei"
	ImportedDeviceIdentityResultImportedDeviceIdentityType_SerialNumber ImportedDeviceIdentityResultImportedDeviceIdentityType = "serialNumber"
	ImportedDeviceIdentityResultImportedDeviceIdentityType_Unknown      ImportedDeviceIdentityResultImportedDeviceIdentityType = "unknown"
)

func PossibleValuesForImportedDeviceIdentityResultImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityResultImportedDeviceIdentityType_Imei),
		string(ImportedDeviceIdentityResultImportedDeviceIdentityType_SerialNumber),
		string(ImportedDeviceIdentityResultImportedDeviceIdentityType_Unknown),
	}
}

func (s *ImportedDeviceIdentityResultImportedDeviceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityResultImportedDeviceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityResultImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityResultImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityResultImportedDeviceIdentityType{
		"imei":         ImportedDeviceIdentityResultImportedDeviceIdentityType_Imei,
		"serialnumber": ImportedDeviceIdentityResultImportedDeviceIdentityType_SerialNumber,
		"unknown":      ImportedDeviceIdentityResultImportedDeviceIdentityType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultImportedDeviceIdentityType(input)
	return &out, nil
}
