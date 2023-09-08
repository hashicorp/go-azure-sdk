package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapanTvRating string

const (
	MediaContentRatingJapanTvRatingallAllowed      MediaContentRatingJapanTvRating = "AllAllowed"
	MediaContentRatingJapanTvRatingallBlocked      MediaContentRatingJapanTvRating = "AllBlocked"
	MediaContentRatingJapanTvRatingexplicitAllowed MediaContentRatingJapanTvRating = "ExplicitAllowed"
)

func PossibleValuesForMediaContentRatingJapanTvRating() []string {
	return []string{
		string(MediaContentRatingJapanTvRatingallAllowed),
		string(MediaContentRatingJapanTvRatingallBlocked),
		string(MediaContentRatingJapanTvRatingexplicitAllowed),
	}
}

func parseMediaContentRatingJapanTvRating(input string) (*MediaContentRatingJapanTvRating, error) {
	vals := map[string]MediaContentRatingJapanTvRating{
		"allallowed":      MediaContentRatingJapanTvRatingallAllowed,
		"allblocked":      MediaContentRatingJapanTvRatingallBlocked,
		"explicitallowed": MediaContentRatingJapanTvRatingexplicitAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingJapanTvRating(input)
	return &out, nil
}
