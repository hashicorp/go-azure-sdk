package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhone81VpnConfigurationAuthenticationMethod string

const (
	WindowsPhone81VpnConfigurationAuthenticationMethod_AzureAD             WindowsPhone81VpnConfigurationAuthenticationMethod = "azureAD"
	WindowsPhone81VpnConfigurationAuthenticationMethod_Certificate         WindowsPhone81VpnConfigurationAuthenticationMethod = "certificate"
	WindowsPhone81VpnConfigurationAuthenticationMethod_DerivedCredential   WindowsPhone81VpnConfigurationAuthenticationMethod = "derivedCredential"
	WindowsPhone81VpnConfigurationAuthenticationMethod_SharedSecret        WindowsPhone81VpnConfigurationAuthenticationMethod = "sharedSecret"
	WindowsPhone81VpnConfigurationAuthenticationMethod_UsernameAndPassword WindowsPhone81VpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindowsPhone81VpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(WindowsPhone81VpnConfigurationAuthenticationMethod_AzureAD),
		string(WindowsPhone81VpnConfigurationAuthenticationMethod_Certificate),
		string(WindowsPhone81VpnConfigurationAuthenticationMethod_DerivedCredential),
		string(WindowsPhone81VpnConfigurationAuthenticationMethod_SharedSecret),
		string(WindowsPhone81VpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WindowsPhone81VpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsPhone81VpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsPhone81VpnConfigurationAuthenticationMethod(input string) (*WindowsPhone81VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]WindowsPhone81VpnConfigurationAuthenticationMethod{
		"azuread":             WindowsPhone81VpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         WindowsPhone81VpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   WindowsPhone81VpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        WindowsPhone81VpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": WindowsPhone81VpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhone81VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
