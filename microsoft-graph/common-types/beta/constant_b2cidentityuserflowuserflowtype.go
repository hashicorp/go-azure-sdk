package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cIdentityUserFlowUserFlowType string

const (
	B2cIdentityUserFlowUserFlowType_PasswordReset  B2cIdentityUserFlowUserFlowType = "passwordReset"
	B2cIdentityUserFlowUserFlowType_ProfileUpdate  B2cIdentityUserFlowUserFlowType = "profileUpdate"
	B2cIdentityUserFlowUserFlowType_ResourceOwner  B2cIdentityUserFlowUserFlowType = "resourceOwner"
	B2cIdentityUserFlowUserFlowType_SignIn         B2cIdentityUserFlowUserFlowType = "signIn"
	B2cIdentityUserFlowUserFlowType_SignUp         B2cIdentityUserFlowUserFlowType = "signUp"
	B2cIdentityUserFlowUserFlowType_SignUpOrSignIn B2cIdentityUserFlowUserFlowType = "signUpOrSignIn"
)

func PossibleValuesForB2cIdentityUserFlowUserFlowType() []string {
	return []string{
		string(B2cIdentityUserFlowUserFlowType_PasswordReset),
		string(B2cIdentityUserFlowUserFlowType_ProfileUpdate),
		string(B2cIdentityUserFlowUserFlowType_ResourceOwner),
		string(B2cIdentityUserFlowUserFlowType_SignIn),
		string(B2cIdentityUserFlowUserFlowType_SignUp),
		string(B2cIdentityUserFlowUserFlowType_SignUpOrSignIn),
	}
}

func (s *B2cIdentityUserFlowUserFlowType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseB2cIdentityUserFlowUserFlowType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseB2cIdentityUserFlowUserFlowType(input string) (*B2cIdentityUserFlowUserFlowType, error) {
	vals := map[string]B2cIdentityUserFlowUserFlowType{
		"passwordreset":  B2cIdentityUserFlowUserFlowType_PasswordReset,
		"profileupdate":  B2cIdentityUserFlowUserFlowType_ProfileUpdate,
		"resourceowner":  B2cIdentityUserFlowUserFlowType_ResourceOwner,
		"signin":         B2cIdentityUserFlowUserFlowType_SignIn,
		"signup":         B2cIdentityUserFlowUserFlowType_SignUp,
		"signuporsignin": B2cIdentityUserFlowUserFlowType_SignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := B2cIdentityUserFlowUserFlowType(input)
	return &out, nil
}
