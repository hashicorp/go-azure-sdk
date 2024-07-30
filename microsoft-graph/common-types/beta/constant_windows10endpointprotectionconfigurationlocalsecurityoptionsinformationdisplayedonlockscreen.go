package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_Administrators                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "administrators"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndInteractiveUsers Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "administratorsAndInteractiveUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndPowerUsers       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "administratorsAndPowerUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_NotConfigured                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_Administrators),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndInteractiveUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndPowerUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen{
		"administrators":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_Administrators,
		"administratorsandinteractiveusers": Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_AdministratorsAndPowerUsers,
		"notconfigured":                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationDisplayedOnLockScreen(input)
	return &out, nil
}
