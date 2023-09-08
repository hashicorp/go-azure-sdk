package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cIdentityUserFlowUserFlowType string

const (
	B2cIdentityUserFlowUserFlowTypepasswordReset  B2cIdentityUserFlowUserFlowType = "PasswordReset"
	B2cIdentityUserFlowUserFlowTypeprofileUpdate  B2cIdentityUserFlowUserFlowType = "ProfileUpdate"
	B2cIdentityUserFlowUserFlowTyperesourceOwner  B2cIdentityUserFlowUserFlowType = "ResourceOwner"
	B2cIdentityUserFlowUserFlowTypesignIn         B2cIdentityUserFlowUserFlowType = "SignIn"
	B2cIdentityUserFlowUserFlowTypesignUp         B2cIdentityUserFlowUserFlowType = "SignUp"
	B2cIdentityUserFlowUserFlowTypesignUpOrSignIn B2cIdentityUserFlowUserFlowType = "SignUpOrSignIn"
)

func PossibleValuesForB2cIdentityUserFlowUserFlowType() []string {
	return []string{
		string(B2cIdentityUserFlowUserFlowTypepasswordReset),
		string(B2cIdentityUserFlowUserFlowTypeprofileUpdate),
		string(B2cIdentityUserFlowUserFlowTyperesourceOwner),
		string(B2cIdentityUserFlowUserFlowTypesignIn),
		string(B2cIdentityUserFlowUserFlowTypesignUp),
		string(B2cIdentityUserFlowUserFlowTypesignUpOrSignIn),
	}
}

func parseB2cIdentityUserFlowUserFlowType(input string) (*B2cIdentityUserFlowUserFlowType, error) {
	vals := map[string]B2cIdentityUserFlowUserFlowType{
		"passwordreset":  B2cIdentityUserFlowUserFlowTypepasswordReset,
		"profileupdate":  B2cIdentityUserFlowUserFlowTypeprofileUpdate,
		"resourceowner":  B2cIdentityUserFlowUserFlowTyperesourceOwner,
		"signin":         B2cIdentityUserFlowUserFlowTypesignIn,
		"signup":         B2cIdentityUserFlowUserFlowTypesignUp,
		"signuporsignin": B2cIdentityUserFlowUserFlowTypesignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := B2cIdentityUserFlowUserFlowType(input)
	return &out, nil
}
