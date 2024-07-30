package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationAuthenticationMethod string

const (
	IosEasEmailProfileConfigurationAuthenticationMethod_Certificate         IosEasEmailProfileConfigurationAuthenticationMethod = "certificate"
	IosEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential   IosEasEmailProfileConfigurationAuthenticationMethod = "derivedCredential"
	IosEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword IosEasEmailProfileConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForIosEasEmailProfileConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosEasEmailProfileConfigurationAuthenticationMethod_Certificate),
		string(IosEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential),
		string(IosEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *IosEasEmailProfileConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosEasEmailProfileConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosEasEmailProfileConfigurationAuthenticationMethod(input string) (*IosEasEmailProfileConfigurationAuthenticationMethod, error) {
	vals := map[string]IosEasEmailProfileConfigurationAuthenticationMethod{
		"certificate":         IosEasEmailProfileConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   IosEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": IosEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationAuthenticationMethod(input)
	return &out, nil
}
