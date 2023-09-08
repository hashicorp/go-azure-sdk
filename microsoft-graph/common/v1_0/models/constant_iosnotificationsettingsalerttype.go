package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosNotificationSettingsAlertType string

const (
	IosNotificationSettingsAlertTypebanner        IosNotificationSettingsAlertType = "Banner"
	IosNotificationSettingsAlertTypedeviceDefault IosNotificationSettingsAlertType = "DeviceDefault"
	IosNotificationSettingsAlertTypemodal         IosNotificationSettingsAlertType = "Modal"
	IosNotificationSettingsAlertTypenone          IosNotificationSettingsAlertType = "None"
)

func PossibleValuesForIosNotificationSettingsAlertType() []string {
	return []string{
		string(IosNotificationSettingsAlertTypebanner),
		string(IosNotificationSettingsAlertTypedeviceDefault),
		string(IosNotificationSettingsAlertTypemodal),
		string(IosNotificationSettingsAlertTypenone),
	}
}

func parseIosNotificationSettingsAlertType(input string) (*IosNotificationSettingsAlertType, error) {
	vals := map[string]IosNotificationSettingsAlertType{
		"banner":        IosNotificationSettingsAlertTypebanner,
		"devicedefault": IosNotificationSettingsAlertTypedeviceDefault,
		"modal":         IosNotificationSettingsAlertTypemodal,
		"none":          IosNotificationSettingsAlertTypenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosNotificationSettingsAlertType(input)
	return &out, nil
}
