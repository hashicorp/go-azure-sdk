package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedKingdomMovieRating string

const (
	MediaContentRatingUnitedKingdomMovieRatingadults            MediaContentRatingUnitedKingdomMovieRating = "Adults"
	MediaContentRatingUnitedKingdomMovieRatingagesAbove12Cinema MediaContentRatingUnitedKingdomMovieRating = "AgesAbove12Cinema"
	MediaContentRatingUnitedKingdomMovieRatingagesAbove12Video  MediaContentRatingUnitedKingdomMovieRating = "AgesAbove12Video"
	MediaContentRatingUnitedKingdomMovieRatingagesAbove15       MediaContentRatingUnitedKingdomMovieRating = "AgesAbove15"
	MediaContentRatingUnitedKingdomMovieRatingallAllowed        MediaContentRatingUnitedKingdomMovieRating = "AllAllowed"
	MediaContentRatingUnitedKingdomMovieRatingallBlocked        MediaContentRatingUnitedKingdomMovieRating = "AllBlocked"
	MediaContentRatingUnitedKingdomMovieRatinggeneral           MediaContentRatingUnitedKingdomMovieRating = "General"
	MediaContentRatingUnitedKingdomMovieRatingparentalGuidance  MediaContentRatingUnitedKingdomMovieRating = "ParentalGuidance"
	MediaContentRatingUnitedKingdomMovieRatinguniversalChildren MediaContentRatingUnitedKingdomMovieRating = "UniversalChildren"
)

func PossibleValuesForMediaContentRatingUnitedKingdomMovieRating() []string {
	return []string{
		string(MediaContentRatingUnitedKingdomMovieRatingadults),
		string(MediaContentRatingUnitedKingdomMovieRatingagesAbove12Cinema),
		string(MediaContentRatingUnitedKingdomMovieRatingagesAbove12Video),
		string(MediaContentRatingUnitedKingdomMovieRatingagesAbove15),
		string(MediaContentRatingUnitedKingdomMovieRatingallAllowed),
		string(MediaContentRatingUnitedKingdomMovieRatingallBlocked),
		string(MediaContentRatingUnitedKingdomMovieRatinggeneral),
		string(MediaContentRatingUnitedKingdomMovieRatingparentalGuidance),
		string(MediaContentRatingUnitedKingdomMovieRatinguniversalChildren),
	}
}

func parseMediaContentRatingUnitedKingdomMovieRating(input string) (*MediaContentRatingUnitedKingdomMovieRating, error) {
	vals := map[string]MediaContentRatingUnitedKingdomMovieRating{
		"adults":            MediaContentRatingUnitedKingdomMovieRatingadults,
		"agesabove12cinema": MediaContentRatingUnitedKingdomMovieRatingagesAbove12Cinema,
		"agesabove12video":  MediaContentRatingUnitedKingdomMovieRatingagesAbove12Video,
		"agesabove15":       MediaContentRatingUnitedKingdomMovieRatingagesAbove15,
		"allallowed":        MediaContentRatingUnitedKingdomMovieRatingallAllowed,
		"allblocked":        MediaContentRatingUnitedKingdomMovieRatingallBlocked,
		"general":           MediaContentRatingUnitedKingdomMovieRatinggeneral,
		"parentalguidance":  MediaContentRatingUnitedKingdomMovieRatingparentalGuidance,
		"universalchildren": MediaContentRatingUnitedKingdomMovieRatinguniversalChildren,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedKingdomMovieRating(input)
	return &out, nil
}
