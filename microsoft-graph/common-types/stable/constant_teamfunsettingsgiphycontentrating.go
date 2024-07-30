package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamFunSettingsGiphyContentRating string

const (
	TeamFunSettingsGiphyContentRating_Moderate TeamFunSettingsGiphyContentRating = "moderate"
	TeamFunSettingsGiphyContentRating_Strict   TeamFunSettingsGiphyContentRating = "strict"
)

func PossibleValuesForTeamFunSettingsGiphyContentRating() []string {
	return []string{
		string(TeamFunSettingsGiphyContentRating_Moderate),
		string(TeamFunSettingsGiphyContentRating_Strict),
	}
}

func (s *TeamFunSettingsGiphyContentRating) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamFunSettingsGiphyContentRating(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamFunSettingsGiphyContentRating(input string) (*TeamFunSettingsGiphyContentRating, error) {
	vals := map[string]TeamFunSettingsGiphyContentRating{
		"moderate": TeamFunSettingsGiphyContentRating_Moderate,
		"strict":   TeamFunSettingsGiphyContentRating_Strict,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamFunSettingsGiphyContentRating(input)
	return &out, nil
}
