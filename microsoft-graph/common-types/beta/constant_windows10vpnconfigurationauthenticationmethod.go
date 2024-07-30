package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationAuthenticationMethod string

const (
	Windows10VpnConfigurationAuthenticationMethod_Certificate         Windows10VpnConfigurationAuthenticationMethod = "certificate"
	Windows10VpnConfigurationAuthenticationMethod_CustomEapXml        Windows10VpnConfigurationAuthenticationMethod = "customEapXml"
	Windows10VpnConfigurationAuthenticationMethod_DerivedCredential   Windows10VpnConfigurationAuthenticationMethod = "derivedCredential"
	Windows10VpnConfigurationAuthenticationMethod_UsernameAndPassword Windows10VpnConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindows10VpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(Windows10VpnConfigurationAuthenticationMethod_Certificate),
		string(Windows10VpnConfigurationAuthenticationMethod_CustomEapXml),
		string(Windows10VpnConfigurationAuthenticationMethod_DerivedCredential),
		string(Windows10VpnConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *Windows10VpnConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10VpnConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10VpnConfigurationAuthenticationMethod(input string) (*Windows10VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]Windows10VpnConfigurationAuthenticationMethod{
		"certificate":         Windows10VpnConfigurationAuthenticationMethod_Certificate,
		"customeapxml":        Windows10VpnConfigurationAuthenticationMethod_CustomEapXml,
		"derivedcredential":   Windows10VpnConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": Windows10VpnConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
