package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseStatus string

const (
	RecommendationBaseStatus_Active            RecommendationBaseStatus = "active"
	RecommendationBaseStatus_CompletedBySystem RecommendationBaseStatus = "completedBySystem"
	RecommendationBaseStatus_CompletedByUser   RecommendationBaseStatus = "completedByUser"
	RecommendationBaseStatus_Dismissed         RecommendationBaseStatus = "dismissed"
	RecommendationBaseStatus_Postponed         RecommendationBaseStatus = "postponed"
)

func PossibleValuesForRecommendationBaseStatus() []string {
	return []string{
		string(RecommendationBaseStatus_Active),
		string(RecommendationBaseStatus_CompletedBySystem),
		string(RecommendationBaseStatus_CompletedByUser),
		string(RecommendationBaseStatus_Dismissed),
		string(RecommendationBaseStatus_Postponed),
	}
}

func (s *RecommendationBaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationBaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationBaseStatus(input string) (*RecommendationBaseStatus, error) {
	vals := map[string]RecommendationBaseStatus{
		"active":            RecommendationBaseStatus_Active,
		"completedbysystem": RecommendationBaseStatus_CompletedBySystem,
		"completedbyuser":   RecommendationBaseStatus_CompletedByUser,
		"dismissed":         RecommendationBaseStatus_Dismissed,
		"postponed":         RecommendationBaseStatus_Postponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseStatus(input)
	return &out, nil
}
