package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnConfigurationAuthenticationMethod string

const (
	VpnConfigurationAuthenticationMethod_AzureAD             VpnConfigurationAuthenticationMethod = "azureAD"
	VpnConfigurationAuthenticationMethod_Certificate         VpnConfigurationAuthenticationMethod = "certificate"
	VpnConfigurationAuthenticationMethod_DerivedCredential   VpnConfigurationAuthenticationMethod = "derivedCredential"
	VpnConfigurationAuthenticationMethod_SharedSecret        VpnConfigurationAuthenticationMethod = "sharedSecret"
	VpnConfigurationAuthenticationMethod_UsernameAndPassword VpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(VpnConfigurationAuthenticationMethod_AzureAD),
		string(VpnConfigurationAuthenticationMethod_Certificate),
		string(VpnConfigurationAuthenticationMethod_DerivedCredential),
		string(VpnConfigurationAuthenticationMethod_SharedSecret),
		string(VpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *VpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnConfigurationAuthenticationMethod(input string) (*VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]VpnConfigurationAuthenticationMethod{
		"azuread":             VpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         VpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   VpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        VpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": VpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
