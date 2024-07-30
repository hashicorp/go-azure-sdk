package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationAuthenticationMethod string

const (
	IosikEv2VpnConfigurationAuthenticationMethod_AzureAD             IosikEv2VpnConfigurationAuthenticationMethod = "azureAD"
	IosikEv2VpnConfigurationAuthenticationMethod_Certificate         IosikEv2VpnConfigurationAuthenticationMethod = "certificate"
	IosikEv2VpnConfigurationAuthenticationMethod_DerivedCredential   IosikEv2VpnConfigurationAuthenticationMethod = "derivedCredential"
	IosikEv2VpnConfigurationAuthenticationMethod_SharedSecret        IosikEv2VpnConfigurationAuthenticationMethod = "sharedSecret"
	IosikEv2VpnConfigurationAuthenticationMethod_UsernameAndPassword IosikEv2VpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForIosikEv2VpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosikEv2VpnConfigurationAuthenticationMethod_AzureAD),
		string(IosikEv2VpnConfigurationAuthenticationMethod_Certificate),
		string(IosikEv2VpnConfigurationAuthenticationMethod_DerivedCredential),
		string(IosikEv2VpnConfigurationAuthenticationMethod_SharedSecret),
		string(IosikEv2VpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *IosikEv2VpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationAuthenticationMethod(input string) (*IosikEv2VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]IosikEv2VpnConfigurationAuthenticationMethod{
		"azuread":             IosikEv2VpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         IosikEv2VpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   IosikEv2VpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        IosikEv2VpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": IosikEv2VpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
