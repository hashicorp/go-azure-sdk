package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate         AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "certificate"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential   AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "derivedCredential"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
