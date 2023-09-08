package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIrelandTvRating string

const (
	MediaContentRatingIrelandTvRatingallAllowed          MediaContentRatingIrelandTvRating = "AllAllowed"
	MediaContentRatingIrelandTvRatingallBlocked          MediaContentRatingIrelandTvRating = "AllBlocked"
	MediaContentRatingIrelandTvRatingchildren            MediaContentRatingIrelandTvRating = "Children"
	MediaContentRatingIrelandTvRatinggeneral             MediaContentRatingIrelandTvRating = "General"
	MediaContentRatingIrelandTvRatingmature              MediaContentRatingIrelandTvRating = "Mature"
	MediaContentRatingIrelandTvRatingparentalSupervision MediaContentRatingIrelandTvRating = "ParentalSupervision"
	MediaContentRatingIrelandTvRatingyoungAdults         MediaContentRatingIrelandTvRating = "YoungAdults"
)

func PossibleValuesForMediaContentRatingIrelandTvRating() []string {
	return []string{
		string(MediaContentRatingIrelandTvRatingallAllowed),
		string(MediaContentRatingIrelandTvRatingallBlocked),
		string(MediaContentRatingIrelandTvRatingchildren),
		string(MediaContentRatingIrelandTvRatinggeneral),
		string(MediaContentRatingIrelandTvRatingmature),
		string(MediaContentRatingIrelandTvRatingparentalSupervision),
		string(MediaContentRatingIrelandTvRatingyoungAdults),
	}
}

func parseMediaContentRatingIrelandTvRating(input string) (*MediaContentRatingIrelandTvRating, error) {
	vals := map[string]MediaContentRatingIrelandTvRating{
		"allallowed":          MediaContentRatingIrelandTvRatingallAllowed,
		"allblocked":          MediaContentRatingIrelandTvRatingallBlocked,
		"children":            MediaContentRatingIrelandTvRatingchildren,
		"general":             MediaContentRatingIrelandTvRatinggeneral,
		"mature":              MediaContentRatingIrelandTvRatingmature,
		"parentalsupervision": MediaContentRatingIrelandTvRatingparentalSupervision,
		"youngadults":         MediaContentRatingIrelandTvRatingyoungAdults,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingIrelandTvRating(input)
	return &out, nil
}
