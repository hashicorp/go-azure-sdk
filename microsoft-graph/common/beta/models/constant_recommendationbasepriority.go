package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBasePriority string

const (
	RecommendationBasePriorityhigh   RecommendationBasePriority = "High"
	RecommendationBasePrioritylow    RecommendationBasePriority = "Low"
	RecommendationBasePrioritymedium RecommendationBasePriority = "Medium"
)

func PossibleValuesForRecommendationBasePriority() []string {
	return []string{
		string(RecommendationBasePriorityhigh),
		string(RecommendationBasePrioritylow),
		string(RecommendationBasePrioritymedium),
	}
}

func parseRecommendationBasePriority(input string) (*RecommendationBasePriority, error) {
	vals := map[string]RecommendationBasePriority{
		"high":   RecommendationBasePriorityhigh,
		"low":    RecommendationBasePrioritylow,
		"medium": RecommendationBasePrioritymedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBasePriority(input)
	return &out, nil
}
