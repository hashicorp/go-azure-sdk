package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseCategory string

const (
	RecommendationBaseCategoryidentityBestPractice RecommendationBaseCategory = "IdentityBestPractice"
	RecommendationBaseCategoryidentitySecureScore  RecommendationBaseCategory = "IdentitySecureScore"
)

func PossibleValuesForRecommendationBaseCategory() []string {
	return []string{
		string(RecommendationBaseCategoryidentityBestPractice),
		string(RecommendationBaseCategoryidentitySecureScore),
	}
}

func parseRecommendationBaseCategory(input string) (*RecommendationBaseCategory, error) {
	vals := map[string]RecommendationBaseCategory{
		"identitybestpractice": RecommendationBaseCategoryidentityBestPractice,
		"identitysecurescore":  RecommendationBaseCategoryidentitySecureScore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseCategory(input)
	return &out, nil
}
