package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync string

const (
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeDay      AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "OneDay"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeMonth    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "OneMonth"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeWeek     AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "OneWeek"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncthreeDays   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "ThreeDays"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynctwoWeeks    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "TwoWeeks"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncunlimited   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "Unlimited"
	AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncuserDefined AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeDay),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeMonth),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeWeek),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncthreeDays),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynctwoWeeks),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncunlimited),
		string(AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidForWorkEasEmailProfileBaseDurationOfEmailToSync(input string) (*AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync{
		"oneday":      AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeDay,
		"onemonth":    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynconeWeek,
		"threedays":   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidForWorkEasEmailProfileBaseDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncunlimited,
		"userdefined": AndroidForWorkEasEmailProfileBaseDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseDurationOfEmailToSync(input)
	return &out, nil
}
