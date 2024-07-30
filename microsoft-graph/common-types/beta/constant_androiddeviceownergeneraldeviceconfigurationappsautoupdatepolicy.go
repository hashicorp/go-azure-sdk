package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Always        AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "always"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Never         AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "never"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_UserChoice    AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "userChoice"
	AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_WiFiOnly      AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy = "wiFiOnly"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Always),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Never),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_UserChoice),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_WiFiOnly),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy{
		"always":        AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Always,
		"never":         AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_Never,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_NotConfigured,
		"userchoice":    AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_UserChoice,
		"wifionly":      AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy_WiFiOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationAppsAutoUpdatePolicy(input)
	return &out, nil
}
