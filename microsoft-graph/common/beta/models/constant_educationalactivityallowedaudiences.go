package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivityAllowedAudiences string

const (
	EducationalActivityAllowedAudiencescontacts               EducationalActivityAllowedAudiences = "Contacts"
	EducationalActivityAllowedAudienceseveryone               EducationalActivityAllowedAudiences = "Everyone"
	EducationalActivityAllowedAudiencesfamily                 EducationalActivityAllowedAudiences = "Family"
	EducationalActivityAllowedAudiencesfederatedOrganizations EducationalActivityAllowedAudiences = "FederatedOrganizations"
	EducationalActivityAllowedAudiencesgroupMembers           EducationalActivityAllowedAudiences = "GroupMembers"
	EducationalActivityAllowedAudiencesme                     EducationalActivityAllowedAudiences = "Me"
	EducationalActivityAllowedAudiencesorganization           EducationalActivityAllowedAudiences = "Organization"
)

func PossibleValuesForEducationalActivityAllowedAudiences() []string {
	return []string{
		string(EducationalActivityAllowedAudiencescontacts),
		string(EducationalActivityAllowedAudienceseveryone),
		string(EducationalActivityAllowedAudiencesfamily),
		string(EducationalActivityAllowedAudiencesfederatedOrganizations),
		string(EducationalActivityAllowedAudiencesgroupMembers),
		string(EducationalActivityAllowedAudiencesme),
		string(EducationalActivityAllowedAudiencesorganization),
	}
}

func parseEducationalActivityAllowedAudiences(input string) (*EducationalActivityAllowedAudiences, error) {
	vals := map[string]EducationalActivityAllowedAudiences{
		"contacts":               EducationalActivityAllowedAudiencescontacts,
		"everyone":               EducationalActivityAllowedAudienceseveryone,
		"family":                 EducationalActivityAllowedAudiencesfamily,
		"federatedorganizations": EducationalActivityAllowedAudiencesfederatedOrganizations,
		"groupmembers":           EducationalActivityAllowedAudiencesgroupMembers,
		"me":                     EducationalActivityAllowedAudiencesme,
		"organization":           EducationalActivityAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EducationalActivityAllowedAudiences(input)
	return &out, nil
}
