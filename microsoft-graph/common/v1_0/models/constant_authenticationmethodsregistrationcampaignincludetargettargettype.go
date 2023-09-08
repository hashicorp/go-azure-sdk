package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType string

const (
	AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypegroup AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType = "Group"
	AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypeuser  AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType = "User"
)

func PossibleValuesForAuthenticationMethodsRegistrationCampaignIncludeTargetTargetType() []string {
	return []string{
		string(AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypegroup),
		string(AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypeuser),
	}
}

func parseAuthenticationMethodsRegistrationCampaignIncludeTargetTargetType(input string) (*AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType, error) {
	vals := map[string]AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType{
		"group": AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypegroup,
		"user":  AuthenticationMethodsRegistrationCampaignIncludeTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodsRegistrationCampaignIncludeTargetTargetType(input)
	return &out, nil
}
