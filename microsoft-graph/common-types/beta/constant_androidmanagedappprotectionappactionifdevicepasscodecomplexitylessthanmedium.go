package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanMedium(input)
	return &out, nil
}
