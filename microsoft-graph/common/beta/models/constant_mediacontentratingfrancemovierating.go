package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingFranceMovieRating string

const (
	MediaContentRatingFranceMovieRatingagesAbove10 MediaContentRatingFranceMovieRating = "AgesAbove10"
	MediaContentRatingFranceMovieRatingagesAbove12 MediaContentRatingFranceMovieRating = "AgesAbove12"
	MediaContentRatingFranceMovieRatingagesAbove16 MediaContentRatingFranceMovieRating = "AgesAbove16"
	MediaContentRatingFranceMovieRatingagesAbove18 MediaContentRatingFranceMovieRating = "AgesAbove18"
	MediaContentRatingFranceMovieRatingallAllowed  MediaContentRatingFranceMovieRating = "AllAllowed"
	MediaContentRatingFranceMovieRatingallBlocked  MediaContentRatingFranceMovieRating = "AllBlocked"
)

func PossibleValuesForMediaContentRatingFranceMovieRating() []string {
	return []string{
		string(MediaContentRatingFranceMovieRatingagesAbove10),
		string(MediaContentRatingFranceMovieRatingagesAbove12),
		string(MediaContentRatingFranceMovieRatingagesAbove16),
		string(MediaContentRatingFranceMovieRatingagesAbove18),
		string(MediaContentRatingFranceMovieRatingallAllowed),
		string(MediaContentRatingFranceMovieRatingallBlocked),
	}
}

func parseMediaContentRatingFranceMovieRating(input string) (*MediaContentRatingFranceMovieRating, error) {
	vals := map[string]MediaContentRatingFranceMovieRating{
		"agesabove10": MediaContentRatingFranceMovieRatingagesAbove10,
		"agesabove12": MediaContentRatingFranceMovieRatingagesAbove12,
		"agesabove16": MediaContentRatingFranceMovieRatingagesAbove16,
		"agesabove18": MediaContentRatingFranceMovieRatingagesAbove18,
		"allallowed":  MediaContentRatingFranceMovieRatingallAllowed,
		"allblocked":  MediaContentRatingFranceMovieRatingallBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingFranceMovieRating(input)
	return &out, nil
}
