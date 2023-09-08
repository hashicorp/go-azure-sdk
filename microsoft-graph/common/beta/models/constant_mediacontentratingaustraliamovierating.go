package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingAustraliaMovieRating string

const (
	MediaContentRatingAustraliaMovieRatingagesAbove15      MediaContentRatingAustraliaMovieRating = "AgesAbove15"
	MediaContentRatingAustraliaMovieRatingagesAbove18      MediaContentRatingAustraliaMovieRating = "AgesAbove18"
	MediaContentRatingAustraliaMovieRatingallAllowed       MediaContentRatingAustraliaMovieRating = "AllAllowed"
	MediaContentRatingAustraliaMovieRatingallBlocked       MediaContentRatingAustraliaMovieRating = "AllBlocked"
	MediaContentRatingAustraliaMovieRatinggeneral          MediaContentRatingAustraliaMovieRating = "General"
	MediaContentRatingAustraliaMovieRatingmature           MediaContentRatingAustraliaMovieRating = "Mature"
	MediaContentRatingAustraliaMovieRatingparentalGuidance MediaContentRatingAustraliaMovieRating = "ParentalGuidance"
)

func PossibleValuesForMediaContentRatingAustraliaMovieRating() []string {
	return []string{
		string(MediaContentRatingAustraliaMovieRatingagesAbove15),
		string(MediaContentRatingAustraliaMovieRatingagesAbove18),
		string(MediaContentRatingAustraliaMovieRatingallAllowed),
		string(MediaContentRatingAustraliaMovieRatingallBlocked),
		string(MediaContentRatingAustraliaMovieRatinggeneral),
		string(MediaContentRatingAustraliaMovieRatingmature),
		string(MediaContentRatingAustraliaMovieRatingparentalGuidance),
	}
}

func parseMediaContentRatingAustraliaMovieRating(input string) (*MediaContentRatingAustraliaMovieRating, error) {
	vals := map[string]MediaContentRatingAustraliaMovieRating{
		"agesabove15":      MediaContentRatingAustraliaMovieRatingagesAbove15,
		"agesabove18":      MediaContentRatingAustraliaMovieRatingagesAbove18,
		"allallowed":       MediaContentRatingAustraliaMovieRatingallAllowed,
		"allblocked":       MediaContentRatingAustraliaMovieRatingallBlocked,
		"general":          MediaContentRatingAustraliaMovieRatinggeneral,
		"mature":           MediaContentRatingAustraliaMovieRatingmature,
		"parentalguidance": MediaContentRatingAustraliaMovieRatingparentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingAustraliaMovieRating(input)
	return &out, nil
}
