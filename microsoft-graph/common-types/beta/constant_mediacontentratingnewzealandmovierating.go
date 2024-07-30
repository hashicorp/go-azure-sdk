package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealandMovieRating string

const (
	MediaContentRatingNewZealandMovieRating_AgesAbove13           MediaContentRatingNewZealandMovieRating = "agesAbove13"
	MediaContentRatingNewZealandMovieRating_AgesAbove15           MediaContentRatingNewZealandMovieRating = "agesAbove15"
	MediaContentRatingNewZealandMovieRating_AgesAbove16           MediaContentRatingNewZealandMovieRating = "agesAbove16"
	MediaContentRatingNewZealandMovieRating_AgesAbove16Restricted MediaContentRatingNewZealandMovieRating = "agesAbove16Restricted"
	MediaContentRatingNewZealandMovieRating_AgesAbove18           MediaContentRatingNewZealandMovieRating = "agesAbove18"
	MediaContentRatingNewZealandMovieRating_AllAllowed            MediaContentRatingNewZealandMovieRating = "allAllowed"
	MediaContentRatingNewZealandMovieRating_AllBlocked            MediaContentRatingNewZealandMovieRating = "allBlocked"
	MediaContentRatingNewZealandMovieRating_General               MediaContentRatingNewZealandMovieRating = "general"
	MediaContentRatingNewZealandMovieRating_Mature                MediaContentRatingNewZealandMovieRating = "mature"
	MediaContentRatingNewZealandMovieRating_ParentalGuidance      MediaContentRatingNewZealandMovieRating = "parentalGuidance"
	MediaContentRatingNewZealandMovieRating_Restricted            MediaContentRatingNewZealandMovieRating = "restricted"
)

func PossibleValuesForMediaContentRatingNewZealandMovieRating() []string {
	return []string{
		string(MediaContentRatingNewZealandMovieRating_AgesAbove13),
		string(MediaContentRatingNewZealandMovieRating_AgesAbove15),
		string(MediaContentRatingNewZealandMovieRating_AgesAbove16),
		string(MediaContentRatingNewZealandMovieRating_AgesAbove16Restricted),
		string(MediaContentRatingNewZealandMovieRating_AgesAbove18),
		string(MediaContentRatingNewZealandMovieRating_AllAllowed),
		string(MediaContentRatingNewZealandMovieRating_AllBlocked),
		string(MediaContentRatingNewZealandMovieRating_General),
		string(MediaContentRatingNewZealandMovieRating_Mature),
		string(MediaContentRatingNewZealandMovieRating_ParentalGuidance),
		string(MediaContentRatingNewZealandMovieRating_Restricted),
	}
}

func (s *MediaContentRatingNewZealandMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingNewZealandMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingNewZealandMovieRating(input string) (*MediaContentRatingNewZealandMovieRating, error) {
	vals := map[string]MediaContentRatingNewZealandMovieRating{
		"agesabove13":           MediaContentRatingNewZealandMovieRating_AgesAbove13,
		"agesabove15":           MediaContentRatingNewZealandMovieRating_AgesAbove15,
		"agesabove16":           MediaContentRatingNewZealandMovieRating_AgesAbove16,
		"agesabove16restricted": MediaContentRatingNewZealandMovieRating_AgesAbove16Restricted,
		"agesabove18":           MediaContentRatingNewZealandMovieRating_AgesAbove18,
		"allallowed":            MediaContentRatingNewZealandMovieRating_AllAllowed,
		"allblocked":            MediaContentRatingNewZealandMovieRating_AllBlocked,
		"general":               MediaContentRatingNewZealandMovieRating_General,
		"mature":                MediaContentRatingNewZealandMovieRating_Mature,
		"parentalguidance":      MediaContentRatingNewZealandMovieRating_ParentalGuidance,
		"restricted":            MediaContentRatingNewZealandMovieRating_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingNewZealandMovieRating(input)
	return &out, nil
}
