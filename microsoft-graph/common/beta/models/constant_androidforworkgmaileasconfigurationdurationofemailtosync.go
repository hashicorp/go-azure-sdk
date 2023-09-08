package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationDurationOfEmailToSync string

const (
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeDay      AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "OneDay"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeMonth    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "OneMonth"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeWeek     AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "OneWeek"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncthreeDays   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "ThreeDays"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSynctwoWeeks    AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "TwoWeeks"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncunlimited   AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "Unlimited"
	AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncuserDefined AndroidForWorkGmailEasConfigurationDurationOfEmailToSync = "UserDefined"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationDurationOfEmailToSync() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeDay),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeMonth),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeWeek),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncthreeDays),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSynctwoWeeks),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncunlimited),
		string(AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncuserDefined),
	}
}

func parseAndroidForWorkGmailEasConfigurationDurationOfEmailToSync(input string) (*AndroidForWorkGmailEasConfigurationDurationOfEmailToSync, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationDurationOfEmailToSync{
		"oneday":      AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeDay,
		"onemonth":    AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeMonth,
		"oneweek":     AndroidForWorkGmailEasConfigurationDurationOfEmailToSynconeWeek,
		"threedays":   AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncthreeDays,
		"twoweeks":    AndroidForWorkGmailEasConfigurationDurationOfEmailToSynctwoWeeks,
		"unlimited":   AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncunlimited,
		"userdefined": AndroidForWorkGmailEasConfigurationDurationOfEmailToSyncuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationDurationOfEmailToSync(input)
	return &out, nil
}
