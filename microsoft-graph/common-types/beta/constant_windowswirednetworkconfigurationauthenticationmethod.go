package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationAuthenticationMethod string

const (
	WindowsWiredNetworkConfigurationAuthenticationMethod_Certificate         WindowsWiredNetworkConfigurationAuthenticationMethod = "certificate"
	WindowsWiredNetworkConfigurationAuthenticationMethod_DerivedCredential   WindowsWiredNetworkConfigurationAuthenticationMethod = "derivedCredential"
	WindowsWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword WindowsWiredNetworkConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationAuthenticationMethod() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationAuthenticationMethod_Certificate),
		string(WindowsWiredNetworkConfigurationAuthenticationMethod_DerivedCredential),
		string(WindowsWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WindowsWiredNetworkConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWiredNetworkConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWiredNetworkConfigurationAuthenticationMethod(input string) (*WindowsWiredNetworkConfigurationAuthenticationMethod, error) {
	vals := map[string]WindowsWiredNetworkConfigurationAuthenticationMethod{
		"certificate":         WindowsWiredNetworkConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   WindowsWiredNetworkConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": WindowsWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationAuthenticationMethod(input)
	return &out, nil
}
