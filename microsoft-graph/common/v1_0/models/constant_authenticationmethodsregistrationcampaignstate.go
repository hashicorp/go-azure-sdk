package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignState string

const (
	AuthenticationMethodsRegistrationCampaignStatedefault  AuthenticationMethodsRegistrationCampaignState = "Default"
	AuthenticationMethodsRegistrationCampaignStatedisabled AuthenticationMethodsRegistrationCampaignState = "Disabled"
	AuthenticationMethodsRegistrationCampaignStateenabled  AuthenticationMethodsRegistrationCampaignState = "Enabled"
)

func PossibleValuesForAuthenticationMethodsRegistrationCampaignState() []string {
	return []string{
		string(AuthenticationMethodsRegistrationCampaignStatedefault),
		string(AuthenticationMethodsRegistrationCampaignStatedisabled),
		string(AuthenticationMethodsRegistrationCampaignStateenabled),
	}
}

func parseAuthenticationMethodsRegistrationCampaignState(input string) (*AuthenticationMethodsRegistrationCampaignState, error) {
	vals := map[string]AuthenticationMethodsRegistrationCampaignState{
		"default":  AuthenticationMethodsRegistrationCampaignStatedefault,
		"disabled": AuthenticationMethodsRegistrationCampaignStatedisabled,
		"enabled":  AuthenticationMethodsRegistrationCampaignStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsRegistrationCampaignState(input)
	return &out, nil
}
