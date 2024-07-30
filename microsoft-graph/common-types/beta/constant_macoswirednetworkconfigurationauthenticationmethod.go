package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationAuthenticationMethod string

const (
	MacOSWiredNetworkConfigurationAuthenticationMethod_Certificate         MacOSWiredNetworkConfigurationAuthenticationMethod = "certificate"
	MacOSWiredNetworkConfigurationAuthenticationMethod_DerivedCredential   MacOSWiredNetworkConfigurationAuthenticationMethod = "derivedCredential"
	MacOSWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword MacOSWiredNetworkConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForMacOSWiredNetworkConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationAuthenticationMethod_Certificate),
		string(MacOSWiredNetworkConfigurationAuthenticationMethod_DerivedCredential),
		string(MacOSWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *MacOSWiredNetworkConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSWiredNetworkConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSWiredNetworkConfigurationAuthenticationMethod(input string) (*MacOSWiredNetworkConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSWiredNetworkConfigurationAuthenticationMethod{
		"certificate":         MacOSWiredNetworkConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   MacOSWiredNetworkConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": MacOSWiredNetworkConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationAuthenticationMethod(input)
	return &out, nil
}
