package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowUserFlowType string

const (
	IdentityUserFlowUserFlowType_PasswordReset  IdentityUserFlowUserFlowType = "passwordReset"
	IdentityUserFlowUserFlowType_ProfileUpdate  IdentityUserFlowUserFlowType = "profileUpdate"
	IdentityUserFlowUserFlowType_ResourceOwner  IdentityUserFlowUserFlowType = "resourceOwner"
	IdentityUserFlowUserFlowType_SignIn         IdentityUserFlowUserFlowType = "signIn"
	IdentityUserFlowUserFlowType_SignUp         IdentityUserFlowUserFlowType = "signUp"
	IdentityUserFlowUserFlowType_SignUpOrSignIn IdentityUserFlowUserFlowType = "signUpOrSignIn"
)

func PossibleValuesForIdentityUserFlowUserFlowType() []string {
	return []string{
		string(IdentityUserFlowUserFlowType_PasswordReset),
		string(IdentityUserFlowUserFlowType_ProfileUpdate),
		string(IdentityUserFlowUserFlowType_ResourceOwner),
		string(IdentityUserFlowUserFlowType_SignIn),
		string(IdentityUserFlowUserFlowType_SignUp),
		string(IdentityUserFlowUserFlowType_SignUpOrSignIn),
	}
}

func (s *IdentityUserFlowUserFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIdentityUserFlowUserFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIdentityUserFlowUserFlowType(input string) (*IdentityUserFlowUserFlowType, error) {
	vals := map[string]IdentityUserFlowUserFlowType{
		"passwordreset":  IdentityUserFlowUserFlowType_PasswordReset,
		"profileupdate":  IdentityUserFlowUserFlowType_ProfileUpdate,
		"resourceowner":  IdentityUserFlowUserFlowType_ResourceOwner,
		"signin":         IdentityUserFlowUserFlowType_SignIn,
		"signup":         IdentityUserFlowUserFlowType_SignUp,
		"signuporsignin": IdentityUserFlowUserFlowType_SignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowUserFlowType(input)
	return &out, nil
}
