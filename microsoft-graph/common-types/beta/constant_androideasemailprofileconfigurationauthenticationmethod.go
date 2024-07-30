package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationAuthenticationMethod string

const (
	AndroidEasEmailProfileConfigurationAuthenticationMethod_Certificate         AndroidEasEmailProfileConfigurationAuthenticationMethod = "certificate"
	AndroidEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential   AndroidEasEmailProfileConfigurationAuthenticationMethod = "derivedCredential"
	AndroidEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword AndroidEasEmailProfileConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationAuthenticationMethod_Certificate),
		string(AndroidEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidEasEmailProfileConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidEasEmailProfileConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidEasEmailProfileConfigurationAuthenticationMethod(input string) (*AndroidEasEmailProfileConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationAuthenticationMethod{
		"certificate":         AndroidEasEmailProfileConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidEasEmailProfileConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidEasEmailProfileConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationAuthenticationMethod(input)
	return &out, nil
}
