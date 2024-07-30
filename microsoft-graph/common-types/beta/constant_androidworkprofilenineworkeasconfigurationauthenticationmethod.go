package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_Certificate         AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "certificate"
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_DerivedCredential   AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "derivedCredential"
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_Certificate),
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
