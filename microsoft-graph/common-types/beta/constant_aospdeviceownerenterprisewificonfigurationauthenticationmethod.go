package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
