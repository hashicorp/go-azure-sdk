package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PersonAnnualEventAllowedAudiences string

const (
	PersonAnnualEventAllowedAudiencescontacts               PersonAnnualEventAllowedAudiences = "Contacts"
	PersonAnnualEventAllowedAudienceseveryone               PersonAnnualEventAllowedAudiences = "Everyone"
	PersonAnnualEventAllowedAudiencesfamily                 PersonAnnualEventAllowedAudiences = "Family"
	PersonAnnualEventAllowedAudiencesfederatedOrganizations PersonAnnualEventAllowedAudiences = "FederatedOrganizations"
	PersonAnnualEventAllowedAudiencesgroupMembers           PersonAnnualEventAllowedAudiences = "GroupMembers"
	PersonAnnualEventAllowedAudiencesme                     PersonAnnualEventAllowedAudiences = "Me"
	PersonAnnualEventAllowedAudiencesorganization           PersonAnnualEventAllowedAudiences = "Organization"
)

func PossibleValuesForPersonAnnualEventAllowedAudiences() []string {
	return []string{
		string(PersonAnnualEventAllowedAudiencescontacts),
		string(PersonAnnualEventAllowedAudienceseveryone),
		string(PersonAnnualEventAllowedAudiencesfamily),
		string(PersonAnnualEventAllowedAudiencesfederatedOrganizations),
		string(PersonAnnualEventAllowedAudiencesgroupMembers),
		string(PersonAnnualEventAllowedAudiencesme),
		string(PersonAnnualEventAllowedAudiencesorganization),
	}
}

func parsePersonAnnualEventAllowedAudiences(input string) (*PersonAnnualEventAllowedAudiences, error) {
	vals := map[string]PersonAnnualEventAllowedAudiences{
		"contacts":               PersonAnnualEventAllowedAudiencescontacts,
		"everyone":               PersonAnnualEventAllowedAudienceseveryone,
		"family":                 PersonAnnualEventAllowedAudiencesfamily,
		"federatedorganizations": PersonAnnualEventAllowedAudiencesfederatedOrganizations,
		"groupmembers":           PersonAnnualEventAllowedAudiencesgroupMembers,
		"me":                     PersonAnnualEventAllowedAudiencesme,
		"organization":           PersonAnnualEventAllowedAudiencesorganization,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PersonAnnualEventAllowedAudiences(input)
	return &out, nil
}
