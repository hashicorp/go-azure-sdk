package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanHigh(input)
	return &out, nil
}
