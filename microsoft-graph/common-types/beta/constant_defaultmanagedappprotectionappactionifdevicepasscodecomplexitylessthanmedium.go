package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium string

const (
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "block"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "warn"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input string) (*DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium{
		"block": DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input)
	return &out, nil
}
