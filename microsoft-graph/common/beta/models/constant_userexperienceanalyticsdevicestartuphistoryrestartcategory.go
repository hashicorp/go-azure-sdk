package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceStartupHistoryRestartCategory string

const (
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryblueScreen            UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "BlueScreen"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategorybootError             UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "BootError"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategorylongPowerButtonPress  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "LongPowerButtonPress"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithUpdate     UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "RestartWithUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithoutUpdate  UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "RestartWithoutUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithUpdate    UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "ShutdownWithUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithoutUpdate UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "ShutdownWithoutUpdate"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryunknown               UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "Unknown"
	UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryupdate                UserExperienceAnalyticsDeviceStartupHistoryRestartCategory = "Update"
)

func PossibleValuesForUserExperienceAnalyticsDeviceStartupHistoryRestartCategory() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryblueScreen),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategorybootError),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategorylongPowerButtonPress),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithoutUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithoutUpdate),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryunknown),
		string(UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryupdate),
	}
}

func parseUserExperienceAnalyticsDeviceStartupHistoryRestartCategory(input string) (*UserExperienceAnalyticsDeviceStartupHistoryRestartCategory, error) {
	vals := map[string]UserExperienceAnalyticsDeviceStartupHistoryRestartCategory{
		"bluescreen":            UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryblueScreen,
		"booterror":             UserExperienceAnalyticsDeviceStartupHistoryRestartCategorybootError,
		"longpowerbuttonpress":  UserExperienceAnalyticsDeviceStartupHistoryRestartCategorylongPowerButtonPress,
		"restartwithupdate":     UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithUpdate,
		"restartwithoutupdate":  UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryrestartWithoutUpdate,
		"shutdownwithupdate":    UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithUpdate,
		"shutdownwithoutupdate": UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryshutdownWithoutUpdate,
		"unknown":               UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryunknown,
		"update":                UserExperienceAnalyticsDeviceStartupHistoryRestartCategoryupdate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceStartupHistoryRestartCategory(input)
	return &out, nil
}
