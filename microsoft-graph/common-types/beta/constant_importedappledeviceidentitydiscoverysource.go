package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityDiscoverySource string

const (
	ImportedAppleDeviceIdentityDiscoverySource_AdminImport             ImportedAppleDeviceIdentityDiscoverySource = "adminImport"
	ImportedAppleDeviceIdentityDiscoverySource_DeviceEnrollmentProgram ImportedAppleDeviceIdentityDiscoverySource = "deviceEnrollmentProgram"
	ImportedAppleDeviceIdentityDiscoverySource_Unknown                 ImportedAppleDeviceIdentityDiscoverySource = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityDiscoverySource() []string {
	return []string{
		string(ImportedAppleDeviceIdentityDiscoverySource_AdminImport),
		string(ImportedAppleDeviceIdentityDiscoverySource_DeviceEnrollmentProgram),
		string(ImportedAppleDeviceIdentityDiscoverySource_Unknown),
	}
}

func (s *ImportedAppleDeviceIdentityDiscoverySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityDiscoverySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityDiscoverySource(input string) (*ImportedAppleDeviceIdentityDiscoverySource, error) {
	vals := map[string]ImportedAppleDeviceIdentityDiscoverySource{
		"adminimport":             ImportedAppleDeviceIdentityDiscoverySource_AdminImport,
		"deviceenrollmentprogram": ImportedAppleDeviceIdentityDiscoverySource_DeviceEnrollmentProgram,
		"unknown":                 ImportedAppleDeviceIdentityDiscoverySource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityDiscoverySource(input)
	return &out, nil
}
