package groupuser

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupType string

const (
	GroupTypeCustom   GroupType = "custom"
	GroupTypeExternal GroupType = "external"
	GroupTypeSystem   GroupType = "system"
)

func PossibleValuesForGroupType() []string {
	return []string{
		string(GroupTypeCustom),
		string(GroupTypeExternal),
		string(GroupTypeSystem),
	}
}

func parseGroupType(input string) (*GroupType, error) {
	vals := map[string]GroupType{
		"custom":   GroupTypeCustom,
		"external": GroupTypeExternal,
		"system":   GroupTypeSystem,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupType(input)
	return &out, nil
}

type UserState string

const (
	UserStateActive  UserState = "active"
	UserStateBlocked UserState = "blocked"
	UserStateDeleted UserState = "deleted"
	UserStatePending UserState = "pending"
)

func PossibleValuesForUserState() []string {
	return []string{
		string(UserStateActive),
		string(UserStateBlocked),
		string(UserStateDeleted),
		string(UserStatePending),
	}
}

func parseUserState(input string) (*UserState, error) {
	vals := map[string]UserState{
		"active":  UserStateActive,
		"blocked": UserStateBlocked,
		"deleted": UserStateDeleted,
		"pending": UserStatePending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserState(input)
	return &out, nil
}
