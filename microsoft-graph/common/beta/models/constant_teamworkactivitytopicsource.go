package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkActivityTopicSource string

const (
	TeamworkActivityTopicSourceentityUrl TeamworkActivityTopicSource = "EntityUrl"
	TeamworkActivityTopicSourcetext      TeamworkActivityTopicSource = "Text"
)

func PossibleValuesForTeamworkActivityTopicSource() []string {
	return []string{
		string(TeamworkActivityTopicSourceentityUrl),
		string(TeamworkActivityTopicSourcetext),
	}
}

func parseTeamworkActivityTopicSource(input string) (*TeamworkActivityTopicSource, error) {
	vals := map[string]TeamworkActivityTopicSource{
		"entityurl": TeamworkActivityTopicSourceentityUrl,
		"text":      TeamworkActivityTopicSourcetext,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkActivityTopicSource(input)
	return &out, nil
}
