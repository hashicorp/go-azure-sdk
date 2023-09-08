package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemAddressAllowedAudiences string

const (
	ItemAddressAllowedAudiencescontacts               ItemAddressAllowedAudiences = "Contacts"
	ItemAddressAllowedAudienceseveryone               ItemAddressAllowedAudiences = "Everyone"
	ItemAddressAllowedAudiencesfamily                 ItemAddressAllowedAudiences = "Family"
	ItemAddressAllowedAudiencesfederatedOrganizations ItemAddressAllowedAudiences = "FederatedOrganizations"
	ItemAddressAllowedAudiencesgroupMembers           ItemAddressAllowedAudiences = "GroupMembers"
	ItemAddressAllowedAudiencesme                     ItemAddressAllowedAudiences = "Me"
	ItemAddressAllowedAudiencesorganization           ItemAddressAllowedAudiences = "Organization"
)

func PossibleValuesForItemAddressAllowedAudiences() []string {
	return []string{
		string(ItemAddressAllowedAudiencescontacts),
		string(ItemAddressAllowedAudienceseveryone),
		string(ItemAddressAllowedAudiencesfamily),
		string(ItemAddressAllowedAudiencesfederatedOrganizations),
		string(ItemAddressAllowedAudiencesgroupMembers),
		string(ItemAddressAllowedAudiencesme),
		string(ItemAddressAllowedAudiencesorganization),
	}
}

func parseItemAddressAllowedAudiences(input string) (*ItemAddressAllowedAudiences, error) {
	vals := map[string]ItemAddressAllowedAudiences{
		"contacts":               ItemAddressAllowedAudiencescontacts,
		"everyone":               ItemAddressAllowedAudienceseveryone,
		"family":                 ItemAddressAllowedAudiencesfamily,
		"federatedorganizations": ItemAddressAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ItemAddressAllowedAudiencesgroupMembers,
		"me":                     ItemAddressAllowedAudiencesme,
		"organization":           ItemAddressAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemAddressAllowedAudiences(input)
	return &out, nil
}
