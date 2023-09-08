package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingFranceTvRating string

const (
	MediaContentRatingFranceTvRatingagesAbove10 MediaContentRatingFranceTvRating = "AgesAbove10"
	MediaContentRatingFranceTvRatingagesAbove12 MediaContentRatingFranceTvRating = "AgesAbove12"
	MediaContentRatingFranceTvRatingagesAbove16 MediaContentRatingFranceTvRating = "AgesAbove16"
	MediaContentRatingFranceTvRatingagesAbove18 MediaContentRatingFranceTvRating = "AgesAbove18"
	MediaContentRatingFranceTvRatingallAllowed  MediaContentRatingFranceTvRating = "AllAllowed"
	MediaContentRatingFranceTvRatingallBlocked  MediaContentRatingFranceTvRating = "AllBlocked"
)

func PossibleValuesForMediaContentRatingFranceTvRating() []string {
	return []string{
		string(MediaContentRatingFranceTvRatingagesAbove10),
		string(MediaContentRatingFranceTvRatingagesAbove12),
		string(MediaContentRatingFranceTvRatingagesAbove16),
		string(MediaContentRatingFranceTvRatingagesAbove18),
		string(MediaContentRatingFranceTvRatingallAllowed),
		string(MediaContentRatingFranceTvRatingallBlocked),
	}
}

func parseMediaContentRatingFranceTvRating(input string) (*MediaContentRatingFranceTvRating, error) {
	vals := map[string]MediaContentRatingFranceTvRating{
		"agesabove10": MediaContentRatingFranceTvRatingagesAbove10,
		"agesabove12": MediaContentRatingFranceTvRatingagesAbove12,
		"agesabove16": MediaContentRatingFranceTvRatingagesAbove16,
		"agesabove18": MediaContentRatingFranceTvRatingagesAbove18,
		"allallowed":  MediaContentRatingFranceTvRatingallAllowed,
		"allblocked":  MediaContentRatingFranceTvRatingallBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingFranceTvRating(input)
	return &out, nil
}
