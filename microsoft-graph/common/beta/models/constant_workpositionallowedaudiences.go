package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkPositionAllowedAudiences string

const (
	WorkPositionAllowedAudiencescontacts               WorkPositionAllowedAudiences = "Contacts"
	WorkPositionAllowedAudienceseveryone               WorkPositionAllowedAudiences = "Everyone"
	WorkPositionAllowedAudiencesfamily                 WorkPositionAllowedAudiences = "Family"
	WorkPositionAllowedAudiencesfederatedOrganizations WorkPositionAllowedAudiences = "FederatedOrganizations"
	WorkPositionAllowedAudiencesgroupMembers           WorkPositionAllowedAudiences = "GroupMembers"
	WorkPositionAllowedAudiencesme                     WorkPositionAllowedAudiences = "Me"
	WorkPositionAllowedAudiencesorganization           WorkPositionAllowedAudiences = "Organization"
)

func PossibleValuesForWorkPositionAllowedAudiences() []string {
	return []string{
		string(WorkPositionAllowedAudiencescontacts),
		string(WorkPositionAllowedAudienceseveryone),
		string(WorkPositionAllowedAudiencesfamily),
		string(WorkPositionAllowedAudiencesfederatedOrganizations),
		string(WorkPositionAllowedAudiencesgroupMembers),
		string(WorkPositionAllowedAudiencesme),
		string(WorkPositionAllowedAudiencesorganization),
	}
}

func parseWorkPositionAllowedAudiences(input string) (*WorkPositionAllowedAudiences, error) {
	vals := map[string]WorkPositionAllowedAudiences{
		"contacts":               WorkPositionAllowedAudiencescontacts,
		"everyone":               WorkPositionAllowedAudienceseveryone,
		"family":                 WorkPositionAllowedAudiencesfamily,
		"federatedorganizations": WorkPositionAllowedAudiencesfederatedOrganizations,
		"groupmembers":           WorkPositionAllowedAudiencesgroupMembers,
		"me":                     WorkPositionAllowedAudiencesme,
		"organization":           WorkPositionAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkPositionAllowedAudiences(input)
	return &out, nil
}
