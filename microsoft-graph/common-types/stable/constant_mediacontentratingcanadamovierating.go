package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingCanadaMovieRating string

const (
	MediaContentRatingCanadaMovieRating_AgesAbove14      MediaContentRatingCanadaMovieRating = "agesAbove14"
	MediaContentRatingCanadaMovieRating_AgesAbove18      MediaContentRatingCanadaMovieRating = "agesAbove18"
	MediaContentRatingCanadaMovieRating_AllAllowed       MediaContentRatingCanadaMovieRating = "allAllowed"
	MediaContentRatingCanadaMovieRating_AllBlocked       MediaContentRatingCanadaMovieRating = "allBlocked"
	MediaContentRatingCanadaMovieRating_General          MediaContentRatingCanadaMovieRating = "general"
	MediaContentRatingCanadaMovieRating_ParentalGuidance MediaContentRatingCanadaMovieRating = "parentalGuidance"
	MediaContentRatingCanadaMovieRating_Restricted       MediaContentRatingCanadaMovieRating = "restricted"
)

func PossibleValuesForMediaContentRatingCanadaMovieRating() []string {
	return []string{
		string(MediaContentRatingCanadaMovieRating_AgesAbove14),
		string(MediaContentRatingCanadaMovieRating_AgesAbove18),
		string(MediaContentRatingCanadaMovieRating_AllAllowed),
		string(MediaContentRatingCanadaMovieRating_AllBlocked),
		string(MediaContentRatingCanadaMovieRating_General),
		string(MediaContentRatingCanadaMovieRating_ParentalGuidance),
		string(MediaContentRatingCanadaMovieRating_Restricted),
	}
}

func (s *MediaContentRatingCanadaMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingCanadaMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingCanadaMovieRating(input string) (*MediaContentRatingCanadaMovieRating, error) {
	vals := map[string]MediaContentRatingCanadaMovieRating{
		"agesabove14":      MediaContentRatingCanadaMovieRating_AgesAbove14,
		"agesabove18":      MediaContentRatingCanadaMovieRating_AgesAbove18,
		"allallowed":       MediaContentRatingCanadaMovieRating_AllAllowed,
		"allblocked":       MediaContentRatingCanadaMovieRating_AllBlocked,
		"general":          MediaContentRatingCanadaMovieRating_General,
		"parentalguidance": MediaContentRatingCanadaMovieRating_ParentalGuidance,
		"restricted":       MediaContentRatingCanadaMovieRating_Restricted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingCanadaMovieRating(input)
	return &out, nil
}
