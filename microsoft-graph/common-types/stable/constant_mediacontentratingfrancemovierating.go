package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingFranceMovieRating string

const (
	MediaContentRatingFranceMovieRating_AgesAbove10 MediaContentRatingFranceMovieRating = "agesAbove10"
	MediaContentRatingFranceMovieRating_AgesAbove12 MediaContentRatingFranceMovieRating = "agesAbove12"
	MediaContentRatingFranceMovieRating_AgesAbove16 MediaContentRatingFranceMovieRating = "agesAbove16"
	MediaContentRatingFranceMovieRating_AgesAbove18 MediaContentRatingFranceMovieRating = "agesAbove18"
	MediaContentRatingFranceMovieRating_AllAllowed  MediaContentRatingFranceMovieRating = "allAllowed"
	MediaContentRatingFranceMovieRating_AllBlocked  MediaContentRatingFranceMovieRating = "allBlocked"
)

func PossibleValuesForMediaContentRatingFranceMovieRating() []string {
	return []string{
		string(MediaContentRatingFranceMovieRating_AgesAbove10),
		string(MediaContentRatingFranceMovieRating_AgesAbove12),
		string(MediaContentRatingFranceMovieRating_AgesAbove16),
		string(MediaContentRatingFranceMovieRating_AgesAbove18),
		string(MediaContentRatingFranceMovieRating_AllAllowed),
		string(MediaContentRatingFranceMovieRating_AllBlocked),
	}
}

func (s *MediaContentRatingFranceMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingFranceMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingFranceMovieRating(input string) (*MediaContentRatingFranceMovieRating, error) {
	vals := map[string]MediaContentRatingFranceMovieRating{
		"agesabove10": MediaContentRatingFranceMovieRating_AgesAbove10,
		"agesabove12": MediaContentRatingFranceMovieRating_AgesAbove12,
		"agesabove16": MediaContentRatingFranceMovieRating_AgesAbove16,
		"agesabove18": MediaContentRatingFranceMovieRating_AgesAbove18,
		"allallowed":  MediaContentRatingFranceMovieRating_AllAllowed,
		"allblocked":  MediaContentRatingFranceMovieRating_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingFranceMovieRating(input)
	return &out, nil
}
