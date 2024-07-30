package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationAuthenticationMethod string

const (
	AppleVpnConfigurationAuthenticationMethod_AzureAD             AppleVpnConfigurationAuthenticationMethod = "azureAD"
	AppleVpnConfigurationAuthenticationMethod_Certificate         AppleVpnConfigurationAuthenticationMethod = "certificate"
	AppleVpnConfigurationAuthenticationMethod_DerivedCredential   AppleVpnConfigurationAuthenticationMethod = "derivedCredential"
	AppleVpnConfigurationAuthenticationMethod_SharedSecret        AppleVpnConfigurationAuthenticationMethod = "sharedSecret"
	AppleVpnConfigurationAuthenticationMethod_UsernameAndPassword AppleVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAppleVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AppleVpnConfigurationAuthenticationMethod_AzureAD),
		string(AppleVpnConfigurationAuthenticationMethod_Certificate),
		string(AppleVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(AppleVpnConfigurationAuthenticationMethod_SharedSecret),
		string(AppleVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AppleVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnConfigurationAuthenticationMethod(input string) (*AppleVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AppleVpnConfigurationAuthenticationMethod{
		"azuread":             AppleVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         AppleVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AppleVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        AppleVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": AppleVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
