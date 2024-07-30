package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingAustraliaTvRating string

const (
	MediaContentRatingAustraliaTvRating_AgesAbove15              MediaContentRatingAustraliaTvRating = "agesAbove15"
	MediaContentRatingAustraliaTvRating_AgesAbove15AdultViolence MediaContentRatingAustraliaTvRating = "agesAbove15AdultViolence"
	MediaContentRatingAustraliaTvRating_AllAllowed               MediaContentRatingAustraliaTvRating = "allAllowed"
	MediaContentRatingAustraliaTvRating_AllBlocked               MediaContentRatingAustraliaTvRating = "allBlocked"
	MediaContentRatingAustraliaTvRating_Children                 MediaContentRatingAustraliaTvRating = "children"
	MediaContentRatingAustraliaTvRating_General                  MediaContentRatingAustraliaTvRating = "general"
	MediaContentRatingAustraliaTvRating_Mature                   MediaContentRatingAustraliaTvRating = "mature"
	MediaContentRatingAustraliaTvRating_ParentalGuidance         MediaContentRatingAustraliaTvRating = "parentalGuidance"
	MediaContentRatingAustraliaTvRating_Preschoolers             MediaContentRatingAustraliaTvRating = "preschoolers"
)

func PossibleValuesForMediaContentRatingAustraliaTvRating() []string {
	return []string{
		string(MediaContentRatingAustraliaTvRating_AgesAbove15),
		string(MediaContentRatingAustraliaTvRating_AgesAbove15AdultViolence),
		string(MediaContentRatingAustraliaTvRating_AllAllowed),
		string(MediaContentRatingAustraliaTvRating_AllBlocked),
		string(MediaContentRatingAustraliaTvRating_Children),
		string(MediaContentRatingAustraliaTvRating_General),
		string(MediaContentRatingAustraliaTvRating_Mature),
		string(MediaContentRatingAustraliaTvRating_ParentalGuidance),
		string(MediaContentRatingAustraliaTvRating_Preschoolers),
	}
}

func (s *MediaContentRatingAustraliaTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingAustraliaTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingAustraliaTvRating(input string) (*MediaContentRatingAustraliaTvRating, error) {
	vals := map[string]MediaContentRatingAustraliaTvRating{
		"agesabove15":              MediaContentRatingAustraliaTvRating_AgesAbove15,
		"agesabove15adultviolence": MediaContentRatingAustraliaTvRating_AgesAbove15AdultViolence,
		"allallowed":               MediaContentRatingAustraliaTvRating_AllAllowed,
		"allblocked":               MediaContentRatingAustraliaTvRating_AllBlocked,
		"children":                 MediaContentRatingAustraliaTvRating_Children,
		"general":                  MediaContentRatingAustraliaTvRating_General,
		"mature":                   MediaContentRatingAustraliaTvRating_Mature,
		"parentalguidance":         MediaContentRatingAustraliaTvRating_ParentalGuidance,
		"preschoolers":             MediaContentRatingAustraliaTvRating_Preschoolers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingAustraliaTvRating(input)
	return &out, nil
}
