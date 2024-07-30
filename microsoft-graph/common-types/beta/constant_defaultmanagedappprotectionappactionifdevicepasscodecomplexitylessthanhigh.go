package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh string

const (
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "block"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "warn"
	DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn),
		string(DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input string) (*DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh{
		"block": DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input)
	return &out, nil
}
