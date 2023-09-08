package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingCanadaMovieRating string

const (
	MediaContentRatingCanadaMovieRatingagesAbove14      MediaContentRatingCanadaMovieRating = "AgesAbove14"
	MediaContentRatingCanadaMovieRatingagesAbove18      MediaContentRatingCanadaMovieRating = "AgesAbove18"
	MediaContentRatingCanadaMovieRatingallAllowed       MediaContentRatingCanadaMovieRating = "AllAllowed"
	MediaContentRatingCanadaMovieRatingallBlocked       MediaContentRatingCanadaMovieRating = "AllBlocked"
	MediaContentRatingCanadaMovieRatinggeneral          MediaContentRatingCanadaMovieRating = "General"
	MediaContentRatingCanadaMovieRatingparentalGuidance MediaContentRatingCanadaMovieRating = "ParentalGuidance"
	MediaContentRatingCanadaMovieRatingrestricted       MediaContentRatingCanadaMovieRating = "Restricted"
)

func PossibleValuesForMediaContentRatingCanadaMovieRating() []string {
	return []string{
		string(MediaContentRatingCanadaMovieRatingagesAbove14),
		string(MediaContentRatingCanadaMovieRatingagesAbove18),
		string(MediaContentRatingCanadaMovieRatingallAllowed),
		string(MediaContentRatingCanadaMovieRatingallBlocked),
		string(MediaContentRatingCanadaMovieRatinggeneral),
		string(MediaContentRatingCanadaMovieRatingparentalGuidance),
		string(MediaContentRatingCanadaMovieRatingrestricted),
	}
}

func parseMediaContentRatingCanadaMovieRating(input string) (*MediaContentRatingCanadaMovieRating, error) {
	vals := map[string]MediaContentRatingCanadaMovieRating{
		"agesabove14":      MediaContentRatingCanadaMovieRatingagesAbove14,
		"agesabove18":      MediaContentRatingCanadaMovieRatingagesAbove18,
		"allallowed":       MediaContentRatingCanadaMovieRatingallAllowed,
		"allblocked":       MediaContentRatingCanadaMovieRatingallBlocked,
		"general":          MediaContentRatingCanadaMovieRatinggeneral,
		"parentalguidance": MediaContentRatingCanadaMovieRatingparentalGuidance,
		"restricted":       MediaContentRatingCanadaMovieRatingrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingCanadaMovieRating(input)
	return &out, nil
}
