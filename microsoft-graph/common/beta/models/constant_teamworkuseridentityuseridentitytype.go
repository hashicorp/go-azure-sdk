package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkUserIdentityUserIdentityType string

const (
	TeamworkUserIdentityUserIdentityTypeaadUser                        TeamworkUserIdentityUserIdentityType = "AadUser"
	TeamworkUserIdentityUserIdentityTypeanonymousGuest                 TeamworkUserIdentityUserIdentityType = "AnonymousGuest"
	TeamworkUserIdentityUserIdentityTypeazureCommunicationServicesUser TeamworkUserIdentityUserIdentityType = "AzureCommunicationServicesUser"
	TeamworkUserIdentityUserIdentityTypeemailUser                      TeamworkUserIdentityUserIdentityType = "EmailUser"
	TeamworkUserIdentityUserIdentityTypefederatedUser                  TeamworkUserIdentityUserIdentityType = "FederatedUser"
	TeamworkUserIdentityUserIdentityTypeonPremiseAadUser               TeamworkUserIdentityUserIdentityType = "OnPremiseAadUser"
	TeamworkUserIdentityUserIdentityTypepersonalMicrosoftAccountUser   TeamworkUserIdentityUserIdentityType = "PersonalMicrosoftAccountUser"
	TeamworkUserIdentityUserIdentityTypephoneUser                      TeamworkUserIdentityUserIdentityType = "PhoneUser"
	TeamworkUserIdentityUserIdentityTypeskypeUser                      TeamworkUserIdentityUserIdentityType = "SkypeUser"
)

func PossibleValuesForTeamworkUserIdentityUserIdentityType() []string {
	return []string{
		string(TeamworkUserIdentityUserIdentityTypeaadUser),
		string(TeamworkUserIdentityUserIdentityTypeanonymousGuest),
		string(TeamworkUserIdentityUserIdentityTypeazureCommunicationServicesUser),
		string(TeamworkUserIdentityUserIdentityTypeemailUser),
		string(TeamworkUserIdentityUserIdentityTypefederatedUser),
		string(TeamworkUserIdentityUserIdentityTypeonPremiseAadUser),
		string(TeamworkUserIdentityUserIdentityTypepersonalMicrosoftAccountUser),
		string(TeamworkUserIdentityUserIdentityTypephoneUser),
		string(TeamworkUserIdentityUserIdentityTypeskypeUser),
	}
}

func parseTeamworkUserIdentityUserIdentityType(input string) (*TeamworkUserIdentityUserIdentityType, error) {
	vals := map[string]TeamworkUserIdentityUserIdentityType{
		"aaduser":                        TeamworkUserIdentityUserIdentityTypeaadUser,
		"anonymousguest":                 TeamworkUserIdentityUserIdentityTypeanonymousGuest,
		"azurecommunicationservicesuser": TeamworkUserIdentityUserIdentityTypeazureCommunicationServicesUser,
		"emailuser":                      TeamworkUserIdentityUserIdentityTypeemailUser,
		"federateduser":                  TeamworkUserIdentityUserIdentityTypefederatedUser,
		"onpremiseaaduser":               TeamworkUserIdentityUserIdentityTypeonPremiseAadUser,
		"personalmicrosoftaccountuser":   TeamworkUserIdentityUserIdentityTypepersonalMicrosoftAccountUser,
		"phoneuser":                      TeamworkUserIdentityUserIdentityTypephoneUser,
		"skypeuser":                      TeamworkUserIdentityUserIdentityTypeskypeUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkUserIdentityUserIdentityType(input)
	return &out, nil
}
