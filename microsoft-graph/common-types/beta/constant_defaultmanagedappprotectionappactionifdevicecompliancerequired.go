package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Block DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "block"
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "warn"
	DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Block),
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn),
		string(DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfDeviceComplianceRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
