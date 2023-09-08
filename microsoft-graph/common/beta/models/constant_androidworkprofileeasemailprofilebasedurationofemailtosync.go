package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync string

const (
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeDay      AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "OneDay"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeMonth    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "OneMonth"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeWeek     AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "OneWeek"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncthreeDays   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "ThreeDays"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynctwoWeeks    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "TwoWeeks"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncunlimited   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "Unlimited"
	AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncuserDefined AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeDay),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeMonth),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeWeek),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncthreeDays),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynctwoWeeks),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncunlimited),
		string(AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync(input string) (*AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeDay,
		"onemonth":    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynconeWeek,
		"threedays":   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncunlimited,
		"userdefined": AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseDurationOfEmailToSync(input)
	return &out, nil
}
