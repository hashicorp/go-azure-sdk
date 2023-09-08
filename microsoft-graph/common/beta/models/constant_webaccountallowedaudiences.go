package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAccountAllowedAudiences string

const (
	WebAccountAllowedAudiencescontacts               WebAccountAllowedAudiences = "Contacts"
	WebAccountAllowedAudienceseveryone               WebAccountAllowedAudiences = "Everyone"
	WebAccountAllowedAudiencesfamily                 WebAccountAllowedAudiences = "Family"
	WebAccountAllowedAudiencesfederatedOrganizations WebAccountAllowedAudiences = "FederatedOrganizations"
	WebAccountAllowedAudiencesgroupMembers           WebAccountAllowedAudiences = "GroupMembers"
	WebAccountAllowedAudiencesme                     WebAccountAllowedAudiences = "Me"
	WebAccountAllowedAudiencesorganization           WebAccountAllowedAudiences = "Organization"
)

func PossibleValuesForWebAccountAllowedAudiences() []string {
	return []string{
		string(WebAccountAllowedAudiencescontacts),
		string(WebAccountAllowedAudienceseveryone),
		string(WebAccountAllowedAudiencesfamily),
		string(WebAccountAllowedAudiencesfederatedOrganizations),
		string(WebAccountAllowedAudiencesgroupMembers),
		string(WebAccountAllowedAudiencesme),
		string(WebAccountAllowedAudiencesorganization),
	}
}

func parseWebAccountAllowedAudiences(input string) (*WebAccountAllowedAudiences, error) {
	vals := map[string]WebAccountAllowedAudiences{
		"contacts":               WebAccountAllowedAudiencescontacts,
		"everyone":               WebAccountAllowedAudienceseveryone,
		"family":                 WebAccountAllowedAudiencesfamily,
		"federatedorganizations": WebAccountAllowedAudiencesfederatedOrganizations,
		"groupmembers":           WebAccountAllowedAudiencesgroupMembers,
		"me":                     WebAccountAllowedAudiencesme,
		"organization":           WebAccountAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebAccountAllowedAudiences(input)
	return &out, nil
}
