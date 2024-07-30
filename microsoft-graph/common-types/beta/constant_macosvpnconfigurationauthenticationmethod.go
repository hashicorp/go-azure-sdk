package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationAuthenticationMethod string

const (
	MacOSVpnConfigurationAuthenticationMethod_AzureAD             MacOSVpnConfigurationAuthenticationMethod = "azureAD"
	MacOSVpnConfigurationAuthenticationMethod_Certificate         MacOSVpnConfigurationAuthenticationMethod = "certificate"
	MacOSVpnConfigurationAuthenticationMethod_DerivedCredential   MacOSVpnConfigurationAuthenticationMethod = "derivedCredential"
	MacOSVpnConfigurationAuthenticationMethod_SharedSecret        MacOSVpnConfigurationAuthenticationMethod = "sharedSecret"
	MacOSVpnConfigurationAuthenticationMethod_UsernameAndPassword MacOSVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForMacOSVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSVpnConfigurationAuthenticationMethod_AzureAD),
		string(MacOSVpnConfigurationAuthenticationMethod_Certificate),
		string(MacOSVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(MacOSVpnConfigurationAuthenticationMethod_SharedSecret),
		string(MacOSVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *MacOSVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSVpnConfigurationAuthenticationMethod(input string) (*MacOSVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSVpnConfigurationAuthenticationMethod{
		"azuread":             MacOSVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         MacOSVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   MacOSVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        MacOSVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": MacOSVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
