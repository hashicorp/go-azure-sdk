package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_MultiAppMode  AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "multiAppMode"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_SingleAppMode AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp = "singleAppMode"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_MultiAppMode),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_SingleAppMode),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp{
		"multiappmode":  AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_MultiAppMode,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_NotConfigured,
		"singleappmode": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp_SingleAppMode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeUseManagedHomeScreenApp(input)
	return &out, nil
}
