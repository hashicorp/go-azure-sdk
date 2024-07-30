package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationAuthenticationMethod string

const (
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_Certificate         AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "certificate"
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_DerivedCredential   AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "derivedCredential"
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_Certificate),
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidForWorkNineWorkEasConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkNineWorkEasConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkNineWorkEasConfigurationAuthenticationMethod(input string) (*AndroidForWorkNineWorkEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidForWorkNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
