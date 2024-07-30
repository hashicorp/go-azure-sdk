package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	AndroidEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	AndroidEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
