package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingGermanyMovieRating string

const (
	MediaContentRatingGermanyMovieRating_Adults      MediaContentRatingGermanyMovieRating = "adults"
	MediaContentRatingGermanyMovieRating_AgesAbove12 MediaContentRatingGermanyMovieRating = "agesAbove12"
	MediaContentRatingGermanyMovieRating_AgesAbove16 MediaContentRatingGermanyMovieRating = "agesAbove16"
	MediaContentRatingGermanyMovieRating_AgesAbove6  MediaContentRatingGermanyMovieRating = "agesAbove6"
	MediaContentRatingGermanyMovieRating_AllAllowed  MediaContentRatingGermanyMovieRating = "allAllowed"
	MediaContentRatingGermanyMovieRating_AllBlocked  MediaContentRatingGermanyMovieRating = "allBlocked"
	MediaContentRatingGermanyMovieRating_General     MediaContentRatingGermanyMovieRating = "general"
)

func PossibleValuesForMediaContentRatingGermanyMovieRating() []string {
	return []string{
		string(MediaContentRatingGermanyMovieRating_Adults),
		string(MediaContentRatingGermanyMovieRating_AgesAbove12),
		string(MediaContentRatingGermanyMovieRating_AgesAbove16),
		string(MediaContentRatingGermanyMovieRating_AgesAbove6),
		string(MediaContentRatingGermanyMovieRating_AllAllowed),
		string(MediaContentRatingGermanyMovieRating_AllBlocked),
		string(MediaContentRatingGermanyMovieRating_General),
	}
}

func (s *MediaContentRatingGermanyMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingGermanyMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingGermanyMovieRating(input string) (*MediaContentRatingGermanyMovieRating, error) {
	vals := map[string]MediaContentRatingGermanyMovieRating{
		"adults":      MediaContentRatingGermanyMovieRating_Adults,
		"agesabove12": MediaContentRatingGermanyMovieRating_AgesAbove12,
		"agesabove16": MediaContentRatingGermanyMovieRating_AgesAbove16,
		"agesabove6":  MediaContentRatingGermanyMovieRating_AgesAbove6,
		"allallowed":  MediaContentRatingGermanyMovieRating_AllAllowed,
		"allblocked":  MediaContentRatingGermanyMovieRating_AllBlocked,
		"general":     MediaContentRatingGermanyMovieRating_General,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingGermanyMovieRating(input)
	return &out, nil
}
