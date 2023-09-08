package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapanMovieRating string

const (
	MediaContentRatingJapanMovieRatingagesAbove15      MediaContentRatingJapanMovieRating = "AgesAbove15"
	MediaContentRatingJapanMovieRatingagesAbove18      MediaContentRatingJapanMovieRating = "AgesAbove18"
	MediaContentRatingJapanMovieRatingallAllowed       MediaContentRatingJapanMovieRating = "AllAllowed"
	MediaContentRatingJapanMovieRatingallBlocked       MediaContentRatingJapanMovieRating = "AllBlocked"
	MediaContentRatingJapanMovieRatinggeneral          MediaContentRatingJapanMovieRating = "General"
	MediaContentRatingJapanMovieRatingparentalGuidance MediaContentRatingJapanMovieRating = "ParentalGuidance"
)

func PossibleValuesForMediaContentRatingJapanMovieRating() []string {
	return []string{
		string(MediaContentRatingJapanMovieRatingagesAbove15),
		string(MediaContentRatingJapanMovieRatingagesAbove18),
		string(MediaContentRatingJapanMovieRatingallAllowed),
		string(MediaContentRatingJapanMovieRatingallBlocked),
		string(MediaContentRatingJapanMovieRatinggeneral),
		string(MediaContentRatingJapanMovieRatingparentalGuidance),
	}
}

func parseMediaContentRatingJapanMovieRating(input string) (*MediaContentRatingJapanMovieRating, error) {
	vals := map[string]MediaContentRatingJapanMovieRating{
		"agesabove15":      MediaContentRatingJapanMovieRatingagesAbove15,
		"agesabove18":      MediaContentRatingJapanMovieRatingagesAbove18,
		"allallowed":       MediaContentRatingJapanMovieRatingallAllowed,
		"allblocked":       MediaContentRatingJapanMovieRatingallBlocked,
		"general":          MediaContentRatingJapanMovieRatinggeneral,
		"parentalguidance": MediaContentRatingJapanMovieRatingparentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingJapanMovieRating(input)
	return &out, nil
}
