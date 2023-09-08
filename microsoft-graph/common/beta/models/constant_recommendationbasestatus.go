package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseStatus string

const (
	RecommendationBaseStatusactive            RecommendationBaseStatus = "Active"
	RecommendationBaseStatuscompletedBySystem RecommendationBaseStatus = "CompletedBySystem"
	RecommendationBaseStatuscompletedByUser   RecommendationBaseStatus = "CompletedByUser"
	RecommendationBaseStatusdismissed         RecommendationBaseStatus = "Dismissed"
	RecommendationBaseStatuspostponed         RecommendationBaseStatus = "Postponed"
)

func PossibleValuesForRecommendationBaseStatus() []string {
	return []string{
		string(RecommendationBaseStatusactive),
		string(RecommendationBaseStatuscompletedBySystem),
		string(RecommendationBaseStatuscompletedByUser),
		string(RecommendationBaseStatusdismissed),
		string(RecommendationBaseStatuspostponed),
	}
}

func parseRecommendationBaseStatus(input string) (*RecommendationBaseStatus, error) {
	vals := map[string]RecommendationBaseStatus{
		"active":            RecommendationBaseStatusactive,
		"completedbysystem": RecommendationBaseStatuscompletedBySystem,
		"completedbyuser":   RecommendationBaseStatuscompletedByUser,
		"dismissed":         RecommendationBaseStatusdismissed,
		"postponed":         RecommendationBaseStatuspostponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseStatus(input)
	return &out, nil
}
