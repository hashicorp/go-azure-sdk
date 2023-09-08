package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonCertificationAllowedAudiences string

const (
	PersonCertificationAllowedAudiencescontacts               PersonCertificationAllowedAudiences = "Contacts"
	PersonCertificationAllowedAudienceseveryone               PersonCertificationAllowedAudiences = "Everyone"
	PersonCertificationAllowedAudiencesfamily                 PersonCertificationAllowedAudiences = "Family"
	PersonCertificationAllowedAudiencesfederatedOrganizations PersonCertificationAllowedAudiences = "FederatedOrganizations"
	PersonCertificationAllowedAudiencesgroupMembers           PersonCertificationAllowedAudiences = "GroupMembers"
	PersonCertificationAllowedAudiencesme                     PersonCertificationAllowedAudiences = "Me"
	PersonCertificationAllowedAudiencesorganization           PersonCertificationAllowedAudiences = "Organization"
)

func PossibleValuesForPersonCertificationAllowedAudiences() []string {
	return []string{
		string(PersonCertificationAllowedAudiencescontacts),
		string(PersonCertificationAllowedAudienceseveryone),
		string(PersonCertificationAllowedAudiencesfamily),
		string(PersonCertificationAllowedAudiencesfederatedOrganizations),
		string(PersonCertificationAllowedAudiencesgroupMembers),
		string(PersonCertificationAllowedAudiencesme),
		string(PersonCertificationAllowedAudiencesorganization),
	}
}

func parsePersonCertificationAllowedAudiences(input string) (*PersonCertificationAllowedAudiences, error) {
	vals := map[string]PersonCertificationAllowedAudiences{
		"contacts":               PersonCertificationAllowedAudiencescontacts,
		"everyone":               PersonCertificationAllowedAudienceseveryone,
		"family":                 PersonCertificationAllowedAudiencesfamily,
		"federatedorganizations": PersonCertificationAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonCertificationAllowedAudiencesgroupMembers,
		"me":                     PersonCertificationAllowedAudiencesme,
		"organization":           PersonCertificationAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonCertificationAllowedAudiences(input)
	return &out, nil
}
