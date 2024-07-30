package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedPersonRelationship string

const (
	RelatedPersonRelationship_AlternateContact RelatedPersonRelationship = "alternateContact"
	RelatedPersonRelationship_Assistant        RelatedPersonRelationship = "assistant"
	RelatedPersonRelationship_Child            RelatedPersonRelationship = "child"
	RelatedPersonRelationship_Colleague        RelatedPersonRelationship = "colleague"
	RelatedPersonRelationship_DirectReport     RelatedPersonRelationship = "directReport"
	RelatedPersonRelationship_DotLineManager   RelatedPersonRelationship = "dotLineManager"
	RelatedPersonRelationship_DotLineReport    RelatedPersonRelationship = "dotLineReport"
	RelatedPersonRelationship_EmergencyContact RelatedPersonRelationship = "emergencyContact"
	RelatedPersonRelationship_Friend           RelatedPersonRelationship = "friend"
	RelatedPersonRelationship_Manager          RelatedPersonRelationship = "manager"
	RelatedPersonRelationship_Other            RelatedPersonRelationship = "other"
	RelatedPersonRelationship_Parent           RelatedPersonRelationship = "parent"
	RelatedPersonRelationship_Sibling          RelatedPersonRelationship = "sibling"
	RelatedPersonRelationship_Sponsor          RelatedPersonRelationship = "sponsor"
	RelatedPersonRelationship_Spouse           RelatedPersonRelationship = "spouse"
)

func PossibleValuesForRelatedPersonRelationship() []string {
	return []string{
		string(RelatedPersonRelationship_AlternateContact),
		string(RelatedPersonRelationship_Assistant),
		string(RelatedPersonRelationship_Child),
		string(RelatedPersonRelationship_Colleague),
		string(RelatedPersonRelationship_DirectReport),
		string(RelatedPersonRelationship_DotLineManager),
		string(RelatedPersonRelationship_DotLineReport),
		string(RelatedPersonRelationship_EmergencyContact),
		string(RelatedPersonRelationship_Friend),
		string(RelatedPersonRelationship_Manager),
		string(RelatedPersonRelationship_Other),
		string(RelatedPersonRelationship_Parent),
		string(RelatedPersonRelationship_Sibling),
		string(RelatedPersonRelationship_Sponsor),
		string(RelatedPersonRelationship_Spouse),
	}
}

func (s *RelatedPersonRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelatedPersonRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelatedPersonRelationship(input string) (*RelatedPersonRelationship, error) {
	vals := map[string]RelatedPersonRelationship{
		"alternatecontact": RelatedPersonRelationship_AlternateContact,
		"assistant":        RelatedPersonRelationship_Assistant,
		"child":            RelatedPersonRelationship_Child,
		"colleague":        RelatedPersonRelationship_Colleague,
		"directreport":     RelatedPersonRelationship_DirectReport,
		"dotlinemanager":   RelatedPersonRelationship_DotLineManager,
		"dotlinereport":    RelatedPersonRelationship_DotLineReport,
		"emergencycontact": RelatedPersonRelationship_EmergencyContact,
		"friend":           RelatedPersonRelationship_Friend,
		"manager":          RelatedPersonRelationship_Manager,
		"other":            RelatedPersonRelationship_Other,
		"parent":           RelatedPersonRelationship_Parent,
		"sibling":          RelatedPersonRelationship_Sibling,
		"sponsor":          RelatedPersonRelationship_Sponsor,
		"spouse":           RelatedPersonRelationship_Spouse,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelatedPersonRelationship(input)
	return &out, nil
}
