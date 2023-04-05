package hybridrunbookworkergroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTypeEnum string

const (
	GroupTypeEnumSystem GroupTypeEnum = "System"
	GroupTypeEnumUser   GroupTypeEnum = "User"
)

func PossibleValuesForGroupTypeEnum() []string {
	return []string{
		string(GroupTypeEnumSystem),
		string(GroupTypeEnumUser),
	}
}
