package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseFeatureAreas string

const (
	RecommendationBaseFeatureAreas_AccessReviews     RecommendationBaseFeatureAreas = "accessReviews"
	RecommendationBaseFeatureAreas_Applications      RecommendationBaseFeatureAreas = "applications"
	RecommendationBaseFeatureAreas_ConditionalAccess RecommendationBaseFeatureAreas = "conditionalAccess"
	RecommendationBaseFeatureAreas_Devices           RecommendationBaseFeatureAreas = "devices"
	RecommendationBaseFeatureAreas_Governance        RecommendationBaseFeatureAreas = "governance"
	RecommendationBaseFeatureAreas_Groups            RecommendationBaseFeatureAreas = "groups"
	RecommendationBaseFeatureAreas_Users             RecommendationBaseFeatureAreas = "users"
)

func PossibleValuesForRecommendationBaseFeatureAreas() []string {
	return []string{
		string(RecommendationBaseFeatureAreas_AccessReviews),
		string(RecommendationBaseFeatureAreas_Applications),
		string(RecommendationBaseFeatureAreas_ConditionalAccess),
		string(RecommendationBaseFeatureAreas_Devices),
		string(RecommendationBaseFeatureAreas_Governance),
		string(RecommendationBaseFeatureAreas_Groups),
		string(RecommendationBaseFeatureAreas_Users),
	}
}

func (s *RecommendationBaseFeatureAreas) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationBaseFeatureAreas(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationBaseFeatureAreas(input string) (*RecommendationBaseFeatureAreas, error) {
	vals := map[string]RecommendationBaseFeatureAreas{
		"accessreviews":     RecommendationBaseFeatureAreas_AccessReviews,
		"applications":      RecommendationBaseFeatureAreas_Applications,
		"conditionalaccess": RecommendationBaseFeatureAreas_ConditionalAccess,
		"devices":           RecommendationBaseFeatureAreas_Devices,
		"governance":        RecommendationBaseFeatureAreas_Governance,
		"groups":            RecommendationBaseFeatureAreas_Groups,
		"users":             RecommendationBaseFeatureAreas_Users,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseFeatureAreas(input)
	return &out, nil
}
