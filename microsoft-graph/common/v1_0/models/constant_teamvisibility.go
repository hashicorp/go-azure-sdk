package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamVisibility string

const (
	TeamVisibilityhiddenMembership TeamVisibility = "HiddenMembership"
	TeamVisibilityprivate          TeamVisibility = "Private"
	TeamVisibilitypublic           TeamVisibility = "Public"
)

func PossibleValuesForTeamVisibility() []string {
	return []string{
		string(TeamVisibilityhiddenMembership),
		string(TeamVisibilityprivate),
		string(TeamVisibilitypublic),
	}
}

func parseTeamVisibility(input string) (*TeamVisibility, error) {
	vals := map[string]TeamVisibility{
		"hiddenmembership": TeamVisibilityhiddenMembership,
		"private":          TeamVisibilityprivate,
		"public":           TeamVisibilitypublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamVisibility(input)
	return &out, nil
}
