package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPublicationAllowedAudiences string

const (
	ItemPublicationAllowedAudiencescontacts               ItemPublicationAllowedAudiences = "Contacts"
	ItemPublicationAllowedAudienceseveryone               ItemPublicationAllowedAudiences = "Everyone"
	ItemPublicationAllowedAudiencesfamily                 ItemPublicationAllowedAudiences = "Family"
	ItemPublicationAllowedAudiencesfederatedOrganizations ItemPublicationAllowedAudiences = "FederatedOrganizations"
	ItemPublicationAllowedAudiencesgroupMembers           ItemPublicationAllowedAudiences = "GroupMembers"
	ItemPublicationAllowedAudiencesme                     ItemPublicationAllowedAudiences = "Me"
	ItemPublicationAllowedAudiencesorganization           ItemPublicationAllowedAudiences = "Organization"
)

func PossibleValuesForItemPublicationAllowedAudiences() []string {
	return []string{
		string(ItemPublicationAllowedAudiencescontacts),
		string(ItemPublicationAllowedAudienceseveryone),
		string(ItemPublicationAllowedAudiencesfamily),
		string(ItemPublicationAllowedAudiencesfederatedOrganizations),
		string(ItemPublicationAllowedAudiencesgroupMembers),
		string(ItemPublicationAllowedAudiencesme),
		string(ItemPublicationAllowedAudiencesorganization),
	}
}

func parseItemPublicationAllowedAudiences(input string) (*ItemPublicationAllowedAudiences, error) {
	vals := map[string]ItemPublicationAllowedAudiences{
		"contacts":               ItemPublicationAllowedAudiencescontacts,
		"everyone":               ItemPublicationAllowedAudienceseveryone,
		"family":                 ItemPublicationAllowedAudiencesfamily,
		"federatedorganizations": ItemPublicationAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ItemPublicationAllowedAudiencesgroupMembers,
		"me":                     ItemPublicationAllowedAudiencesme,
		"organization":           ItemPublicationAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemPublicationAllowedAudiences(input)
	return &out, nil
}
