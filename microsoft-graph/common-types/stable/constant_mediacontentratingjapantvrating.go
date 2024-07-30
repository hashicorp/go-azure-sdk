package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MediaContentRatingJapanTvRating string

const (
	MediaContentRatingJapanTvRating_AllAllowed      MediaContentRatingJapanTvRating = "allAllowed"
	MediaContentRatingJapanTvRating_AllBlocked      MediaContentRatingJapanTvRating = "allBlocked"
	MediaContentRatingJapanTvRating_ExplicitAllowed MediaContentRatingJapanTvRating = "explicitAllowed"
)

func PossibleValuesForMediaContentRatingJapanTvRating() []string {
	return []string{
		string(MediaContentRatingJapanTvRating_AllAllowed),
		string(MediaContentRatingJapanTvRating_AllBlocked),
		string(MediaContentRatingJapanTvRating_ExplicitAllowed),
	}
}

func (s *MediaContentRatingJapanTvRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMediaContentRatingJapanTvRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMediaContentRatingJapanTvRating(input string) (*MediaContentRatingJapanTvRating, error) {
	vals := map[string]MediaContentRatingJapanTvRating{
		"allallowed":      MediaContentRatingJapanTvRating_AllAllowed,
		"allblocked":      MediaContentRatingJapanTvRating_AllBlocked,
		"explicitallowed": MediaContentRatingJapanTvRating_ExplicitAllowed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MediaContentRatingJapanTvRating(input)
	return &out, nil
}
