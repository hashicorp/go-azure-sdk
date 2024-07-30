package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingFranceTvRating string

const (
	MediaContentRatingFranceTvRating_AgesAbove10 MediaContentRatingFranceTvRating = "agesAbove10"
	MediaContentRatingFranceTvRating_AgesAbove12 MediaContentRatingFranceTvRating = "agesAbove12"
	MediaContentRatingFranceTvRating_AgesAbove16 MediaContentRatingFranceTvRating = "agesAbove16"
	MediaContentRatingFranceTvRating_AgesAbove18 MediaContentRatingFranceTvRating = "agesAbove18"
	MediaContentRatingFranceTvRating_AllAllowed  MediaContentRatingFranceTvRating = "allAllowed"
	MediaContentRatingFranceTvRating_AllBlocked  MediaContentRatingFranceTvRating = "allBlocked"
)

func PossibleValuesForMediaContentRatingFranceTvRating() []string {
	return []string{
		string(MediaContentRatingFranceTvRating_AgesAbove10),
		string(MediaContentRatingFranceTvRating_AgesAbove12),
		string(MediaContentRatingFranceTvRating_AgesAbove16),
		string(MediaContentRatingFranceTvRating_AgesAbove18),
		string(MediaContentRatingFranceTvRating_AllAllowed),
		string(MediaContentRatingFranceTvRating_AllBlocked),
	}
}

func (s *MediaContentRatingFranceTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingFranceTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingFranceTvRating(input string) (*MediaContentRatingFranceTvRating, error) {
	vals := map[string]MediaContentRatingFranceTvRating{
		"agesabove10": MediaContentRatingFranceTvRating_AgesAbove10,
		"agesabove12": MediaContentRatingFranceTvRating_AgesAbove12,
		"agesabove16": MediaContentRatingFranceTvRating_AgesAbove16,
		"agesabove18": MediaContentRatingFranceTvRating_AgesAbove18,
		"allallowed":  MediaContentRatingFranceTvRating_AllAllowed,
		"allblocked":  MediaContentRatingFranceTvRating_AllBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingFranceTvRating(input)
	return &out, nil
}
