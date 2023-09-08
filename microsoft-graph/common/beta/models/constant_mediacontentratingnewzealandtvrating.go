package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealandTvRating string

const (
	MediaContentRatingNewZealandTvRatingadults           MediaContentRatingNewZealandTvRating = "Adults"
	MediaContentRatingNewZealandTvRatingallAllowed       MediaContentRatingNewZealandTvRating = "AllAllowed"
	MediaContentRatingNewZealandTvRatingallBlocked       MediaContentRatingNewZealandTvRating = "AllBlocked"
	MediaContentRatingNewZealandTvRatinggeneral          MediaContentRatingNewZealandTvRating = "General"
	MediaContentRatingNewZealandTvRatingparentalGuidance MediaContentRatingNewZealandTvRating = "ParentalGuidance"
)

func PossibleValuesForMediaContentRatingNewZealandTvRating() []string {
	return []string{
		string(MediaContentRatingNewZealandTvRatingadults),
		string(MediaContentRatingNewZealandTvRatingallAllowed),
		string(MediaContentRatingNewZealandTvRatingallBlocked),
		string(MediaContentRatingNewZealandTvRatinggeneral),
		string(MediaContentRatingNewZealandTvRatingparentalGuidance),
	}
}

func parseMediaContentRatingNewZealandTvRating(input string) (*MediaContentRatingNewZealandTvRating, error) {
	vals := map[string]MediaContentRatingNewZealandTvRating{
		"adults":           MediaContentRatingNewZealandTvRatingadults,
		"allallowed":       MediaContentRatingNewZealandTvRatingallAllowed,
		"allblocked":       MediaContentRatingNewZealandTvRatingallBlocked,
		"general":          MediaContentRatingNewZealandTvRatinggeneral,
		"parentalguidance": MediaContentRatingNewZealandTvRatingparentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingNewZealandTvRating(input)
	return &out, nil
}
