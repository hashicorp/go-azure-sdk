package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemFacetAllowedAudiences string

const (
	ItemFacetAllowedAudiencescontacts               ItemFacetAllowedAudiences = "Contacts"
	ItemFacetAllowedAudienceseveryone               ItemFacetAllowedAudiences = "Everyone"
	ItemFacetAllowedAudiencesfamily                 ItemFacetAllowedAudiences = "Family"
	ItemFacetAllowedAudiencesfederatedOrganizations ItemFacetAllowedAudiences = "FederatedOrganizations"
	ItemFacetAllowedAudiencesgroupMembers           ItemFacetAllowedAudiences = "GroupMembers"
	ItemFacetAllowedAudiencesme                     ItemFacetAllowedAudiences = "Me"
	ItemFacetAllowedAudiencesorganization           ItemFacetAllowedAudiences = "Organization"
)

func PossibleValuesForItemFacetAllowedAudiences() []string {
	return []string{
		string(ItemFacetAllowedAudiencescontacts),
		string(ItemFacetAllowedAudienceseveryone),
		string(ItemFacetAllowedAudiencesfamily),
		string(ItemFacetAllowedAudiencesfederatedOrganizations),
		string(ItemFacetAllowedAudiencesgroupMembers),
		string(ItemFacetAllowedAudiencesme),
		string(ItemFacetAllowedAudiencesorganization),
	}
}

func parseItemFacetAllowedAudiences(input string) (*ItemFacetAllowedAudiences, error) {
	vals := map[string]ItemFacetAllowedAudiences{
		"contacts":               ItemFacetAllowedAudiencescontacts,
		"everyone":               ItemFacetAllowedAudienceseveryone,
		"family":                 ItemFacetAllowedAudiencesfamily,
		"federatedorganizations": ItemFacetAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ItemFacetAllowedAudiencesgroupMembers,
		"me":                     ItemFacetAllowedAudiencesme,
		"organization":           ItemFacetAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemFacetAllowedAudiences(input)
	return &out, nil
}
