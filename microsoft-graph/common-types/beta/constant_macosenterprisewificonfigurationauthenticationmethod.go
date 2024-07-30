package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	MacOSEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	MacOSEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	MacOSEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *MacOSEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*MacOSEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         MacOSEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   MacOSEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": MacOSEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
