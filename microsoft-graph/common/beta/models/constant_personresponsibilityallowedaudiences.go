package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonResponsibilityAllowedAudiences string

const (
	PersonResponsibilityAllowedAudiencescontacts               PersonResponsibilityAllowedAudiences = "Contacts"
	PersonResponsibilityAllowedAudienceseveryone               PersonResponsibilityAllowedAudiences = "Everyone"
	PersonResponsibilityAllowedAudiencesfamily                 PersonResponsibilityAllowedAudiences = "Family"
	PersonResponsibilityAllowedAudiencesfederatedOrganizations PersonResponsibilityAllowedAudiences = "FederatedOrganizations"
	PersonResponsibilityAllowedAudiencesgroupMembers           PersonResponsibilityAllowedAudiences = "GroupMembers"
	PersonResponsibilityAllowedAudiencesme                     PersonResponsibilityAllowedAudiences = "Me"
	PersonResponsibilityAllowedAudiencesorganization           PersonResponsibilityAllowedAudiences = "Organization"
)

func PossibleValuesForPersonResponsibilityAllowedAudiences() []string {
	return []string{
		string(PersonResponsibilityAllowedAudiencescontacts),
		string(PersonResponsibilityAllowedAudienceseveryone),
		string(PersonResponsibilityAllowedAudiencesfamily),
		string(PersonResponsibilityAllowedAudiencesfederatedOrganizations),
		string(PersonResponsibilityAllowedAudiencesgroupMembers),
		string(PersonResponsibilityAllowedAudiencesme),
		string(PersonResponsibilityAllowedAudiencesorganization),
	}
}

func parsePersonResponsibilityAllowedAudiences(input string) (*PersonResponsibilityAllowedAudiences, error) {
	vals := map[string]PersonResponsibilityAllowedAudiences{
		"contacts":               PersonResponsibilityAllowedAudiencescontacts,
		"everyone":               PersonResponsibilityAllowedAudienceseveryone,
		"family":                 PersonResponsibilityAllowedAudiencesfamily,
		"federatedorganizations": PersonResponsibilityAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonResponsibilityAllowedAudiencesgroupMembers,
		"me":                     PersonResponsibilityAllowedAudiencesme,
		"organization":           PersonResponsibilityAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonResponsibilityAllowedAudiences(input)
	return &out, nil
}
