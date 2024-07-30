package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingNewZealandTvRating string

const (
	MediaContentRatingNewZealandTvRating_Adults           MediaContentRatingNewZealandTvRating = "adults"
	MediaContentRatingNewZealandTvRating_AllAllowed       MediaContentRatingNewZealandTvRating = "allAllowed"
	MediaContentRatingNewZealandTvRating_AllBlocked       MediaContentRatingNewZealandTvRating = "allBlocked"
	MediaContentRatingNewZealandTvRating_General          MediaContentRatingNewZealandTvRating = "general"
	MediaContentRatingNewZealandTvRating_ParentalGuidance MediaContentRatingNewZealandTvRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingNewZealandTvRating() []string {
	return []string{
		string(MediaContentRatingNewZealandTvRating_Adults),
		string(MediaContentRatingNewZealandTvRating_AllAllowed),
		string(MediaContentRatingNewZealandTvRating_AllBlocked),
		string(MediaContentRatingNewZealandTvRating_General),
		string(MediaContentRatingNewZealandTvRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingNewZealandTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingNewZealandTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingNewZealandTvRating(input string) (*MediaContentRatingNewZealandTvRating, error) {
	vals := map[string]MediaContentRatingNewZealandTvRating{
		"adults":           MediaContentRatingNewZealandTvRating_Adults,
		"allallowed":       MediaContentRatingNewZealandTvRating_AllAllowed,
		"allblocked":       MediaContentRatingNewZealandTvRating_AllBlocked,
		"general":          MediaContentRatingNewZealandTvRating_General,
		"parentalguidance": MediaContentRatingNewZealandTvRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingNewZealandTvRating(input)
	return &out, nil
}
