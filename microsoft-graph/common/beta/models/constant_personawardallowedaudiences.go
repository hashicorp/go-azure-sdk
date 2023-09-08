package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAwardAllowedAudiences string

const (
	PersonAwardAllowedAudiencescontacts               PersonAwardAllowedAudiences = "Contacts"
	PersonAwardAllowedAudienceseveryone               PersonAwardAllowedAudiences = "Everyone"
	PersonAwardAllowedAudiencesfamily                 PersonAwardAllowedAudiences = "Family"
	PersonAwardAllowedAudiencesfederatedOrganizations PersonAwardAllowedAudiences = "FederatedOrganizations"
	PersonAwardAllowedAudiencesgroupMembers           PersonAwardAllowedAudiences = "GroupMembers"
	PersonAwardAllowedAudiencesme                     PersonAwardAllowedAudiences = "Me"
	PersonAwardAllowedAudiencesorganization           PersonAwardAllowedAudiences = "Organization"
)

func PossibleValuesForPersonAwardAllowedAudiences() []string {
	return []string{
		string(PersonAwardAllowedAudiencescontacts),
		string(PersonAwardAllowedAudienceseveryone),
		string(PersonAwardAllowedAudiencesfamily),
		string(PersonAwardAllowedAudiencesfederatedOrganizations),
		string(PersonAwardAllowedAudiencesgroupMembers),
		string(PersonAwardAllowedAudiencesme),
		string(PersonAwardAllowedAudiencesorganization),
	}
}

func parsePersonAwardAllowedAudiences(input string) (*PersonAwardAllowedAudiences, error) {
	vals := map[string]PersonAwardAllowedAudiences{
		"contacts":               PersonAwardAllowedAudiencescontacts,
		"everyone":               PersonAwardAllowedAudienceseveryone,
		"family":                 PersonAwardAllowedAudiencesfamily,
		"federatedorganizations": PersonAwardAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonAwardAllowedAudiencesgroupMembers,
		"me":                     PersonAwardAllowedAudiencesme,
		"organization":           PersonAwardAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAwardAllowedAudiences(input)
	return &out, nil
}
