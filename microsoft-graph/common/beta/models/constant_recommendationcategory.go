package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationCategory string

const (
	RecommendationCategoryidentityBestPractice RecommendationCategory = "IdentityBestPractice"
	RecommendationCategoryidentitySecureScore  RecommendationCategory = "IdentitySecureScore"
)

func PossibleValuesForRecommendationCategory() []string {
	return []string{
		string(RecommendationCategoryidentityBestPractice),
		string(RecommendationCategoryidentitySecureScore),
	}
}

func parseRecommendationCategory(input string) (*RecommendationCategory, error) {
	vals := map[string]RecommendationCategory{
		"identitybestpractice": RecommendationCategoryidentityBestPractice,
		"identitysecurescore":  RecommendationCategoryidentitySecureScore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationCategory(input)
	return &out, nil
}
