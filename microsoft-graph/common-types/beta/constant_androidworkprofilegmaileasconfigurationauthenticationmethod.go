package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_Certificate         AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "certificate"
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_DerivedCredential   AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "derivedCredential"
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_UsernameAndPassword AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_Certificate),
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidWorkProfileGmailEasConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGmailEasConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGmailEasConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileGmailEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidWorkProfileGmailEasConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
