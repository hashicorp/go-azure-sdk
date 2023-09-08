package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync string

const (
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeDay      WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "OneDay"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeMonth    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "OneMonth"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeWeek     WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "OneWeek"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncthreeDays   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "ThreeDays"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynctwoWeeks    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "TwoWeeks"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncunlimited   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "Unlimited"
	WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncuserDefined WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeDay),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeMonth),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeWeek),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncthreeDays),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynctwoWeeks),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncunlimited),
		string(WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseWindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync(input string) (*WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
