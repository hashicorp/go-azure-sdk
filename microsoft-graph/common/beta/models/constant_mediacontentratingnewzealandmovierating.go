package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealandMovieRating string

const (
	MediaContentRatingNewZealandMovieRatingagesAbove13           MediaContentRatingNewZealandMovieRating = "AgesAbove13"
	MediaContentRatingNewZealandMovieRatingagesAbove15           MediaContentRatingNewZealandMovieRating = "AgesAbove15"
	MediaContentRatingNewZealandMovieRatingagesAbove16           MediaContentRatingNewZealandMovieRating = "AgesAbove16"
	MediaContentRatingNewZealandMovieRatingagesAbove16Restricted MediaContentRatingNewZealandMovieRating = "AgesAbove16Restricted"
	MediaContentRatingNewZealandMovieRatingagesAbove18           MediaContentRatingNewZealandMovieRating = "AgesAbove18"
	MediaContentRatingNewZealandMovieRatingallAllowed            MediaContentRatingNewZealandMovieRating = "AllAllowed"
	MediaContentRatingNewZealandMovieRatingallBlocked            MediaContentRatingNewZealandMovieRating = "AllBlocked"
	MediaContentRatingNewZealandMovieRatinggeneral               MediaContentRatingNewZealandMovieRating = "General"
	MediaContentRatingNewZealandMovieRatingmature                MediaContentRatingNewZealandMovieRating = "Mature"
	MediaContentRatingNewZealandMovieRatingparentalGuidance      MediaContentRatingNewZealandMovieRating = "ParentalGuidance"
	MediaContentRatingNewZealandMovieRatingrestricted            MediaContentRatingNewZealandMovieRating = "Restricted"
)

func PossibleValuesForMediaContentRatingNewZealandMovieRating() []string {
	return []string{
		string(MediaContentRatingNewZealandMovieRatingagesAbove13),
		string(MediaContentRatingNewZealandMovieRatingagesAbove15),
		string(MediaContentRatingNewZealandMovieRatingagesAbove16),
		string(MediaContentRatingNewZealandMovieRatingagesAbove16Restricted),
		string(MediaContentRatingNewZealandMovieRatingagesAbove18),
		string(MediaContentRatingNewZealandMovieRatingallAllowed),
		string(MediaContentRatingNewZealandMovieRatingallBlocked),
		string(MediaContentRatingNewZealandMovieRatinggeneral),
		string(MediaContentRatingNewZealandMovieRatingmature),
		string(MediaContentRatingNewZealandMovieRatingparentalGuidance),
		string(MediaContentRatingNewZealandMovieRatingrestricted),
	}
}

func parseMediaContentRatingNewZealandMovieRating(input string) (*MediaContentRatingNewZealandMovieRating, error) {
	vals := map[string]MediaContentRatingNewZealandMovieRating{
		"agesabove13":           MediaContentRatingNewZealandMovieRatingagesAbove13,
		"agesabove15":           MediaContentRatingNewZealandMovieRatingagesAbove15,
		"agesabove16":           MediaContentRatingNewZealandMovieRatingagesAbove16,
		"agesabove16restricted": MediaContentRatingNewZealandMovieRatingagesAbove16Restricted,
		"agesabove18":           MediaContentRatingNewZealandMovieRatingagesAbove18,
		"allallowed":            MediaContentRatingNewZealandMovieRatingallAllowed,
		"allblocked":            MediaContentRatingNewZealandMovieRatingallBlocked,
		"general":               MediaContentRatingNewZealandMovieRatinggeneral,
		"mature":                MediaContentRatingNewZealandMovieRatingmature,
		"parentalguidance":      MediaContentRatingNewZealandMovieRatingparentalGuidance,
		"restricted":            MediaContentRatingNewZealandMovieRatingrestricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingNewZealandMovieRating(input)
	return &out, nil
}
