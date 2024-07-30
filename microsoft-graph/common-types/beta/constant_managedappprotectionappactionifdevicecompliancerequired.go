package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	ManagedAppProtectionAppActionIfDeviceComplianceRequired_Block ManagedAppProtectionAppActionIfDeviceComplianceRequired = "block"
	ManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn  ManagedAppProtectionAppActionIfDeviceComplianceRequired = "warn"
	ManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe  ManagedAppProtectionAppActionIfDeviceComplianceRequired = "wipe"
)

func PossibleValuesForManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequired_Block),
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn),
		string(ManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe),
	}
}

func (s *ManagedAppProtectionAppActionIfDeviceComplianceRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAppActionIfDeviceComplianceRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*ManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]ManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": ManagedAppProtectionAppActionIfDeviceComplianceRequired_Block,
		"warn":  ManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn,
		"wipe":  ManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
