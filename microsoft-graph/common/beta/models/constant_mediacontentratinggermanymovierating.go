package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingGermanyMovieRating string

const (
	MediaContentRatingGermanyMovieRatingadults      MediaContentRatingGermanyMovieRating = "Adults"
	MediaContentRatingGermanyMovieRatingagesAbove12 MediaContentRatingGermanyMovieRating = "AgesAbove12"
	MediaContentRatingGermanyMovieRatingagesAbove16 MediaContentRatingGermanyMovieRating = "AgesAbove16"
	MediaContentRatingGermanyMovieRatingagesAbove6  MediaContentRatingGermanyMovieRating = "AgesAbove6"
	MediaContentRatingGermanyMovieRatingallAllowed  MediaContentRatingGermanyMovieRating = "AllAllowed"
	MediaContentRatingGermanyMovieRatingallBlocked  MediaContentRatingGermanyMovieRating = "AllBlocked"
	MediaContentRatingGermanyMovieRatinggeneral     MediaContentRatingGermanyMovieRating = "General"
)

func PossibleValuesForMediaContentRatingGermanyMovieRating() []string {
	return []string{
		string(MediaContentRatingGermanyMovieRatingadults),
		string(MediaContentRatingGermanyMovieRatingagesAbove12),
		string(MediaContentRatingGermanyMovieRatingagesAbove16),
		string(MediaContentRatingGermanyMovieRatingagesAbove6),
		string(MediaContentRatingGermanyMovieRatingallAllowed),
		string(MediaContentRatingGermanyMovieRatingallBlocked),
		string(MediaContentRatingGermanyMovieRatinggeneral),
	}
}

func parseMediaContentRatingGermanyMovieRating(input string) (*MediaContentRatingGermanyMovieRating, error) {
	vals := map[string]MediaContentRatingGermanyMovieRating{
		"adults":      MediaContentRatingGermanyMovieRatingadults,
		"agesabove12": MediaContentRatingGermanyMovieRatingagesAbove12,
		"agesabove16": MediaContentRatingGermanyMovieRatingagesAbove16,
		"agesabove6":  MediaContentRatingGermanyMovieRatingagesAbove6,
		"allallowed":  MediaContentRatingGermanyMovieRatingallAllowed,
		"allblocked":  MediaContentRatingGermanyMovieRatingallBlocked,
		"general":     MediaContentRatingGermanyMovieRatinggeneral,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingGermanyMovieRating(input)
	return &out, nil
}
