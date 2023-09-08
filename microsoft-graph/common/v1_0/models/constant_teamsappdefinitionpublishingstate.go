package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinitionPublishingState string

const (
	TeamsAppDefinitionPublishingStatepublished TeamsAppDefinitionPublishingState = "Published"
	TeamsAppDefinitionPublishingStaterejected  TeamsAppDefinitionPublishingState = "Rejected"
	TeamsAppDefinitionPublishingStatesubmitted TeamsAppDefinitionPublishingState = "Submitted"
)

func PossibleValuesForTeamsAppDefinitionPublishingState() []string {
	return []string{
		string(TeamsAppDefinitionPublishingStatepublished),
		string(TeamsAppDefinitionPublishingStaterejected),
		string(TeamsAppDefinitionPublishingStatesubmitted),
	}
}

func parseTeamsAppDefinitionPublishingState(input string) (*TeamsAppDefinitionPublishingState, error) {
	vals := map[string]TeamsAppDefinitionPublishingState{
		"published": TeamsAppDefinitionPublishingStatepublished,
		"rejected":  TeamsAppDefinitionPublishingStaterejected,
		"submitted": TeamsAppDefinitionPublishingStatesubmitted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAppDefinitionPublishingState(input)
	return &out, nil
}
