package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xIdentityUserFlowUserFlowType string

const (
	B2xIdentityUserFlowUserFlowTypepasswordReset  B2xIdentityUserFlowUserFlowType = "PasswordReset"
	B2xIdentityUserFlowUserFlowTypeprofileUpdate  B2xIdentityUserFlowUserFlowType = "ProfileUpdate"
	B2xIdentityUserFlowUserFlowTyperesourceOwner  B2xIdentityUserFlowUserFlowType = "ResourceOwner"
	B2xIdentityUserFlowUserFlowTypesignIn         B2xIdentityUserFlowUserFlowType = "SignIn"
	B2xIdentityUserFlowUserFlowTypesignUp         B2xIdentityUserFlowUserFlowType = "SignUp"
	B2xIdentityUserFlowUserFlowTypesignUpOrSignIn B2xIdentityUserFlowUserFlowType = "SignUpOrSignIn"
)

func PossibleValuesForB2xIdentityUserFlowUserFlowType() []string {
	return []string{
		string(B2xIdentityUserFlowUserFlowTypepasswordReset),
		string(B2xIdentityUserFlowUserFlowTypeprofileUpdate),
		string(B2xIdentityUserFlowUserFlowTyperesourceOwner),
		string(B2xIdentityUserFlowUserFlowTypesignIn),
		string(B2xIdentityUserFlowUserFlowTypesignUp),
		string(B2xIdentityUserFlowUserFlowTypesignUpOrSignIn),
	}
}

func parseB2xIdentityUserFlowUserFlowType(input string) (*B2xIdentityUserFlowUserFlowType, error) {
	vals := map[string]B2xIdentityUserFlowUserFlowType{
		"passwordreset":  B2xIdentityUserFlowUserFlowTypepasswordReset,
		"profileupdate":  B2xIdentityUserFlowUserFlowTypeprofileUpdate,
		"resourceowner":  B2xIdentityUserFlowUserFlowTyperesourceOwner,
		"signin":         B2xIdentityUserFlowUserFlowTypesignIn,
		"signup":         B2xIdentityUserFlowUserFlowTypesignUp,
		"signuporsignin": B2xIdentityUserFlowUserFlowTypesignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := B2xIdentityUserFlowUserFlowType(input)
	return &out, nil
}
