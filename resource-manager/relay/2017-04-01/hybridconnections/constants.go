package hybridconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRights string

const (
	AccessRightsListen AccessRights = "Listen"
	AccessRightsManage AccessRights = "Manage"
	AccessRightsSend   AccessRights = "Send"
)

func PossibleValuesForAccessRights() []string {
	return []string{
		string(AccessRightsListen),
		string(AccessRightsManage),
		string(AccessRightsSend),
	}
}

type KeyType string

const (
	KeyTypePrimaryKey   KeyType = "PrimaryKey"
	KeyTypeSecondaryKey KeyType = "SecondaryKey"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimaryKey),
		string(KeyTypeSecondaryKey),
	}
}
