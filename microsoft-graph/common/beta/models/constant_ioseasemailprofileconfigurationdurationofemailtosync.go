package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationDurationOfEmailToSync string

const (
	IosEasEmailProfileConfigurationDurationOfEmailToSynconeDay      IosEasEmailProfileConfigurationDurationOfEmailToSync = "OneDay"
	IosEasEmailProfileConfigurationDurationOfEmailToSynconeMonth    IosEasEmailProfileConfigurationDurationOfEmailToSync = "OneMonth"
	IosEasEmailProfileConfigurationDurationOfEmailToSynconeWeek     IosEasEmailProfileConfigurationDurationOfEmailToSync = "OneWeek"
	IosEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays   IosEasEmailProfileConfigurationDurationOfEmailToSync = "ThreeDays"
	IosEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks    IosEasEmailProfileConfigurationDurationOfEmailToSync = "TwoWeeks"
	IosEasEmailProfileConfigurationDurationOfEmailToSyncunlimited   IosEasEmailProfileConfigurationDurationOfEmailToSync = "Unlimited"
	IosEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined IosEasEmailProfileConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForIosEasEmailProfileConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(IosEasEmailProfileConfigurationDurationOfEmailToSynconeDay),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSynconeMonth),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSynconeWeek),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSyncunlimited),
		string(IosEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseIosEasEmailProfileConfigurationDurationOfEmailToSync(input string) (*IosEasEmailProfileConfigurationDurationOfEmailToSync, error) {
	vals := map[string]IosEasEmailProfileConfigurationDurationOfEmailToSync{
		"oneday":      IosEasEmailProfileConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    IosEasEmailProfileConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     IosEasEmailProfileConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   IosEasEmailProfileConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    IosEasEmailProfileConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   IosEasEmailProfileConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": IosEasEmailProfileConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
