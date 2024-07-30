package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerVpnConfigurationAuthenticationMethod string

const (
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_AzureAD             AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "azureAD"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_Certificate         AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "certificate"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_DerivedCredential   AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "derivedCredential"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_SharedSecret        AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "sharedSecret"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_UsernameAndPassword AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidDeviceOwnerVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_AzureAD),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_Certificate),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_SharedSecret),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidDeviceOwnerVpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerVpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerVpnConfigurationAuthenticationMethod(input string) (*AndroidDeviceOwnerVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidDeviceOwnerVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_AzureAD,
		"certificate":         AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_DerivedCredential,
		"sharedsecret":        AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_SharedSecret,
		"usernameandpassword": AndroidDeviceOwnerVpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
