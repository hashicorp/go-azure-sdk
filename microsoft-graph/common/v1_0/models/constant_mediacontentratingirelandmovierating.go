package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIrelandMovieRating string

const (
	MediaContentRatingIrelandMovieRatingadults           MediaContentRatingIrelandMovieRating = "Adults"
	MediaContentRatingIrelandMovieRatingagesAbove12      MediaContentRatingIrelandMovieRating = "AgesAbove12"
	MediaContentRatingIrelandMovieRatingagesAbove15      MediaContentRatingIrelandMovieRating = "AgesAbove15"
	MediaContentRatingIrelandMovieRatingagesAbove16      MediaContentRatingIrelandMovieRating = "AgesAbove16"
	MediaContentRatingIrelandMovieRatingallAllowed       MediaContentRatingIrelandMovieRating = "AllAllowed"
	MediaContentRatingIrelandMovieRatingallBlocked       MediaContentRatingIrelandMovieRating = "AllBlocked"
	MediaContentRatingIrelandMovieRatinggeneral          MediaContentRatingIrelandMovieRating = "General"
	MediaContentRatingIrelandMovieRatingparentalGuidance MediaContentRatingIrelandMovieRating = "ParentalGuidance"
)

func PossibleValuesForMediaContentRatingIrelandMovieRating() []string {
	return []string{
		string(MediaContentRatingIrelandMovieRatingadults),
		string(MediaContentRatingIrelandMovieRatingagesAbove12),
		string(MediaContentRatingIrelandMovieRatingagesAbove15),
		string(MediaContentRatingIrelandMovieRatingagesAbove16),
		string(MediaContentRatingIrelandMovieRatingallAllowed),
		string(MediaContentRatingIrelandMovieRatingallBlocked),
		string(MediaContentRatingIrelandMovieRatinggeneral),
		string(MediaContentRatingIrelandMovieRatingparentalGuidance),
	}
}

func parseMediaContentRatingIrelandMovieRating(input string) (*MediaContentRatingIrelandMovieRating, error) {
	vals := map[string]MediaContentRatingIrelandMovieRating{
		"adults":           MediaContentRatingIrelandMovieRatingadults,
		"agesabove12":      MediaContentRatingIrelandMovieRatingagesAbove12,
		"agesabove15":      MediaContentRatingIrelandMovieRatingagesAbove15,
		"agesabove16":      MediaContentRatingIrelandMovieRatingagesAbove16,
		"allallowed":       MediaContentRatingIrelandMovieRatingallAllowed,
		"allblocked":       MediaContentRatingIrelandMovieRatingallBlocked,
		"general":          MediaContentRatingIrelandMovieRatinggeneral,
		"parentalguidance": MediaContentRatingIrelandMovieRatingparentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingIrelandMovieRating(input)
	return &out, nil
}
