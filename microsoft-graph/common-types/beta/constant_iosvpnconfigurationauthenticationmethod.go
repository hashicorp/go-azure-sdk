package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationAuthenticationMethod string

const (
	IosVpnConfigurationAuthenticationMethod_AzureAD             IosVpnConfigurationAuthenticationMethod = "azureAD"
	IosVpnConfigurationAuthenticationMethod_Certificate         IosVpnConfigurationAuthenticationMethod = "certificate"
	IosVpnConfigurationAuthenticationMethod_DerivedCredential   IosVpnConfigurationAuthenticationMethod = "derivedCredential"
	IosVpnConfigurationAuthenticationMethod_SharedSecret        IosVpnConfigurationAuthenticationMethod = "sharedSecret"
	IosVpnConfigurationAuthenticationMethod_UsernameAndPassword IosVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForIosVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosVpnConfigurationAuthenticationMethod_AzureAD),
		string(IosVpnConfigurationAuthenticationMethod_Certificate),
		string(IosVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(IosVpnConfigurationAuthenticationMethod_SharedSecret),
		string(IosVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *IosVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosVpnConfigurationAuthenticationMethod(input string) (*IosVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]IosVpnConfigurationAuthenticationMethod{
		"azuread":             IosVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         IosVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   IosVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        IosVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": IosVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
