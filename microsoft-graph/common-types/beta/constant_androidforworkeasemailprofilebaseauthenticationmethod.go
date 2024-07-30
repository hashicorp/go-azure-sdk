package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseAuthenticationMethod string

const (
	AndroidForWorkEasEmailProfileBaseAuthenticationMethod_Certificate         AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "certificate"
	AndroidForWorkEasEmailProfileBaseAuthenticationMethod_DerivedCredential   AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "derivedCredential"
	AndroidForWorkEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethod_Certificate),
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethod_DerivedCredential),
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidForWorkEasEmailProfileBaseAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkEasEmailProfileBaseAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkEasEmailProfileBaseAuthenticationMethod(input string) (*AndroidForWorkEasEmailProfileBaseAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseAuthenticationMethod{
		"certificate":         AndroidForWorkEasEmailProfileBaseAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidForWorkEasEmailProfileBaseAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidForWorkEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseAuthenticationMethod(input)
	return &out, nil
}
