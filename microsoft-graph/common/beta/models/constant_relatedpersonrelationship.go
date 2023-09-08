package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedPersonRelationship string

const (
	RelatedPersonRelationshipalternateContact RelatedPersonRelationship = "AlternateContact"
	RelatedPersonRelationshipassistant        RelatedPersonRelationship = "Assistant"
	RelatedPersonRelationshipchild            RelatedPersonRelationship = "Child"
	RelatedPersonRelationshipcolleague        RelatedPersonRelationship = "Colleague"
	RelatedPersonRelationshipdirectReport     RelatedPersonRelationship = "DirectReport"
	RelatedPersonRelationshipdotLineManager   RelatedPersonRelationship = "DotLineManager"
	RelatedPersonRelationshipdotLineReport    RelatedPersonRelationship = "DotLineReport"
	RelatedPersonRelationshipemergencyContact RelatedPersonRelationship = "EmergencyContact"
	RelatedPersonRelationshipfriend           RelatedPersonRelationship = "Friend"
	RelatedPersonRelationshipmanager          RelatedPersonRelationship = "Manager"
	RelatedPersonRelationshipother            RelatedPersonRelationship = "Other"
	RelatedPersonRelationshipparent           RelatedPersonRelationship = "Parent"
	RelatedPersonRelationshipsibling          RelatedPersonRelationship = "Sibling"
	RelatedPersonRelationshipsponsor          RelatedPersonRelationship = "Sponsor"
	RelatedPersonRelationshipspouse           RelatedPersonRelationship = "Spouse"
)

func PossibleValuesForRelatedPersonRelationship() []string {
	return []string{
		string(RelatedPersonRelationshipalternateContact),
		string(RelatedPersonRelationshipassistant),
		string(RelatedPersonRelationshipchild),
		string(RelatedPersonRelationshipcolleague),
		string(RelatedPersonRelationshipdirectReport),
		string(RelatedPersonRelationshipdotLineManager),
		string(RelatedPersonRelationshipdotLineReport),
		string(RelatedPersonRelationshipemergencyContact),
		string(RelatedPersonRelationshipfriend),
		string(RelatedPersonRelationshipmanager),
		string(RelatedPersonRelationshipother),
		string(RelatedPersonRelationshipparent),
		string(RelatedPersonRelationshipsibling),
		string(RelatedPersonRelationshipsponsor),
		string(RelatedPersonRelationshipspouse),
	}
}

func parseRelatedPersonRelationship(input string) (*RelatedPersonRelationship, error) {
	vals := map[string]RelatedPersonRelationship{
		"alternatecontact": RelatedPersonRelationshipalternateContact,
		"assistant":        RelatedPersonRelationshipassistant,
		"child":            RelatedPersonRelationshipchild,
		"colleague":        RelatedPersonRelationshipcolleague,
		"directreport":     RelatedPersonRelationshipdirectReport,
		"dotlinemanager":   RelatedPersonRelationshipdotLineManager,
		"dotlinereport":    RelatedPersonRelationshipdotLineReport,
		"emergencycontact": RelatedPersonRelationshipemergencyContact,
		"friend":           RelatedPersonRelationshipfriend,
		"manager":          RelatedPersonRelationshipmanager,
		"other":            RelatedPersonRelationshipother,
		"parent":           RelatedPersonRelationshipparent,
		"sibling":          RelatedPersonRelationshipsibling,
		"sponsor":          RelatedPersonRelationshipsponsor,
		"spouse":           RelatedPersonRelationshipspouse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelatedPersonRelationship(input)
	return &out, nil
}
