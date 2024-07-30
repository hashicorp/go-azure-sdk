package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedStatesMovieRating string

const (
	MediaContentRatingUnitedStatesMovieRating_Adults             MediaContentRatingUnitedStatesMovieRating = "adults"
	MediaContentRatingUnitedStatesMovieRating_AllAllowed         MediaContentRatingUnitedStatesMovieRating = "allAllowed"
	MediaContentRatingUnitedStatesMovieRating_AllBlocked         MediaContentRatingUnitedStatesMovieRating = "allBlocked"
	MediaContentRatingUnitedStatesMovieRating_General            MediaContentRatingUnitedStatesMovieRating = "general"
	MediaContentRatingUnitedStatesMovieRating_ParentalGuidance   MediaContentRatingUnitedStatesMovieRating = "parentalGuidance"
	MediaContentRatingUnitedStatesMovieRating_ParentalGuidance13 MediaContentRatingUnitedStatesMovieRating = "parentalGuidance13"
	MediaContentRatingUnitedStatesMovieRating_Restricted         MediaContentRatingUnitedStatesMovieRating = "restricted"
)

func PossibleValuesForMediaContentRatingUnitedStatesMovieRating() []string {
	return []string{
		string(MediaContentRatingUnitedStatesMovieRating_Adults),
		string(MediaContentRatingUnitedStatesMovieRating_AllAllowed),
		string(MediaContentRatingUnitedStatesMovieRating_AllBlocked),
		string(MediaContentRatingUnitedStatesMovieRating_General),
		string(MediaContentRatingUnitedStatesMovieRating_ParentalGuidance),
		string(MediaContentRatingUnitedStatesMovieRating_ParentalGuidance13),
		string(MediaContentRatingUnitedStatesMovieRating_Restricted),
	}
}

func (s *MediaContentRatingUnitedStatesMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingUnitedStatesMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingUnitedStatesMovieRating(input string) (*MediaContentRatingUnitedStatesMovieRating, error) {
	vals := map[string]MediaContentRatingUnitedStatesMovieRating{
		"adults":             MediaContentRatingUnitedStatesMovieRating_Adults,
		"allallowed":         MediaContentRatingUnitedStatesMovieRating_AllAllowed,
		"allblocked":         MediaContentRatingUnitedStatesMovieRating_AllBlocked,
		"general":            MediaContentRatingUnitedStatesMovieRating_General,
		"parentalguidance":   MediaContentRatingUnitedStatesMovieRating_ParentalGuidance,
		"parentalguidance13": MediaContentRatingUnitedStatesMovieRating_ParentalGuidance13,
		"restricted":         MediaContentRatingUnitedStatesMovieRating_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedStatesMovieRating(input)
	return &out, nil
}
