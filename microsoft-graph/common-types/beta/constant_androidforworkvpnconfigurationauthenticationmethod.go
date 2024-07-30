package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConfigurationAuthenticationMethod string

const (
	AndroidForWorkVpnConfigurationAuthenticationMethod_AzureAD             AndroidForWorkVpnConfigurationAuthenticationMethod = "azureAD"
	AndroidForWorkVpnConfigurationAuthenticationMethod_Certificate         AndroidForWorkVpnConfigurationAuthenticationMethod = "certificate"
	AndroidForWorkVpnConfigurationAuthenticationMethod_DerivedCredential   AndroidForWorkVpnConfigurationAuthenticationMethod = "derivedCredential"
	AndroidForWorkVpnConfigurationAuthenticationMethod_SharedSecret        AndroidForWorkVpnConfigurationAuthenticationMethod = "sharedSecret"
	AndroidForWorkVpnConfigurationAuthenticationMethod_UsernameAndPassword AndroidForWorkVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidForWorkVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkVpnConfigurationAuthenticationMethod_AzureAD),
		string(AndroidForWorkVpnConfigurationAuthenticationMethod_Certificate),
		string(AndroidForWorkVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidForWorkVpnConfigurationAuthenticationMethod_SharedSecret),
		string(AndroidForWorkVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidForWorkVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkVpnConfigurationAuthenticationMethod(input string) (*AndroidForWorkVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidForWorkVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         AndroidForWorkVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidForWorkVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        AndroidForWorkVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": AndroidForWorkVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
