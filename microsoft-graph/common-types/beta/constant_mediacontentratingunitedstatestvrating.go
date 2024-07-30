package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingUnitedStatesTvRating string

const (
	MediaContentRatingUnitedStatesTvRating_Adults           MediaContentRatingUnitedStatesTvRating = "adults"
	MediaContentRatingUnitedStatesTvRating_AllAllowed       MediaContentRatingUnitedStatesTvRating = "allAllowed"
	MediaContentRatingUnitedStatesTvRating_AllBlocked       MediaContentRatingUnitedStatesTvRating = "allBlocked"
	MediaContentRatingUnitedStatesTvRating_ChildrenAbove14  MediaContentRatingUnitedStatesTvRating = "childrenAbove14"
	MediaContentRatingUnitedStatesTvRating_ChildrenAbove7   MediaContentRatingUnitedStatesTvRating = "childrenAbove7"
	MediaContentRatingUnitedStatesTvRating_ChildrenAll      MediaContentRatingUnitedStatesTvRating = "childrenAll"
	MediaContentRatingUnitedStatesTvRating_General          MediaContentRatingUnitedStatesTvRating = "general"
	MediaContentRatingUnitedStatesTvRating_ParentalGuidance MediaContentRatingUnitedStatesTvRating = "parentalGuidance"
)

func PossibleValuesForMediaContentRatingUnitedStatesTvRating() []string {
	return []string{
		string(MediaContentRatingUnitedStatesTvRating_Adults),
		string(MediaContentRatingUnitedStatesTvRating_AllAllowed),
		string(MediaContentRatingUnitedStatesTvRating_AllBlocked),
		string(MediaContentRatingUnitedStatesTvRating_ChildrenAbove14),
		string(MediaContentRatingUnitedStatesTvRating_ChildrenAbove7),
		string(MediaContentRatingUnitedStatesTvRating_ChildrenAll),
		string(MediaContentRatingUnitedStatesTvRating_General),
		string(MediaContentRatingUnitedStatesTvRating_ParentalGuidance),
	}
}

func (s *MediaContentRatingUnitedStatesTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingUnitedStatesTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingUnitedStatesTvRating(input string) (*MediaContentRatingUnitedStatesTvRating, error) {
	vals := map[string]MediaContentRatingUnitedStatesTvRating{
		"adults":           MediaContentRatingUnitedStatesTvRating_Adults,
		"allallowed":       MediaContentRatingUnitedStatesTvRating_AllAllowed,
		"allblocked":       MediaContentRatingUnitedStatesTvRating_AllBlocked,
		"childrenabove14":  MediaContentRatingUnitedStatesTvRating_ChildrenAbove14,
		"childrenabove7":   MediaContentRatingUnitedStatesTvRating_ChildrenAbove7,
		"childrenall":      MediaContentRatingUnitedStatesTvRating_ChildrenAll,
		"general":          MediaContentRatingUnitedStatesTvRating_General,
		"parentalguidance": MediaContentRatingUnitedStatesTvRating_ParentalGuidance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingUnitedStatesTvRating(input)
	return &out, nil
}
