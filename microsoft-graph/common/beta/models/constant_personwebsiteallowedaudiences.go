package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonWebsiteAllowedAudiences string

const (
	PersonWebsiteAllowedAudiencescontacts               PersonWebsiteAllowedAudiences = "Contacts"
	PersonWebsiteAllowedAudienceseveryone               PersonWebsiteAllowedAudiences = "Everyone"
	PersonWebsiteAllowedAudiencesfamily                 PersonWebsiteAllowedAudiences = "Family"
	PersonWebsiteAllowedAudiencesfederatedOrganizations PersonWebsiteAllowedAudiences = "FederatedOrganizations"
	PersonWebsiteAllowedAudiencesgroupMembers           PersonWebsiteAllowedAudiences = "GroupMembers"
	PersonWebsiteAllowedAudiencesme                     PersonWebsiteAllowedAudiences = "Me"
	PersonWebsiteAllowedAudiencesorganization           PersonWebsiteAllowedAudiences = "Organization"
)

func PossibleValuesForPersonWebsiteAllowedAudiences() []string {
	return []string{
		string(PersonWebsiteAllowedAudiencescontacts),
		string(PersonWebsiteAllowedAudienceseveryone),
		string(PersonWebsiteAllowedAudiencesfamily),
		string(PersonWebsiteAllowedAudiencesfederatedOrganizations),
		string(PersonWebsiteAllowedAudiencesgroupMembers),
		string(PersonWebsiteAllowedAudiencesme),
		string(PersonWebsiteAllowedAudiencesorganization),
	}
}

func parsePersonWebsiteAllowedAudiences(input string) (*PersonWebsiteAllowedAudiences, error) {
	vals := map[string]PersonWebsiteAllowedAudiences{
		"contacts":               PersonWebsiteAllowedAudiencescontacts,
		"everyone":               PersonWebsiteAllowedAudienceseveryone,
		"family":                 PersonWebsiteAllowedAudiencesfamily,
		"federatedorganizations": PersonWebsiteAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonWebsiteAllowedAudiencesgroupMembers,
		"me":                     PersonWebsiteAllowedAudiencesme,
		"organization":           PersonWebsiteAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonWebsiteAllowedAudiences(input)
	return &out, nil
}
