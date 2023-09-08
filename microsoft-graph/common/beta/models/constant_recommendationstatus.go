package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationStatus string

const (
	RecommendationStatusactive            RecommendationStatus = "Active"
	RecommendationStatuscompletedBySystem RecommendationStatus = "CompletedBySystem"
	RecommendationStatuscompletedByUser   RecommendationStatus = "CompletedByUser"
	RecommendationStatusdismissed         RecommendationStatus = "Dismissed"
	RecommendationStatuspostponed         RecommendationStatus = "Postponed"
)

func PossibleValuesForRecommendationStatus() []string {
	return []string{
		string(RecommendationStatusactive),
		string(RecommendationStatuscompletedBySystem),
		string(RecommendationStatuscompletedByUser),
		string(RecommendationStatusdismissed),
		string(RecommendationStatuspostponed),
	}
}

func parseRecommendationStatus(input string) (*RecommendationStatus, error) {
	vals := map[string]RecommendationStatus{
		"active":            RecommendationStatusactive,
		"completedbysystem": RecommendationStatuscompletedBySystem,
		"completedbyuser":   RecommendationStatuscompletedByUser,
		"dismissed":         RecommendationStatusdismissed,
		"postponed":         RecommendationStatuspostponed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationStatus(input)
	return &out, nil
}
