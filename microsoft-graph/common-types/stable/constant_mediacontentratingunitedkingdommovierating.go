package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedKingdomMovieRating string

const (
	MediaContentRatingUnitedKingdomMovieRating_Adults            MediaContentRatingUnitedKingdomMovieRating = "adults"
	MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Cinema MediaContentRatingUnitedKingdomMovieRating = "agesAbove12Cinema"
	MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Video  MediaContentRatingUnitedKingdomMovieRating = "agesAbove12Video"
	MediaContentRatingUnitedKingdomMovieRating_AgesAbove15       MediaContentRatingUnitedKingdomMovieRating = "agesAbove15"
	MediaContentRatingUnitedKingdomMovieRating_AllAllowed        MediaContentRatingUnitedKingdomMovieRating = "allAllowed"
	MediaContentRatingUnitedKingdomMovieRating_AllBlocked        MediaContentRatingUnitedKingdomMovieRating = "allBlocked"
	MediaContentRatingUnitedKingdomMovieRating_General           MediaContentRatingUnitedKingdomMovieRating = "general"
	MediaContentRatingUnitedKingdomMovieRating_ParentalGuidance  MediaContentRatingUnitedKingdomMovieRating = "parentalGuidance"
	MediaContentRatingUnitedKingdomMovieRating_UniversalChildren MediaContentRatingUnitedKingdomMovieRating = "universalChildren"
)

func PossibleValuesForMediaContentRatingUnitedKingdomMovieRating() []string {
	return []string{
		string(MediaContentRatingUnitedKingdomMovieRating_Adults),
		string(MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Cinema),
		string(MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Video),
		string(MediaContentRatingUnitedKingdomMovieRating_AgesAbove15),
		string(MediaContentRatingUnitedKingdomMovieRating_AllAllowed),
		string(MediaContentRatingUnitedKingdomMovieRating_AllBlocked),
		string(MediaContentRatingUnitedKingdomMovieRating_General),
		string(MediaContentRatingUnitedKingdomMovieRating_ParentalGuidance),
		string(MediaContentRatingUnitedKingdomMovieRating_UniversalChildren),
	}
}

func (s *MediaContentRatingUnitedKingdomMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingUnitedKingdomMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingUnitedKingdomMovieRating(input string) (*MediaContentRatingUnitedKingdomMovieRating, error) {
	vals := map[string]MediaContentRatingUnitedKingdomMovieRating{
		"adults":            MediaContentRatingUnitedKingdomMovieRating_Adults,
		"agesabove12cinema": MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Cinema,
		"agesabove12video":  MediaContentRatingUnitedKingdomMovieRating_AgesAbove12Video,
		"agesabove15":       MediaContentRatingUnitedKingdomMovieRating_AgesAbove15,
		"allallowed":        MediaContentRatingUnitedKingdomMovieRating_AllAllowed,
		"allblocked":        MediaContentRatingUnitedKingdomMovieRating_AllBlocked,
		"general":           MediaContentRatingUnitedKingdomMovieRating_General,
		"parentalguidance":  MediaContentRatingUnitedKingdomMovieRating_ParentalGuidance,
		"universalchildren": MediaContentRatingUnitedKingdomMovieRating_UniversalChildren,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedKingdomMovieRating(input)
	return &out, nil
}
