package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnotationAllowedAudiences string

const (
	PersonAnnotationAllowedAudiencescontacts               PersonAnnotationAllowedAudiences = "Contacts"
	PersonAnnotationAllowedAudienceseveryone               PersonAnnotationAllowedAudiences = "Everyone"
	PersonAnnotationAllowedAudiencesfamily                 PersonAnnotationAllowedAudiences = "Family"
	PersonAnnotationAllowedAudiencesfederatedOrganizations PersonAnnotationAllowedAudiences = "FederatedOrganizations"
	PersonAnnotationAllowedAudiencesgroupMembers           PersonAnnotationAllowedAudiences = "GroupMembers"
	PersonAnnotationAllowedAudiencesme                     PersonAnnotationAllowedAudiences = "Me"
	PersonAnnotationAllowedAudiencesorganization           PersonAnnotationAllowedAudiences = "Organization"
)

func PossibleValuesForPersonAnnotationAllowedAudiences() []string {
	return []string{
		string(PersonAnnotationAllowedAudiencescontacts),
		string(PersonAnnotationAllowedAudienceseveryone),
		string(PersonAnnotationAllowedAudiencesfamily),
		string(PersonAnnotationAllowedAudiencesfederatedOrganizations),
		string(PersonAnnotationAllowedAudiencesgroupMembers),
		string(PersonAnnotationAllowedAudiencesme),
		string(PersonAnnotationAllowedAudiencesorganization),
	}
}

func parsePersonAnnotationAllowedAudiences(input string) (*PersonAnnotationAllowedAudiences, error) {
	vals := map[string]PersonAnnotationAllowedAudiences{
		"contacts":               PersonAnnotationAllowedAudiencescontacts,
		"everyone":               PersonAnnotationAllowedAudienceseveryone,
		"family":                 PersonAnnotationAllowedAudiencesfamily,
		"federatedorganizations": PersonAnnotationAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonAnnotationAllowedAudiencesgroupMembers,
		"me":                     PersonAnnotationAllowedAudiencesme,
		"organization":           PersonAnnotationAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnotationAllowedAudiences(input)
	return &out, nil
}
