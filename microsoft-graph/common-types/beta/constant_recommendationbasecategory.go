package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendationBaseCategory string

const (
	RecommendationBaseCategory_IdentityBestPractice RecommendationBaseCategory = "identityBestPractice"
	RecommendationBaseCategory_IdentitySecureScore  RecommendationBaseCategory = "identitySecureScore"
)

func PossibleValuesForRecommendationBaseCategory() []string {
	return []string{
		string(RecommendationBaseCategory_IdentityBestPractice),
		string(RecommendationBaseCategory_IdentitySecureScore),
	}
}

func (s *RecommendationBaseCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecommendationBaseCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecommendationBaseCategory(input string) (*RecommendationBaseCategory, error) {
	vals := map[string]RecommendationBaseCategory{
		"identitybestpractice": RecommendationBaseCategory_IdentityBestPractice,
		"identitysecurescore":  RecommendationBaseCategory_IdentitySecureScore,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendationBaseCategory(input)
	return &out, nil
}
