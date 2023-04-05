package namespaces

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

type NamespaceType string

const (
	NamespaceTypeMessaging       NamespaceType = "Messaging"
	NamespaceTypeNotificationHub NamespaceType = "NotificationHub"
)

func PossibleValuesForNamespaceType() []string {
	return []string{
		string(NamespaceTypeMessaging),
		string(NamespaceTypeNotificationHub),
	}
}

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNameFree     SkuName = "Free"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNameFree),
		string(SkuNameStandard),
	}
}
