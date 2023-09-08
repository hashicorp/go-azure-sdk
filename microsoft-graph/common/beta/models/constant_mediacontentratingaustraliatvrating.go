package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingAustraliaTvRating string

const (
	MediaContentRatingAustraliaTvRatingagesAbove15              MediaContentRatingAustraliaTvRating = "AgesAbove15"
	MediaContentRatingAustraliaTvRatingagesAbove15AdultViolence MediaContentRatingAustraliaTvRating = "AgesAbove15AdultViolence"
	MediaContentRatingAustraliaTvRatingallAllowed               MediaContentRatingAustraliaTvRating = "AllAllowed"
	MediaContentRatingAustraliaTvRatingallBlocked               MediaContentRatingAustraliaTvRating = "AllBlocked"
	MediaContentRatingAustraliaTvRatingchildren                 MediaContentRatingAustraliaTvRating = "Children"
	MediaContentRatingAustraliaTvRatinggeneral                  MediaContentRatingAustraliaTvRating = "General"
	MediaContentRatingAustraliaTvRatingmature                   MediaContentRatingAustraliaTvRating = "Mature"
	MediaContentRatingAustraliaTvRatingparentalGuidance         MediaContentRatingAustraliaTvRating = "ParentalGuidance"
	MediaContentRatingAustraliaTvRatingpreschoolers             MediaContentRatingAustraliaTvRating = "Preschoolers"
)

func PossibleValuesForMediaContentRatingAustraliaTvRating() []string {
	return []string{
		string(MediaContentRatingAustraliaTvRatingagesAbove15),
		string(MediaContentRatingAustraliaTvRatingagesAbove15AdultViolence),
		string(MediaContentRatingAustraliaTvRatingallAllowed),
		string(MediaContentRatingAustraliaTvRatingallBlocked),
		string(MediaContentRatingAustraliaTvRatingchildren),
		string(MediaContentRatingAustraliaTvRatinggeneral),
		string(MediaContentRatingAustraliaTvRatingmature),
		string(MediaContentRatingAustraliaTvRatingparentalGuidance),
		string(MediaContentRatingAustraliaTvRatingpreschoolers),
	}
}

func parseMediaContentRatingAustraliaTvRating(input string) (*MediaContentRatingAustraliaTvRating, error) {
	vals := map[string]MediaContentRatingAustraliaTvRating{
		"agesabove15":              MediaContentRatingAustraliaTvRatingagesAbove15,
		"agesabove15adultviolence": MediaContentRatingAustraliaTvRatingagesAbove15AdultViolence,
		"allallowed":               MediaContentRatingAustraliaTvRatingallAllowed,
		"allblocked":               MediaContentRatingAustraliaTvRatingallBlocked,
		"children":                 MediaContentRatingAustraliaTvRatingchildren,
		"general":                  MediaContentRatingAustraliaTvRatinggeneral,
		"mature":                   MediaContentRatingAustraliaTvRatingmature,
		"parentalguidance":         MediaContentRatingAustraliaTvRatingparentalGuidance,
		"preschoolers":             MediaContentRatingAustraliaTvRatingpreschoolers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingAustraliaTvRating(input)
	return &out, nil
}
