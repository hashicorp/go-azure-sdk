package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync string

const (
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeDay      AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "OneDay"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeMonth    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "OneMonth"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeWeek     AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "OneWeek"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncthreeDays   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "ThreeDays"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "TwoWeeks"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncunlimited   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "Unlimited"
	AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncuserDefined AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeDay),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeMonth),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeWeek),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncthreeDays),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncunlimited),
		string(AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync(input string) (*AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
