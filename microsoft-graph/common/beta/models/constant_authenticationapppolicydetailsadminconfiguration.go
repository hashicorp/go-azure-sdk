package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsAdminConfiguration string

const (
	AuthenticationAppPolicyDetailsAdminConfigurationdisabled      AuthenticationAppPolicyDetailsAdminConfiguration = "Disabled"
	AuthenticationAppPolicyDetailsAdminConfigurationenabled       AuthenticationAppPolicyDetailsAdminConfiguration = "Enabled"
	AuthenticationAppPolicyDetailsAdminConfigurationnotApplicable AuthenticationAppPolicyDetailsAdminConfiguration = "NotApplicable"
)

func PossibleValuesForAuthenticationAppPolicyDetailsAdminConfiguration() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsAdminConfigurationdisabled),
		string(AuthenticationAppPolicyDetailsAdminConfigurationenabled),
		string(AuthenticationAppPolicyDetailsAdminConfigurationnotApplicable),
	}
}

func parseAuthenticationAppPolicyDetailsAdminConfiguration(input string) (*AuthenticationAppPolicyDetailsAdminConfiguration, error) {
	vals := map[string]AuthenticationAppPolicyDetailsAdminConfiguration{
		"disabled":      AuthenticationAppPolicyDetailsAdminConfigurationdisabled,
		"enabled":       AuthenticationAppPolicyDetailsAdminConfigurationenabled,
		"notapplicable": AuthenticationAppPolicyDetailsAdminConfigurationnotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsAdminConfiguration(input)
	return &out, nil
}
