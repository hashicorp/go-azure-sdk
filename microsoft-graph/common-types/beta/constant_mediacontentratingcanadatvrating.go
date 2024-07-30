package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingCanadaTvRating string

const (
	MediaContentRatingCanadaTvRating_AgesAbove14      MediaContentRatingCanadaTvRating = "agesAbove14"
	MediaContentRatingCanadaTvRating_AgesAbove18      MediaContentRatingCanadaTvRating = "agesAbove18"
	MediaContentRatingCanadaTvRating_AllAllowed       MediaContentRatingCanadaTvRating = "allAllowed"
	MediaContentRatingCanadaTvRating_AllBlocked       MediaContentRatingCanadaTvRating = "allBlocked"
	MediaContentRatingCanadaTvRating_Children         MediaContentRatingCanadaTvRating = "children"
	MediaContentRatingCanadaTvRating_ChildrenAbove8   MediaContentRatingCanadaTvRating = "childrenAbove8"
	MediaContentRatingCanadaTvRating_General          MediaContentRatingCanadaTvRating = "general"
	MediaContentRatingCanadaTvRating_ParentalGuidance MediaContentRatingCanadaTvRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingCanadaTvRating() []string {
	return []string{
		string(MediaContentRatingCanadaTvRating_AgesAbove14),
		string(MediaContentRatingCanadaTvRating_AgesAbove18),
		string(MediaContentRatingCanadaTvRating_AllAllowed),
		string(MediaContentRatingCanadaTvRating_AllBlocked),
		string(MediaContentRatingCanadaTvRating_Children),
		string(MediaContentRatingCanadaTvRating_ChildrenAbove8),
		string(MediaContentRatingCanadaTvRating_General),
		string(MediaContentRatingCanadaTvRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingCanadaTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingCanadaTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingCanadaTvRating(input string) (*MediaContentRatingCanadaTvRating, error) {
	vals := map[string]MediaContentRatingCanadaTvRating{
		"agesabove14":      MediaContentRatingCanadaTvRating_AgesAbove14,
		"agesabove18":      MediaContentRatingCanadaTvRating_AgesAbove18,
		"allallowed":       MediaContentRatingCanadaTvRating_AllAllowed,
		"allblocked":       MediaContentRatingCanadaTvRating_AllBlocked,
		"children":         MediaContentRatingCanadaTvRating_Children,
		"childrenabove8":   MediaContentRatingCanadaTvRating_ChildrenAbove8,
		"general":          MediaContentRatingCanadaTvRating_General,
		"parentalguidance": MediaContentRatingCanadaTvRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingCanadaTvRating(input)
	return &out, nil
}
