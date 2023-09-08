package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationDurationOfEmailToSync string

const (
	Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeDay      Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "OneDay"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeMonth    Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "OneMonth"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeWeek     Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "OneWeek"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSyncthreeDays   Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "ThreeDays"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks    Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "TwoWeeks"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSyncunlimited   Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "Unlimited"
	Windows10EasEmailProfileConfigurationDurationOfEmailToSyncuserDefined Windows10EasEmailProfileConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeDay),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeMonth),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeWeek),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSyncthreeDays),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSyncunlimited),
		string(Windows10EasEmailProfileConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseWindows10EasEmailProfileConfigurationDurationOfEmailToSync(input string) (*Windows10EasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     Windows10EasEmailProfileConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   Windows10EasEmailProfileConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    Windows10EasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   Windows10EasEmailProfileConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": Windows10EasEmailProfileConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
