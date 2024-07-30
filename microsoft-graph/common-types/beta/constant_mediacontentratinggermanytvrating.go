package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingGermanyTvRating string

const (
	MediaContentRatingGermanyTvRating_Adults      MediaContentRatingGermanyTvRating = "adults"
	MediaContentRatingGermanyTvRating_AgesAbove12 MediaContentRatingGermanyTvRating = "agesAbove12"
	MediaContentRatingGermanyTvRating_AgesAbove16 MediaContentRatingGermanyTvRating = "agesAbove16"
	MediaContentRatingGermanyTvRating_AgesAbove6  MediaContentRatingGermanyTvRating = "agesAbove6"
	MediaContentRatingGermanyTvRating_AllAllowed  MediaContentRatingGermanyTvRating = "allAllowed"
	MediaContentRatingGermanyTvRating_AllBlocked  MediaContentRatingGermanyTvRating = "allBlocked"
	MediaContentRatingGermanyTvRating_General     MediaContentRatingGermanyTvRating = "general"
)

func PossibleValuesForMediaContentRatingGermanyTvRating() []string {
	return []string{
		string(MediaContentRatingGermanyTvRating_Adults),
		string(MediaContentRatingGermanyTvRating_AgesAbove12),
		string(MediaContentRatingGermanyTvRating_AgesAbove16),
		string(MediaContentRatingGermanyTvRating_AgesAbove6),
		string(MediaContentRatingGermanyTvRating_AllAllowed),
		string(MediaContentRatingGermanyTvRating_AllBlocked),
		string(MediaContentRatingGermanyTvRating_General),
	}
}

func (s *MediaContentRatingGermanyTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingGermanyTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingGermanyTvRating(input string) (*MediaContentRatingGermanyTvRating, error) {
	vals := map[string]MediaContentRatingGermanyTvRating{
		"adults":      MediaContentRatingGermanyTvRating_Adults,
		"agesabove12": MediaContentRatingGermanyTvRating_AgesAbove12,
		"agesabove16": MediaContentRatingGermanyTvRating_AgesAbove16,
		"agesabove6":  MediaContentRatingGermanyTvRating_AgesAbove6,
		"allallowed":  MediaContentRatingGermanyTvRating_AllAllowed,
		"allblocked":  MediaContentRatingGermanyTvRating_AllBlocked,
		"general":     MediaContentRatingGermanyTvRating_General,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingGermanyTvRating(input)
	return &out, nil
}
