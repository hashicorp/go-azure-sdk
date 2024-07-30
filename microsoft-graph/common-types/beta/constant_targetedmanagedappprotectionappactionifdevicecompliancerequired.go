package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Block TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "block"
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "warn"
	TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired = "wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Block),
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn),
		string(TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe),
	}
}

func (s *TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAppActionIfDeviceComplianceRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Block,
		"warn":  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn,
		"wipe":  TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
