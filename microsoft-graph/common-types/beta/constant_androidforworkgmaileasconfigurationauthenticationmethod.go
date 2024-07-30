package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationAuthenticationMethod string

const (
	AndroidForWorkGmailEasConfigurationAuthenticationMethod_Certificate         AndroidForWorkGmailEasConfigurationAuthenticationMethod = "certificate"
	AndroidForWorkGmailEasConfigurationAuthenticationMethod_DerivedCredential   AndroidForWorkGmailEasConfigurationAuthenticationMethod = "derivedCredential"
	AndroidForWorkGmailEasConfigurationAuthenticationMethod_UsernameAndPassword AndroidForWorkGmailEasConfigurationAuthenticationMethod = "usernameAndPassword"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethod_Certificate),
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethod_DerivedCredential),
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethod_UsernameAndPassword),
	}
}

func (s *AndroidForWorkGmailEasConfigurationAuthenticationMethod) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGmailEasConfigurationAuthenticationMethod(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGmailEasConfigurationAuthenticationMethod(input string) (*AndroidForWorkGmailEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkGmailEasConfigurationAuthenticationMethod_Certificate,
		"derivedcredential":   AndroidForWorkGmailEasConfigurationAuthenticationMethod_DerivedCredential,
		"usernameandpassword": AndroidForWorkGmailEasConfigurationAuthenticationMethod_UsernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
