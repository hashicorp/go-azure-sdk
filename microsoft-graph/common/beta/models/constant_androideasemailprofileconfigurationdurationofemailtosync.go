package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationDurationOfEmailToSync string

const (
	AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeDay      AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "OneDay"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeMonth    AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "OneMonth"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeWeek     AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "OneWeek"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays   AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "ThreeDays"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks    AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "TwoWeeks"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSyncunlimited   AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "Unlimited"
	AndroidEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined AndroidEasEmailProfileConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeDay),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeMonth),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeWeek),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSyncunlimited),
		string(AndroidEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidEasEmailProfileConfigurationDurationOfEmailToSync(input string) (*AndroidEasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidEasEmailProfileConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   AndroidEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidEasEmailProfileConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": AndroidEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
