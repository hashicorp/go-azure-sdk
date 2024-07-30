package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelatedContactRelationship string

const (
	RelatedContactRelationship_Aide     RelatedContactRelationship = "aide"
	RelatedContactRelationship_Child    RelatedContactRelationship = "child"
	RelatedContactRelationship_Doctor   RelatedContactRelationship = "doctor"
	RelatedContactRelationship_Guardian RelatedContactRelationship = "guardian"
	RelatedContactRelationship_Other    RelatedContactRelationship = "other"
	RelatedContactRelationship_Parent   RelatedContactRelationship = "parent"
	RelatedContactRelationship_Relative RelatedContactRelationship = "relative"
)

func PossibleValuesForRelatedContactRelationship() []string {
	return []string{
		string(RelatedContactRelationship_Aide),
		string(RelatedContactRelationship_Child),
		string(RelatedContactRelationship_Doctor),
		string(RelatedContactRelationship_Guardian),
		string(RelatedContactRelationship_Other),
		string(RelatedContactRelationship_Parent),
		string(RelatedContactRelationship_Relative),
	}
}

func (s *RelatedContactRelationship) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRelatedContactRelationship(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRelatedContactRelationship(input string) (*RelatedContactRelationship, error) {
	vals := map[string]RelatedContactRelationship{
		"aide":     RelatedContactRelationship_Aide,
		"child":    RelatedContactRelationship_Child,
		"doctor":   RelatedContactRelationship_Doctor,
		"guardian": RelatedContactRelationship_Guardian,
		"other":    RelatedContactRelationship_Other,
		"parent":   RelatedContactRelationship_Parent,
		"relative": RelatedContactRelationship_Relative,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RelatedContactRelationship(input)
	return &out, nil
}
