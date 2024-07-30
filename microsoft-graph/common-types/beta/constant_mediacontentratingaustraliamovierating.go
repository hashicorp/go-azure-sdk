package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingAustraliaMovieRating string

const (
	MediaContentRatingAustraliaMovieRating_AgesAbove15      MediaContentRatingAustraliaMovieRating = "agesAbove15"
	MediaContentRatingAustraliaMovieRating_AgesAbove18      MediaContentRatingAustraliaMovieRating = "agesAbove18"
	MediaContentRatingAustraliaMovieRating_AllAllowed       MediaContentRatingAustraliaMovieRating = "allAllowed"
	MediaContentRatingAustraliaMovieRating_AllBlocked       MediaContentRatingAustraliaMovieRating = "allBlocked"
	MediaContentRatingAustraliaMovieRating_General          MediaContentRatingAustraliaMovieRating = "general"
	MediaContentRatingAustraliaMovieRating_Mature           MediaContentRatingAustraliaMovieRating = "mature"
	MediaContentRatingAustraliaMovieRating_ParentalGuidance MediaContentRatingAustraliaMovieRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingAustraliaMovieRating() []string {
	return []string{
		string(MediaContentRatingAustraliaMovieRating_AgesAbove15),
		string(MediaContentRatingAustraliaMovieRating_AgesAbove18),
		string(MediaContentRatingAustraliaMovieRating_AllAllowed),
		string(MediaContentRatingAustraliaMovieRating_AllBlocked),
		string(MediaContentRatingAustraliaMovieRating_General),
		string(MediaContentRatingAustraliaMovieRating_Mature),
		string(MediaContentRatingAustraliaMovieRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingAustraliaMovieRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingAustraliaMovieRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingAustraliaMovieRating(input string) (*MediaContentRatingAustraliaMovieRating, error) {
	vals := map[string]MediaContentRatingAustraliaMovieRating{
		"agesabove15":      MediaContentRatingAustraliaMovieRating_AgesAbove15,
		"agesabove18":      MediaContentRatingAustraliaMovieRating_AgesAbove18,
		"allallowed":       MediaContentRatingAustraliaMovieRating_AllAllowed,
		"allblocked":       MediaContentRatingAustraliaMovieRating_AllBlocked,
		"general":          MediaContentRatingAustraliaMovieRating_General,
		"mature":           MediaContentRatingAustraliaMovieRating_Mature,
		"parentalguidance": MediaContentRatingAustraliaMovieRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingAustraliaMovieRating(input)
	return &out, nil
}
