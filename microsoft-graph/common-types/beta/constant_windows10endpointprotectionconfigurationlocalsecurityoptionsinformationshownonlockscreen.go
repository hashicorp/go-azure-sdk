package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_DoNotDisplayUser          Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "doNotDisplayUser"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_NotConfigured             Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "notConfigured"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameDomainUser Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "userDisplayNameDomainUser"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameOnly       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen = "userDisplayNameOnly"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_DoNotDisplayUser),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_NotConfigured),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameDomainUser),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameOnly),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen{
		"donotdisplayuser":          Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_DoNotDisplayUser,
		"notconfigured":             Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_NotConfigured,
		"userdisplaynamedomainuser": Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameDomainUser,
		"userdisplaynameonly":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen_UserDisplayNameOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsInformationShownOnLockScreen(input)
	return &out, nil
}
