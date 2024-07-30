package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior string

const (
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_ElevateWithoutPrompting                Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "elevateWithoutPrompting"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_NotConfigured                          Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "notConfigured"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsent                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "promptForConsent"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentForNonWindowsBinaries  Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "promptForConsentForNonWindowsBinaries"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentOnTheSecureDesktop     Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "promptForConsentOnTheSecureDesktop"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentials                   Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "promptForCredentials"
	Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior = "promptForCredentialsOnTheSecureDesktop"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_ElevateWithoutPrompting),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_NotConfigured),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsent),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentForNonWindowsBinaries),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentOnTheSecureDesktop),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentials),
		string(Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop),
	}
}

func (s *Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior(input string) (*Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior{
		"elevatewithoutprompting":                Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_ElevateWithoutPrompting,
		"notconfigured":                          Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_NotConfigured,
		"promptforconsent":                       Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsent,
		"promptforconsentfornonwindowsbinaries":  Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentForNonWindowsBinaries,
		"promptforconsentonthesecuredesktop":     Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForConsentOnTheSecureDesktop,
		"promptforcredentials":                   Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentials,
		"promptforcredentialsonthesecuredesktop": Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior_PromptForCredentialsOnTheSecureDesktop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationLocalSecurityOptionsAdministratorElevationPromptBehavior(input)
	return &out, nil
}
