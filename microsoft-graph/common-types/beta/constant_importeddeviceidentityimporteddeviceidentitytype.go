package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityImportedDeviceIdentityType string

const (
	ImportedDeviceIdentityImportedDeviceIdentityType_Imei         ImportedDeviceIdentityImportedDeviceIdentityType = "imei"
	ImportedDeviceIdentityImportedDeviceIdentityType_SerialNumber ImportedDeviceIdentityImportedDeviceIdentityType = "serialNumber"
	ImportedDeviceIdentityImportedDeviceIdentityType_Unknown      ImportedDeviceIdentityImportedDeviceIdentityType = "unknown"
)

func PossibleValuesForImportedDeviceIdentityImportedDeviceIdentityType() []string {
	return []string{
		string(ImportedDeviceIdentityImportedDeviceIdentityType_Imei),
		string(ImportedDeviceIdentityImportedDeviceIdentityType_SerialNumber),
		string(ImportedDeviceIdentityImportedDeviceIdentityType_Unknown),
	}
}

func (s *ImportedDeviceIdentityImportedDeviceIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityImportedDeviceIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityImportedDeviceIdentityType(input string) (*ImportedDeviceIdentityImportedDeviceIdentityType, error) {
	vals := map[string]ImportedDeviceIdentityImportedDeviceIdentityType{
		"imei":         ImportedDeviceIdentityImportedDeviceIdentityType_Imei,
		"serialnumber": ImportedDeviceIdentityImportedDeviceIdentityType_SerialNumber,
		"unknown":      ImportedDeviceIdentityImportedDeviceIdentityType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityImportedDeviceIdentityType(input)
	return &out, nil
}
