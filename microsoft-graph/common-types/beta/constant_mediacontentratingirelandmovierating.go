package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIrelandMovieRating string

const (
	MediaContentRatingIrelandMovieRating_Adults           MediaContentRatingIrelandMovieRating = "adults"
	MediaContentRatingIrelandMovieRating_AgesAbove12      MediaContentRatingIrelandMovieRating = "agesAbove12"
	MediaContentRatingIrelandMovieRating_AgesAbove15      MediaContentRatingIrelandMovieRating = "agesAbove15"
	MediaContentRatingIrelandMovieRating_AgesAbove16      MediaContentRatingIrelandMovieRating = "agesAbove16"
	MediaContentRatingIrelandMovieRating_AllAllowed       MediaContentRatingIrelandMovieRating = "allAllowed"
	MediaContentRatingIrelandMovieRating_AllBlocked       MediaContentRatingIrelandMovieRating = "allBlocked"
	MediaContentRatingIrelandMovieRating_General          MediaContentRatingIrelandMovieRating = "general"
	MediaContentRatingIrelandMovieRating_ParentalGuidance MediaContentRatingIrelandMovieRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingIrelandMovieRating() []string {
	return []string{
		string(MediaContentRatingIrelandMovieRating_Adults),
		string(MediaContentRatingIrelandMovieRating_AgesAbove12),
		string(MediaContentRatingIrelandMovieRating_AgesAbove15),
		string(MediaContentRatingIrelandMovieRating_AgesAbove16),
		string(MediaContentRatingIrelandMovieRating_AllAllowed),
		string(MediaContentRatingIrelandMovieRating_AllBlocked),
		string(MediaContentRatingIrelandMovieRating_General),
		string(MediaContentRatingIrelandMovieRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingIrelandMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingIrelandMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingIrelandMovieRating(input string) (*MediaContentRatingIrelandMovieRating, error) {
	vals := map[string]MediaContentRatingIrelandMovieRating{
		"adults":           MediaContentRatingIrelandMovieRating_Adults,
		"agesabove12":      MediaContentRatingIrelandMovieRating_AgesAbove12,
		"agesabove15":      MediaContentRatingIrelandMovieRating_AgesAbove15,
		"agesabove16":      MediaContentRatingIrelandMovieRating_AgesAbove16,
		"allallowed":       MediaContentRatingIrelandMovieRating_AllAllowed,
		"allblocked":       MediaContentRatingIrelandMovieRating_AllBlocked,
		"general":          MediaContentRatingIrelandMovieRating_General,
		"parentalguidance": MediaContentRatingIrelandMovieRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingIrelandMovieRating(input)
	return &out, nil
}
