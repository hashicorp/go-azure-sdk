package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync string

const (
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeDay      AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "OneDay"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeMonth    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "OneMonth"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeWeek     AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "OneWeek"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncthreeDays   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "ThreeDays"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynctwoWeeks    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "TwoWeeks"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncunlimited   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "Unlimited"
	AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncuserDefined AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeDay),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeMonth),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeWeek),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncthreeDays),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynctwoWeeks),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncunlimited),
		string(AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync(input string) (*AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
