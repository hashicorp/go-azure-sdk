package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultDiscoverySource string

const (
	ImportedAppleDeviceIdentityResultDiscoverySource_AdminImport             ImportedAppleDeviceIdentityResultDiscoverySource = "adminImport"
	ImportedAppleDeviceIdentityResultDiscoverySource_DeviceEnrollmentProgram ImportedAppleDeviceIdentityResultDiscoverySource = "deviceEnrollmentProgram"
	ImportedAppleDeviceIdentityResultDiscoverySource_Unknown                 ImportedAppleDeviceIdentityResultDiscoverySource = "unknown"
)

func PossibleValuesForImportedAppleDeviceIdentityResultDiscoverySource() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultDiscoverySource_AdminImport),
		string(ImportedAppleDeviceIdentityResultDiscoverySource_DeviceEnrollmentProgram),
		string(ImportedAppleDeviceIdentityResultDiscoverySource_Unknown),
	}
}

func (s *ImportedAppleDeviceIdentityResultDiscoverySource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityResultDiscoverySource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityResultDiscoverySource(input string) (*ImportedAppleDeviceIdentityResultDiscoverySource, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultDiscoverySource{
		"adminimport":             ImportedAppleDeviceIdentityResultDiscoverySource_AdminImport,
		"deviceenrollmentprogram": ImportedAppleDeviceIdentityResultDiscoverySource_DeviceEnrollmentProgram,
		"unknown":                 ImportedAppleDeviceIdentityResultDiscoverySource_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultDiscoverySource(input)
	return &out, nil
}
