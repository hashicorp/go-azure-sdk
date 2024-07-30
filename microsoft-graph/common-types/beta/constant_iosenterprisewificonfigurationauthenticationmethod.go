package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	IosEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         IosEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	IosEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   IosEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	IosEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword IosEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(IosEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(IosEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *IosEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*IosEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         IosEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   IosEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": IosEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
