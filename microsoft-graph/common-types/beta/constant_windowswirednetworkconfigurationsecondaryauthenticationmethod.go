package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod string

const (
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_Certificate         WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "certificate"
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_DerivedCredential   WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "derivedCredential"
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_UsernameAndPassword WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationSecondaryAuthenticationMethod() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_Certificate),
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_DerivedCredential),
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWiredNetworkConfigurationSecondaryAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWiredNetworkConfigurationSecondaryAuthenticationMethod(input string) (*WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod, error) {
	vals := map[string]WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod{
		"certificate":         WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_Certificate,
		"derivedcredential":   WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_DerivedCredential,
		"usernameandpassword": WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod(input)
	return &out, nil
}
