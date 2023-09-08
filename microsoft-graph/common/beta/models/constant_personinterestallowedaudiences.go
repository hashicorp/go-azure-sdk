package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonInterestAllowedAudiences string

const (
	PersonInterestAllowedAudiencescontacts               PersonInterestAllowedAudiences = "Contacts"
	PersonInterestAllowedAudienceseveryone               PersonInterestAllowedAudiences = "Everyone"
	PersonInterestAllowedAudiencesfamily                 PersonInterestAllowedAudiences = "Family"
	PersonInterestAllowedAudiencesfederatedOrganizations PersonInterestAllowedAudiences = "FederatedOrganizations"
	PersonInterestAllowedAudiencesgroupMembers           PersonInterestAllowedAudiences = "GroupMembers"
	PersonInterestAllowedAudiencesme                     PersonInterestAllowedAudiences = "Me"
	PersonInterestAllowedAudiencesorganization           PersonInterestAllowedAudiences = "Organization"
)

func PossibleValuesForPersonInterestAllowedAudiences() []string {
	return []string{
		string(PersonInterestAllowedAudiencescontacts),
		string(PersonInterestAllowedAudienceseveryone),
		string(PersonInterestAllowedAudiencesfamily),
		string(PersonInterestAllowedAudiencesfederatedOrganizations),
		string(PersonInterestAllowedAudiencesgroupMembers),
		string(PersonInterestAllowedAudiencesme),
		string(PersonInterestAllowedAudiencesorganization),
	}
}

func parsePersonInterestAllowedAudiences(input string) (*PersonInterestAllowedAudiences, error) {
	vals := map[string]PersonInterestAllowedAudiences{
		"contacts":               PersonInterestAllowedAudiencescontacts,
		"everyone":               PersonInterestAllowedAudienceseveryone,
		"family":                 PersonInterestAllowedAudiencesfamily,
		"federatedorganizations": PersonInterestAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonInterestAllowedAudiencesgroupMembers,
		"me":                     PersonInterestAllowedAudiencesme,
		"organization":           PersonInterestAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonInterestAllowedAudiences(input)
	return &out, nil
}
