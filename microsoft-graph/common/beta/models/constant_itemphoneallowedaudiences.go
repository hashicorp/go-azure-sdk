package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPhoneAllowedAudiences string

const (
	ItemPhoneAllowedAudiencescontacts               ItemPhoneAllowedAudiences = "Contacts"
	ItemPhoneAllowedAudienceseveryone               ItemPhoneAllowedAudiences = "Everyone"
	ItemPhoneAllowedAudiencesfamily                 ItemPhoneAllowedAudiences = "Family"
	ItemPhoneAllowedAudiencesfederatedOrganizations ItemPhoneAllowedAudiences = "FederatedOrganizations"
	ItemPhoneAllowedAudiencesgroupMembers           ItemPhoneAllowedAudiences = "GroupMembers"
	ItemPhoneAllowedAudiencesme                     ItemPhoneAllowedAudiences = "Me"
	ItemPhoneAllowedAudiencesorganization           ItemPhoneAllowedAudiences = "Organization"
)

func PossibleValuesForItemPhoneAllowedAudiences() []string {
	return []string{
		string(ItemPhoneAllowedAudiencescontacts),
		string(ItemPhoneAllowedAudienceseveryone),
		string(ItemPhoneAllowedAudiencesfamily),
		string(ItemPhoneAllowedAudiencesfederatedOrganizations),
		string(ItemPhoneAllowedAudiencesgroupMembers),
		string(ItemPhoneAllowedAudiencesme),
		string(ItemPhoneAllowedAudiencesorganization),
	}
}

func parseItemPhoneAllowedAudiences(input string) (*ItemPhoneAllowedAudiences, error) {
	vals := map[string]ItemPhoneAllowedAudiences{
		"contacts":               ItemPhoneAllowedAudiencescontacts,
		"everyone":               ItemPhoneAllowedAudienceseveryone,
		"family":                 ItemPhoneAllowedAudiencesfamily,
		"federatedorganizations": ItemPhoneAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ItemPhoneAllowedAudiencesgroupMembers,
		"me":                     ItemPhoneAllowedAudiencesme,
		"organization":           ItemPhoneAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPhoneAllowedAudiences(input)
	return &out, nil
}
