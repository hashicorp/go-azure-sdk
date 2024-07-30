package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingIrelandTvRating string

const (
	MediaContentRatingIrelandTvRating_AllAllowed          MediaContentRatingIrelandTvRating = "allAllowed"
	MediaContentRatingIrelandTvRating_AllBlocked          MediaContentRatingIrelandTvRating = "allBlocked"
	MediaContentRatingIrelandTvRating_Children            MediaContentRatingIrelandTvRating = "children"
	MediaContentRatingIrelandTvRating_General             MediaContentRatingIrelandTvRating = "general"
	MediaContentRatingIrelandTvRating_Mature              MediaContentRatingIrelandTvRating = "mature"
	MediaContentRatingIrelandTvRating_ParentalSupervision MediaContentRatingIrelandTvRating = "parentalSupervision"
	MediaContentRatingIrelandTvRating_YoungAdults         MediaContentRatingIrelandTvRating = "youngAdults"
)

func PossibleValuesForMediaContentRatingIrelandTvRating() []string {
	return []string{
		string(MediaContentRatingIrelandTvRating_AllAllowed),
		string(MediaContentRatingIrelandTvRating_AllBlocked),
		string(MediaContentRatingIrelandTvRating_Children),
		string(MediaContentRatingIrelandTvRating_General),
		string(MediaContentRatingIrelandTvRating_Mature),
		string(MediaContentRatingIrelandTvRating_ParentalSupervision),
		string(MediaContentRatingIrelandTvRating_YoungAdults),
	}
}

func (s *MediaContentRatingIrelandTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingIrelandTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingIrelandTvRating(input string) (*MediaContentRatingIrelandTvRating, error) {
	vals := map[string]MediaContentRatingIrelandTvRating{
		"allallowed":          MediaContentRatingIrelandTvRating_AllAllowed,
		"allblocked":          MediaContentRatingIrelandTvRating_AllBlocked,
		"children":            MediaContentRatingIrelandTvRating_Children,
		"general":             MediaContentRatingIrelandTvRating_General,
		"mature":              MediaContentRatingIrelandTvRating_Mature,
		"parentalsupervision": MediaContentRatingIrelandTvRating_ParentalSupervision,
		"youngadults":         MediaContentRatingIrelandTvRating_YoungAdults,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingIrelandTvRating(input)
	return &out, nil
}
