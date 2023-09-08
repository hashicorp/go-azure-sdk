package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonNameAllowedAudiences string

const (
	PersonNameAllowedAudiencescontacts               PersonNameAllowedAudiences = "Contacts"
	PersonNameAllowedAudienceseveryone               PersonNameAllowedAudiences = "Everyone"
	PersonNameAllowedAudiencesfamily                 PersonNameAllowedAudiences = "Family"
	PersonNameAllowedAudiencesfederatedOrganizations PersonNameAllowedAudiences = "FederatedOrganizations"
	PersonNameAllowedAudiencesgroupMembers           PersonNameAllowedAudiences = "GroupMembers"
	PersonNameAllowedAudiencesme                     PersonNameAllowedAudiences = "Me"
	PersonNameAllowedAudiencesorganization           PersonNameAllowedAudiences = "Organization"
)

func PossibleValuesForPersonNameAllowedAudiences() []string {
	return []string{
		string(PersonNameAllowedAudiencescontacts),
		string(PersonNameAllowedAudienceseveryone),
		string(PersonNameAllowedAudiencesfamily),
		string(PersonNameAllowedAudiencesfederatedOrganizations),
		string(PersonNameAllowedAudiencesgroupMembers),
		string(PersonNameAllowedAudiencesme),
		string(PersonNameAllowedAudiencesorganization),
	}
}

func parsePersonNameAllowedAudiences(input string) (*PersonNameAllowedAudiences, error) {
	vals := map[string]PersonNameAllowedAudiences{
		"contacts":               PersonNameAllowedAudiencescontacts,
		"everyone":               PersonNameAllowedAudienceseveryone,
		"family":                 PersonNameAllowedAudiencesfamily,
		"federatedorganizations": PersonNameAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonNameAllowedAudiencesgroupMembers,
		"me":                     PersonNameAllowedAudiencesme,
		"organization":           PersonNameAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonNameAllowedAudiences(input)
	return &out, nil
}
