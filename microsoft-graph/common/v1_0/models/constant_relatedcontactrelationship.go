package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedContactRelationship string

const (
	RelatedContactRelationshipaide     RelatedContactRelationship = "Aide"
	RelatedContactRelationshipchild    RelatedContactRelationship = "Child"
	RelatedContactRelationshipdoctor   RelatedContactRelationship = "Doctor"
	RelatedContactRelationshipguardian RelatedContactRelationship = "Guardian"
	RelatedContactRelationshipother    RelatedContactRelationship = "Other"
	RelatedContactRelationshipparent   RelatedContactRelationship = "Parent"
	RelatedContactRelationshiprelative RelatedContactRelationship = "Relative"
)

func PossibleValuesForRelatedContactRelationship() []string {
	return []string{
		string(RelatedContactRelationshipaide),
		string(RelatedContactRelationshipchild),
		string(RelatedContactRelationshipdoctor),
		string(RelatedContactRelationshipguardian),
		string(RelatedContactRelationshipother),
		string(RelatedContactRelationshipparent),
		string(RelatedContactRelationshiprelative),
	}
}

func parseRelatedContactRelationship(input string) (*RelatedContactRelationship, error) {
	vals := map[string]RelatedContactRelationship{
		"aide":     RelatedContactRelationshipaide,
		"child":    RelatedContactRelationshipchild,
		"doctor":   RelatedContactRelationshipdoctor,
		"guardian": RelatedContactRelationshipguardian,
		"other":    RelatedContactRelationshipother,
		"parent":   RelatedContactRelationshipparent,
		"relative": RelatedContactRelationshiprelative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelatedContactRelationship(input)
	return &out, nil
}
