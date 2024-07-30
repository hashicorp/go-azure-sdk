package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapanMovieRating string

const (
	MediaContentRatingJapanMovieRating_AgesAbove15      MediaContentRatingJapanMovieRating = "agesAbove15"
	MediaContentRatingJapanMovieRating_AgesAbove18      MediaContentRatingJapanMovieRating = "agesAbove18"
	MediaContentRatingJapanMovieRating_AllAllowed       MediaContentRatingJapanMovieRating = "allAllowed"
	MediaContentRatingJapanMovieRating_AllBlocked       MediaContentRatingJapanMovieRating = "allBlocked"
	MediaContentRatingJapanMovieRating_General          MediaContentRatingJapanMovieRating = "general"
	MediaContentRatingJapanMovieRating_ParentalGuidance MediaContentRatingJapanMovieRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingJapanMovieRating() []string {
	return []string{
		string(MediaContentRatingJapanMovieRating_AgesAbove15),
		string(MediaContentRatingJapanMovieRating_AgesAbove18),
		string(MediaContentRatingJapanMovieRating_AllAllowed),
		string(MediaContentRatingJapanMovieRating_AllBlocked),
		string(MediaContentRatingJapanMovieRating_General),
		string(MediaContentRatingJapanMovieRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingJapanMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingJapanMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingJapanMovieRating(input string) (*MediaContentRatingJapanMovieRating, error) {
	vals := map[string]MediaContentRatingJapanMovieRating{
		"agesabove15":      MediaContentRatingJapanMovieRating_AgesAbove15,
		"agesabove18":      MediaContentRatingJapanMovieRating_AgesAbove18,
		"allallowed":       MediaContentRatingJapanMovieRating_AllAllowed,
		"allblocked":       MediaContentRatingJapanMovieRating_AllBlocked,
		"general":          MediaContentRatingJapanMovieRating_General,
		"parentalguidance": MediaContentRatingJapanMovieRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingJapanMovieRating(input)
	return &out, nil
}
