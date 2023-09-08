package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedStatesTvRating string

const (
	MediaContentRatingUnitedStatesTvRatingadults           MediaContentRatingUnitedStatesTvRating = "Adults"
	MediaContentRatingUnitedStatesTvRatingallAllowed       MediaContentRatingUnitedStatesTvRating = "AllAllowed"
	MediaContentRatingUnitedStatesTvRatingallBlocked       MediaContentRatingUnitedStatesTvRating = "AllBlocked"
	MediaContentRatingUnitedStatesTvRatingchildrenAbove14  MediaContentRatingUnitedStatesTvRating = "ChildrenAbove14"
	MediaContentRatingUnitedStatesTvRatingchildrenAbove7   MediaContentRatingUnitedStatesTvRating = "ChildrenAbove7"
	MediaContentRatingUnitedStatesTvRatingchildrenAll      MediaContentRatingUnitedStatesTvRating = "ChildrenAll"
	MediaContentRatingUnitedStatesTvRatinggeneral          MediaContentRatingUnitedStatesTvRating = "General"
	MediaContentRatingUnitedStatesTvRatingparentalGuidance MediaContentRatingUnitedStatesTvRating = "ParentalGuidance"
)

func PossibleValuesForMediaContentRatingUnitedStatesTvRating() []string {
	return []string{
		string(MediaContentRatingUnitedStatesTvRatingadults),
		string(MediaContentRatingUnitedStatesTvRatingallAllowed),
		string(MediaContentRatingUnitedStatesTvRatingallBlocked),
		string(MediaContentRatingUnitedStatesTvRatingchildrenAbove14),
		string(MediaContentRatingUnitedStatesTvRatingchildrenAbove7),
		string(MediaContentRatingUnitedStatesTvRatingchildrenAll),
		string(MediaContentRatingUnitedStatesTvRatinggeneral),
		string(MediaContentRatingUnitedStatesTvRatingparentalGuidance),
	}
}

func parseMediaContentRatingUnitedStatesTvRating(input string) (*MediaContentRatingUnitedStatesTvRating, error) {
	vals := map[string]MediaContentRatingUnitedStatesTvRating{
		"adults":           MediaContentRatingUnitedStatesTvRatingadults,
		"allallowed":       MediaContentRatingUnitedStatesTvRatingallAllowed,
		"allblocked":       MediaContentRatingUnitedStatesTvRatingallBlocked,
		"childrenabove14":  MediaContentRatingUnitedStatesTvRatingchildrenAbove14,
		"childrenabove7":   MediaContentRatingUnitedStatesTvRatingchildrenAbove7,
		"childrenall":      MediaContentRatingUnitedStatesTvRatingchildrenAll,
		"general":          MediaContentRatingUnitedStatesTvRatinggeneral,
		"parentalguidance": MediaContentRatingUnitedStatesTvRatingparentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedStatesTvRating(input)
	return &out, nil
}
