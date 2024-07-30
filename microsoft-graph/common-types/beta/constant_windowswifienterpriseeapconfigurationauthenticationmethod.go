package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod string

const (
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_Certificate         WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "certificate"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_DerivedCredential   WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "derivedCredential"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_UsernameAndPassword WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationAuthenticationMethod() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_Certificate),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_DerivedCredential),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationAuthenticationMethod(input string) (*WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod{
		"certificate":         WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod(input)
	return &out, nil
}
