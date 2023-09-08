package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowUserFlowType string

const (
	IdentityUserFlowUserFlowTypepasswordReset  IdentityUserFlowUserFlowType = "PasswordReset"
	IdentityUserFlowUserFlowTypeprofileUpdate  IdentityUserFlowUserFlowType = "ProfileUpdate"
	IdentityUserFlowUserFlowTyperesourceOwner  IdentityUserFlowUserFlowType = "ResourceOwner"
	IdentityUserFlowUserFlowTypesignIn         IdentityUserFlowUserFlowType = "SignIn"
	IdentityUserFlowUserFlowTypesignUp         IdentityUserFlowUserFlowType = "SignUp"
	IdentityUserFlowUserFlowTypesignUpOrSignIn IdentityUserFlowUserFlowType = "SignUpOrSignIn"
)

func PossibleValuesForIdentityUserFlowUserFlowType() []string {
	return []string{
		string(IdentityUserFlowUserFlowTypepasswordReset),
		string(IdentityUserFlowUserFlowTypeprofileUpdate),
		string(IdentityUserFlowUserFlowTyperesourceOwner),
		string(IdentityUserFlowUserFlowTypesignIn),
		string(IdentityUserFlowUserFlowTypesignUp),
		string(IdentityUserFlowUserFlowTypesignUpOrSignIn),
	}
}

func parseIdentityUserFlowUserFlowType(input string) (*IdentityUserFlowUserFlowType, error) {
	vals := map[string]IdentityUserFlowUserFlowType{
		"passwordreset":  IdentityUserFlowUserFlowTypepasswordReset,
		"profileupdate":  IdentityUserFlowUserFlowTypeprofileUpdate,
		"resourceowner":  IdentityUserFlowUserFlowTyperesourceOwner,
		"signin":         IdentityUserFlowUserFlowTypesignIn,
		"signup":         IdentityUserFlowUserFlowTypesignUp,
		"signuporsignin": IdentityUserFlowUserFlowTypesignUpOrSignIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IdentityUserFlowUserFlowType(input)
	return &out, nil
}
