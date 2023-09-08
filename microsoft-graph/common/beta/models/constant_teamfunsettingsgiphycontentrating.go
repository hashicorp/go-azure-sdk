package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamFunSettingsGiphyContentRating string

const (
	TeamFunSettingsGiphyContentRatingmoderate TeamFunSettingsGiphyContentRating = "Moderate"
	TeamFunSettingsGiphyContentRatingstrict   TeamFunSettingsGiphyContentRating = "Strict"
)

func PossibleValuesForTeamFunSettingsGiphyContentRating() []string {
	return []string{
		string(TeamFunSettingsGiphyContentRatingmoderate),
		string(TeamFunSettingsGiphyContentRatingstrict),
	}
}

func parseTeamFunSettingsGiphyContentRating(input string) (*TeamFunSettingsGiphyContentRating, error) {
	vals := map[string]TeamFunSettingsGiphyContentRating{
		"moderate": TeamFunSettingsGiphyContentRatingmoderate,
		"strict":   TeamFunSettingsGiphyContentRatingstrict,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamFunSettingsGiphyContentRating(input)
	return &out, nil
}
