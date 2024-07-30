package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_AutomaticallyDenyElevationRequests     Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior = "automaticallyDenyElevationRequests"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_NotConfigured                          Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior = "notConfigured"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentials                   Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior = "promptForCredentials"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior = "promptForCredentialsOnTheSecureDesktop"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_AutomaticallyDenyElevationRequests),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_NotConfigured),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentials),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior{
		"automaticallydenyelevationrequests":     Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_AutomaticallyDenyElevationRequests,
		"notconfigured":                          Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_NotConfigured,
		"promptforcredentials":                   Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentials,
		"promptforcredentialsonthesecuredesktop": Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsStandardUserElevationPromptBehavior(input)
	return &out, nil
}
