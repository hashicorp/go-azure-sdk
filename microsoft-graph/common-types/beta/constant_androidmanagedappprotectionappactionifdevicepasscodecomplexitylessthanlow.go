package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow string

const (
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Block AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "block"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Warn  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "warn"
	AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Wipe  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Block),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Warn),
		string(AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input string) (*AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow{
		"block": AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDevicePasscodeComplexityLessThanLow(input)
	return &out, nil
}
