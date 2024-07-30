package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidVpnConfigurationAuthenticationMethod string

const (
	AndroidVpnConfigurationAuthenticationMethod_AzureAD             AndroidVpnConfigurationAuthenticationMethod = "azureAD"
	AndroidVpnConfigurationAuthenticationMethod_Certificate         AndroidVpnConfigurationAuthenticationMethod = "certificate"
	AndroidVpnConfigurationAuthenticationMethod_DerivedCredential   AndroidVpnConfigurationAuthenticationMethod = "derivedCredential"
	AndroidVpnConfigurationAuthenticationMethod_SharedSecret        AndroidVpnConfigurationAuthenticationMethod = "sharedSecret"
	AndroidVpnConfigurationAuthenticationMethod_UsernameAndPassword AndroidVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidVpnConfigurationAuthenticationMethod_AzureAD),
		string(AndroidVpnConfigurationAuthenticationMethod_Certificate),
		string(AndroidVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidVpnConfigurationAuthenticationMethod_SharedSecret),
		string(AndroidVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidVpnConfigurationAuthenticationMethod(input string) (*AndroidVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         AndroidVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        AndroidVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": AndroidVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
