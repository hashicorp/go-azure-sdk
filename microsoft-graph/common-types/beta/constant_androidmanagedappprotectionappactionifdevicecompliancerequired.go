package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Block AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "block"
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "warn"
	AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Block),
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn),
		string(AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfDeviceComplianceRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
