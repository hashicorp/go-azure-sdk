package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectParticipationAllowedAudiences string

const (
	ProjectParticipationAllowedAudiencescontacts               ProjectParticipationAllowedAudiences = "Contacts"
	ProjectParticipationAllowedAudienceseveryone               ProjectParticipationAllowedAudiences = "Everyone"
	ProjectParticipationAllowedAudiencesfamily                 ProjectParticipationAllowedAudiences = "Family"
	ProjectParticipationAllowedAudiencesfederatedOrganizations ProjectParticipationAllowedAudiences = "FederatedOrganizations"
	ProjectParticipationAllowedAudiencesgroupMembers           ProjectParticipationAllowedAudiences = "GroupMembers"
	ProjectParticipationAllowedAudiencesme                     ProjectParticipationAllowedAudiences = "Me"
	ProjectParticipationAllowedAudiencesorganization           ProjectParticipationAllowedAudiences = "Organization"
)

func PossibleValuesForProjectParticipationAllowedAudiences() []string {
	return []string{
		string(ProjectParticipationAllowedAudiencescontacts),
		string(ProjectParticipationAllowedAudienceseveryone),
		string(ProjectParticipationAllowedAudiencesfamily),
		string(ProjectParticipationAllowedAudiencesfederatedOrganizations),
		string(ProjectParticipationAllowedAudiencesgroupMembers),
		string(ProjectParticipationAllowedAudiencesme),
		string(ProjectParticipationAllowedAudiencesorganization),
	}
}

func parseProjectParticipationAllowedAudiences(input string) (*ProjectParticipationAllowedAudiences, error) {
	vals := map[string]ProjectParticipationAllowedAudiences{
		"contacts":               ProjectParticipationAllowedAudiencescontacts,
		"everyone":               ProjectParticipationAllowedAudienceseveryone,
		"family":                 ProjectParticipationAllowedAudiencesfamily,
		"federatedorganizations": ProjectParticipationAllowedAudiencesfederatedOrganizations,
		"groupmembers":           ProjectParticipationAllowedAudiencesgroupMembers,
		"me":                     ProjectParticipationAllowedAudiencesme,
		"organization":           ProjectParticipationAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProjectParticipationAllowedAudiences(input)
	return &out, nil
}
