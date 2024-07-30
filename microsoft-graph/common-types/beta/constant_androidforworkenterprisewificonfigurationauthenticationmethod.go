package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
