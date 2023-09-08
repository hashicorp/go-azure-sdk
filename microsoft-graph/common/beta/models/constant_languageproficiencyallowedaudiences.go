package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LanguageProficiencyAllowedAudiences string

const (
	LanguageProficiencyAllowedAudiencescontacts               LanguageProficiencyAllowedAudiences = "Contacts"
	LanguageProficiencyAllowedAudienceseveryone               LanguageProficiencyAllowedAudiences = "Everyone"
	LanguageProficiencyAllowedAudiencesfamily                 LanguageProficiencyAllowedAudiences = "Family"
	LanguageProficiencyAllowedAudiencesfederatedOrganizations LanguageProficiencyAllowedAudiences = "FederatedOrganizations"
	LanguageProficiencyAllowedAudiencesgroupMembers           LanguageProficiencyAllowedAudiences = "GroupMembers"
	LanguageProficiencyAllowedAudiencesme                     LanguageProficiencyAllowedAudiences = "Me"
	LanguageProficiencyAllowedAudiencesorganization           LanguageProficiencyAllowedAudiences = "Organization"
)

func PossibleValuesForLanguageProficiencyAllowedAudiences() []string {
	return []string{
		string(LanguageProficiencyAllowedAudiencescontacts),
		string(LanguageProficiencyAllowedAudienceseveryone),
		string(LanguageProficiencyAllowedAudiencesfamily),
		string(LanguageProficiencyAllowedAudiencesfederatedOrganizations),
		string(LanguageProficiencyAllowedAudiencesgroupMembers),
		string(LanguageProficiencyAllowedAudiencesme),
		string(LanguageProficiencyAllowedAudiencesorganization),
	}
}

func parseLanguageProficiencyAllowedAudiences(input string) (*LanguageProficiencyAllowedAudiences, error) {
	vals := map[string]LanguageProficiencyAllowedAudiences{
		"contacts":               LanguageProficiencyAllowedAudiencescontacts,
		"everyone":               LanguageProficiencyAllowedAudienceseveryone,
		"family":                 LanguageProficiencyAllowedAudiencesfamily,
		"federatedorganizations": LanguageProficiencyAllowedAudiencesfederatedOrganizations,
		"groupmembers":           LanguageProficiencyAllowedAudiencesgroupMembers,
		"me":                     LanguageProficiencyAllowedAudiencesme,
		"organization":           LanguageProficiencyAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LanguageProficiencyAllowedAudiences(input)
	return &out, nil
}
