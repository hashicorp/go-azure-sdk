package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedKingdomTvRating string

const (
	MediaContentRatingUnitedKingdomTvRating_AllAllowed MediaContentRatingUnitedKingdomTvRating = "allAllowed"
	MediaContentRatingUnitedKingdomTvRating_AllBlocked MediaContentRatingUnitedKingdomTvRating = "allBlocked"
	MediaContentRatingUnitedKingdomTvRating_Caution    MediaContentRatingUnitedKingdomTvRating = "caution"
)

func PossibleValuesForMediaContentRatingUnitedKingdomTvRating() []string {
	return []string{
		string(MediaContentRatingUnitedKingdomTvRating_AllAllowed),
		string(MediaContentRatingUnitedKingdomTvRating_AllBlocked),
		string(MediaContentRatingUnitedKingdomTvRating_Caution),
	}
}

func (s *MediaContentRatingUnitedKingdomTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingUnitedKingdomTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingUnitedKingdomTvRating(input string) (*MediaContentRatingUnitedKingdomTvRating, error) {
	vals := map[string]MediaContentRatingUnitedKingdomTvRating{
		"allallowed": MediaContentRatingUnitedKingdomTvRating_AllAllowed,
		"allblocked": MediaContentRatingUnitedKingdomTvRating_AllBlocked,
		"caution":    MediaContentRatingUnitedKingdomTvRating_Caution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedKingdomTvRating(input)
	return &out, nil
}
