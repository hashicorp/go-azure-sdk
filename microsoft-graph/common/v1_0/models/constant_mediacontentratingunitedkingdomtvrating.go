package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedKingdomTvRating string

const (
	MediaContentRatingUnitedKingdomTvRatingallAllowed MediaContentRatingUnitedKingdomTvRating = "AllAllowed"
	MediaContentRatingUnitedKingdomTvRatingallBlocked MediaContentRatingUnitedKingdomTvRating = "AllBlocked"
	MediaContentRatingUnitedKingdomTvRatingcaution    MediaContentRatingUnitedKingdomTvRating = "Caution"
)

func PossibleValuesForMediaContentRatingUnitedKingdomTvRating() []string {
	return []string{
		string(MediaContentRatingUnitedKingdomTvRatingallAllowed),
		string(MediaContentRatingUnitedKingdomTvRatingallBlocked),
		string(MediaContentRatingUnitedKingdomTvRatingcaution),
	}
}

func parseMediaContentRatingUnitedKingdomTvRating(input string) (*MediaContentRatingUnitedKingdomTvRating, error) {
	vals := map[string]MediaContentRatingUnitedKingdomTvRating{
		"allallowed": MediaContentRatingUnitedKingdomTvRatingallAllowed,
		"allblocked": MediaContentRatingUnitedKingdomTvRatingallBlocked,
		"caution":    MediaContentRatingUnitedKingdomTvRatingcaution,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedKingdomTvRating(input)
	return &out, nil
}
