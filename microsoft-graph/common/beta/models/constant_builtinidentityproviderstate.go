package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BuiltInIdentityProviderState string

const (
	BuiltInIdentityProviderStatedisabled BuiltInIdentityProviderState = "Disabled"
	BuiltInIdentityProviderStateenabled  BuiltInIdentityProviderState = "Enabled"
)

func PossibleValuesForBuiltInIdentityProviderState() []string {
	return []string{
		string(BuiltInIdentityProviderStatedisabled),
		string(BuiltInIdentityProviderStateenabled),
	}
}

func parseBuiltInIdentityProviderState(input string) (*BuiltInIdentityProviderState, error) {
	vals := map[string]BuiltInIdentityProviderState{
		"disabled": BuiltInIdentityProviderStatedisabled,
		"enabled":  BuiltInIdentityProviderStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BuiltInIdentityProviderState(input)
	return &out, nil
}
