package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_HomeButtonOnly    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "homeButtonOnly"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NavigationEnabled AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "navigationEnabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NotConfigured     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_HomeButtonOnly),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NavigationEnabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation{
		"homebuttononly":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_HomeButtonOnly,
		"navigationenabled": AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NavigationEnabled,
		"notconfigured":     AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskCustomizationSystemNavigation(input)
	return &out, nil
}
