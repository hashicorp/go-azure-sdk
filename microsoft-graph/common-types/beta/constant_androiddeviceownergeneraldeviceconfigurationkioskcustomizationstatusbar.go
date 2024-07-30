package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotConfigured                     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "notificationsAndSystemInfoEnabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_SystemInfoOnly                    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar = "systemInfoOnly"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_SystemInfoOnly),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar{
		"notconfigured":                     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotConfigured,
		"notificationsandsysteminfoenabled": AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_NotificationsAndSystemInfoEnabled,
		"systeminfoonly":                    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar_SystemInfoOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationStatusBar(input)
	return &out, nil
}
