package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileVpnConfigurationAuthenticationMethod_AzureAD             AndroidWorkProfileVpnConfigurationAuthenticationMethod = "azureAD"
	AndroidWorkProfileVpnConfigurationAuthenticationMethod_Certificate         AndroidWorkProfileVpnConfigurationAuthenticationMethod = "certificate"
	AndroidWorkProfileVpnConfigurationAuthenticationMethod_DerivedCredential   AndroidWorkProfileVpnConfigurationAuthenticationMethod = "derivedCredential"
	AndroidWorkProfileVpnConfigurationAuthenticationMethod_SharedSecret        AndroidWorkProfileVpnConfigurationAuthenticationMethod = "sharedSecret"
	AndroidWorkProfileVpnConfigurationAuthenticationMethod_UsernameAndPassword AndroidWorkProfileVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethod_AzureAD),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethod_Certificate),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethod_SharedSecret),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidWorkProfileVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileVpnConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidWorkProfileVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         AndroidWorkProfileVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidWorkProfileVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        AndroidWorkProfileVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": AndroidWorkProfileVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
