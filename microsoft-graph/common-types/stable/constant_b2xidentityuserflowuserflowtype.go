package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xIdentityUserFlowUserFlowType string

const (
	B2xIdentityUserFlowUserFlowType_PasswordReset  B2xIdentityUserFlowUserFlowType = "passwordReset"
	B2xIdentityUserFlowUserFlowType_ProfileUpdate  B2xIdentityUserFlowUserFlowType = "profileUpdate"
	B2xIdentityUserFlowUserFlowType_ResourceOwner  B2xIdentityUserFlowUserFlowType = "resourceOwner"
	B2xIdentityUserFlowUserFlowType_SignIn         B2xIdentityUserFlowUserFlowType = "signIn"
	B2xIdentityUserFlowUserFlowType_SignUp         B2xIdentityUserFlowUserFlowType = "signUp"
	B2xIdentityUserFlowUserFlowType_SignUpOrSignIn B2xIdentityUserFlowUserFlowType = "signUpOrSignIn"
)

func PossibleValuesForB2xIdentityUserFlowUserFlowType() []string {
	return []string{
		string(B2xIdentityUserFlowUserFlowType_PasswordReset),
		string(B2xIdentityUserFlowUserFlowType_ProfileUpdate),
		string(B2xIdentityUserFlowUserFlowType_ResourceOwner),
		string(B2xIdentityUserFlowUserFlowType_SignIn),
		string(B2xIdentityUserFlowUserFlowType_SignUp),
		string(B2xIdentityUserFlowUserFlowType_SignUpOrSignIn),
	}
}

func (s *B2xIdentityUserFlowUserFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseB2xIdentityUserFlowUserFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseB2xIdentityUserFlowUserFlowType(input string) (*B2xIdentityUserFlowUserFlowType, error) {
	vals := map[string]B2xIdentityUserFlowUserFlowType{
		"passwordreset":  B2xIdentityUserFlowUserFlowType_PasswordReset,
		"profileupdate":  B2xIdentityUserFlowUserFlowType_ProfileUpdate,
		"resourceowner":  B2xIdentityUserFlowUserFlowType_ResourceOwner,
		"signin":         B2xIdentityUserFlowUserFlowType_SignIn,
		"signup":         B2xIdentityUserFlowUserFlowType_SignUp,
		"signuporsignin": B2xIdentityUserFlowUserFlowType_SignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := B2xIdentityUserFlowUserFlowType(input)
	return &out, nil
}
