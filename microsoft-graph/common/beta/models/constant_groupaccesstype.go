package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAccessType string

const (
	GroupAccessTypenone    GroupAccessType = "None"
	GroupAccessTypeprivate GroupAccessType = "Private"
	GroupAccessTypepublic  GroupAccessType = "Public"
	GroupAccessTypesecret  GroupAccessType = "Secret"
)

func PossibleValuesForGroupAccessType() []string {
	return []string{
		string(GroupAccessTypenone),
		string(GroupAccessTypeprivate),
		string(GroupAccessTypepublic),
		string(GroupAccessTypesecret),
	}
}

func parseGroupAccessType(input string) (*GroupAccessType, error) {
	vals := map[string]GroupAccessType{
		"none":    GroupAccessTypenone,
		"private": GroupAccessTypeprivate,
		"public":  GroupAccessTypepublic,
		"secret":  GroupAccessTypesecret,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupAccessType(input)
	return &out, nil
}
