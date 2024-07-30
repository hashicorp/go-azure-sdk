package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_Floating      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "floating"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_SwipeUp       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType = "swipeUp"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_Floating),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_SwipeUp),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType{
		"floating":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_Floating,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_NotConfigured,
		"swipeup":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType_SwipeUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeVirtualHomeButtonType(input)
	return &out, nil
}
