package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync string

const (
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeDay      AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "OneDay"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeMonth    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "OneMonth"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeWeek     AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "OneWeek"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncthreeDays   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "ThreeDays"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "TwoWeeks"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncunlimited   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "Unlimited"
	AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncuserDefined AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeDay),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeMonth),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeWeek),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncthreeDays),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncunlimited),
		string(AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync(input string) (*AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
