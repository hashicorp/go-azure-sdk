package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
