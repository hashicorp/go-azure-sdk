package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBasePriority string

const (
	RecommendationBasePriority_High   RecommendationBasePriority = "high"
	RecommendationBasePriority_Low    RecommendationBasePriority = "low"
	RecommendationBasePriority_Medium RecommendationBasePriority = "medium"
)

func PossibleValuesForRecommendationBasePriority() []string {
	return []string{
		string(RecommendationBasePriority_High),
		string(RecommendationBasePriority_Low),
		string(RecommendationBasePriority_Medium),
	}
}

func (s *RecommendationBasePriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationBasePriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationBasePriority(input string) (*RecommendationBasePriority, error) {
	vals := map[string]RecommendationBasePriority{
		"high":   RecommendationBasePriority_High,
		"low":    RecommendationBasePriority_Low,
		"medium": RecommendationBasePriority_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBasePriority(input)
	return &out, nil
}
