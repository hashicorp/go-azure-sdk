package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_Administrators                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "administrators"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndInteractiveUsers Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "administratorsAndInteractiveUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndPowerUsers       Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "administratorsAndPowerUsers"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_NotConfigured                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser = "notConfigured"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_Administrators),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndInteractiveUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndPowerUsers),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_NotConfigured),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser{
		"administrators":                    Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_Administrators,
		"administratorsandinteractiveusers": Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndInteractiveUsers,
		"administratorsandpowerusers":       Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_AdministratorsAndPowerUsers,
		"notconfigured":                     Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsFormatAndEjectOfRemovableMediaAllowedUser(input)
	return &out, nil
}
