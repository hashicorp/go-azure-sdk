package groupuser

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
