package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkillProficiencyAllowedAudiences string

const (
	SkillProficiencyAllowedAudiencescontacts               SkillProficiencyAllowedAudiences = "Contacts"
	SkillProficiencyAllowedAudienceseveryone               SkillProficiencyAllowedAudiences = "Everyone"
	SkillProficiencyAllowedAudiencesfamily                 SkillProficiencyAllowedAudiences = "Family"
	SkillProficiencyAllowedAudiencesfederatedOrganizations SkillProficiencyAllowedAudiences = "FederatedOrganizations"
	SkillProficiencyAllowedAudiencesgroupMembers           SkillProficiencyAllowedAudiences = "GroupMembers"
	SkillProficiencyAllowedAudiencesme                     SkillProficiencyAllowedAudiences = "Me"
	SkillProficiencyAllowedAudiencesorganization           SkillProficiencyAllowedAudiences = "Organization"
)

func PossibleValuesForSkillProficiencyAllowedAudiences() []string {
	return []string{
		string(SkillProficiencyAllowedAudiencescontacts),
		string(SkillProficiencyAllowedAudienceseveryone),
		string(SkillProficiencyAllowedAudiencesfamily),
		string(SkillProficiencyAllowedAudiencesfederatedOrganizations),
		string(SkillProficiencyAllowedAudiencesgroupMembers),
		string(SkillProficiencyAllowedAudiencesme),
		string(SkillProficiencyAllowedAudiencesorganization),
	}
}

func parseSkillProficiencyAllowedAudiences(input string) (*SkillProficiencyAllowedAudiences, error) {
	vals := map[string]SkillProficiencyAllowedAudiences{
		"contacts":               SkillProficiencyAllowedAudiencescontacts,
		"everyone":               SkillProficiencyAllowedAudienceseveryone,
		"family":                 SkillProficiencyAllowedAudiencesfamily,
		"federatedorganizations": SkillProficiencyAllowedAudiencesfederatedOrganizations,
		"groupmembers":           SkillProficiencyAllowedAudiencesgroupMembers,
		"me":                     SkillProficiencyAllowedAudiencesme,
		"organization":           SkillProficiencyAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkillProficiencyAllowedAudiences(input)
	return &out, nil
}
