package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountInformationAllowedAudiences string

const (
	UserAccountInformationAllowedAudiencescontacts               UserAccountInformationAllowedAudiences = "Contacts"
	UserAccountInformationAllowedAudienceseveryone               UserAccountInformationAllowedAudiences = "Everyone"
	UserAccountInformationAllowedAudiencesfamily                 UserAccountInformationAllowedAudiences = "Family"
	UserAccountInformationAllowedAudiencesfederatedOrganizations UserAccountInformationAllowedAudiences = "FederatedOrganizations"
	UserAccountInformationAllowedAudiencesgroupMembers           UserAccountInformationAllowedAudiences = "GroupMembers"
	UserAccountInformationAllowedAudiencesme                     UserAccountInformationAllowedAudiences = "Me"
	UserAccountInformationAllowedAudiencesorganization           UserAccountInformationAllowedAudiences = "Organization"
)

func PossibleValuesForUserAccountInformationAllowedAudiences() []string {
	return []string{
		string(UserAccountInformationAllowedAudiencescontacts),
		string(UserAccountInformationAllowedAudienceseveryone),
		string(UserAccountInformationAllowedAudiencesfamily),
		string(UserAccountInformationAllowedAudiencesfederatedOrganizations),
		string(UserAccountInformationAllowedAudiencesgroupMembers),
		string(UserAccountInformationAllowedAudiencesme),
		string(UserAccountInformationAllowedAudiencesorganization),
	}
}

func parseUserAccountInformationAllowedAudiences(input string) (*UserAccountInformationAllowedAudiences, error) {
	vals := map[string]UserAccountInformationAllowedAudiences{
		"contacts":               UserAccountInformationAllowedAudiencescontacts,
		"everyone":               UserAccountInformationAllowedAudienceseveryone,
		"family":                 UserAccountInformationAllowedAudiencesfamily,
		"federatedorganizations": UserAccountInformationAllowedAudiencesfederatedOrganizations,
		"groupmembers":           UserAccountInformationAllowedAudiencesgroupMembers,
		"me":                     UserAccountInformationAllowedAudiencesme,
		"organization":           UserAccountInformationAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountInformationAllowedAudiences(input)
	return &out, nil
}
