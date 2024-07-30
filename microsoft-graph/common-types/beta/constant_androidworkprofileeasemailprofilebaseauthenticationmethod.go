package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod string

const (
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_Certificate         AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "certificate"
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_DerivedCredential   AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "derivedCredential"
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_Certificate),
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_DerivedCredential),
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileEasEmailProfileBaseAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileEasEmailProfileBaseAuthenticationMethod(input string) (*AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod{
		"certificate":         AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod(input)
	return &out, nil
}
