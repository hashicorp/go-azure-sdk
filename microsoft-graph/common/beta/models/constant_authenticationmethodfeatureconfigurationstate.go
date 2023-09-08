package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodFeatureConfigurationState string

const (
	AuthenticationMethodFeatureConfigurationStatedefault  AuthenticationMethodFeatureConfigurationState = "Default"
	AuthenticationMethodFeatureConfigurationStatedisabled AuthenticationMethodFeatureConfigurationState = "Disabled"
	AuthenticationMethodFeatureConfigurationStateenabled  AuthenticationMethodFeatureConfigurationState = "Enabled"
)

func PossibleValuesForAuthenticationMethodFeatureConfigurationState() []string {
	return []string{
		string(AuthenticationMethodFeatureConfigurationStatedefault),
		string(AuthenticationMethodFeatureConfigurationStatedisabled),
		string(AuthenticationMethodFeatureConfigurationStateenabled),
	}
}

func parseAuthenticationMethodFeatureConfigurationState(input string) (*AuthenticationMethodFeatureConfigurationState, error) {
	vals := map[string]AuthenticationMethodFeatureConfigurationState{
		"default":  AuthenticationMethodFeatureConfigurationStatedefault,
		"disabled": AuthenticationMethodFeatureConfigurationStatedisabled,
		"enabled":  AuthenticationMethodFeatureConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodFeatureConfigurationState(input)
	return &out, nil
}
