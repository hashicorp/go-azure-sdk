package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemEmailAllowedAudiences string

const (
	ItemEmailAllowedAudiencescontacts               ItemEmailAllowedAudiences = "Contacts"
	ItemEmailAllowedAudienceseveryone               ItemEmailAllowedAudiences = "Everyone"
	ItemEmailAllowedAudiencesfamily                 ItemEmailAllowedAudiences = "Family"
	ItemEmailAllowedAudiencesfederatedOrganizations ItemEmailAllowedAudiences = "FederatedOrganizations"
	ItemEmailAllowedAudiencesgroupMembers           ItemEmailAllowedAudiences = "GroupMembers"
	ItemEmailAllowedAudiencesme                     ItemEmailAllowedAudiences = "Me"
	ItemEmailAllowedAudiencesorganization           ItemEmailAllowedAudiences = "Organization"
)

func PossibleValuesForItemEmailAllowedAudiences() []string {
	return []string{
		string(ItemEmailAllowedAudiencescontacts),
		string(ItemEmailAllowedAudienceseveryone),
		string(ItemEmailAllowedAudiencesfamily),
		string(ItemEmailAllowedAudiencesfederatedOrganizations),
		string(ItemEmailAllowedAudiencesgroupMembers),
		string(ItemEmailAllowedAudiencesme),
		string(ItemEmailAllowedAudiencesorganization),
	}
}

func parseItemEmailAllowedAudiences(input string) (*ItemEmailAllowedAudiences, error) {
	vals := map[string]ItemEmailAllowedAudiences{
		"contacts":               ItemEmailAllowedAudiencescontacts,
		"everyone":               ItemEmailAllowedAudienceseveryone,
		"family":                 ItemEmailAllowedAudiencesfamily,
		"federatedorganizations": ItemEmailAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ItemEmailAllowedAudiencesgroupMembers,
		"me":                     ItemEmailAllowedAudiencesme,
		"organization":           ItemEmailAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemEmailAllowedAudiences(input)
	return &out, nil
}
