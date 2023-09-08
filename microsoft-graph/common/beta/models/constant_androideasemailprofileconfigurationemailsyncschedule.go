package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationEmailSyncSchedule string

const (
	AndroidEasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive AndroidEasEmailProfileConfigurationEmailSyncSchedule = "AsMessagesArrive"
	AndroidEasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage   AndroidEasEmailProfileConfigurationEmailSyncSchedule = "BasedOnMyUsage"
	AndroidEasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes   AndroidEasEmailProfileConfigurationEmailSyncSchedule = "FifteenMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncSchedulemanual           AndroidEasEmailProfileConfigurationEmailSyncSchedule = "Manual"
	AndroidEasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes     AndroidEasEmailProfileConfigurationEmailSyncSchedule = "SixtyMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes    AndroidEasEmailProfileConfigurationEmailSyncSchedule = "ThirtyMinutes"
	AndroidEasEmailProfileConfigurationEmailSyncScheduleuserDefined      AndroidEasEmailProfileConfigurationEmailSyncSchedule = "UserDefined"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedulemanual),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes),
		string(AndroidEasEmailProfileConfigurationEmailSyncScheduleuserDefined),
	}
}

func parseAndroidEasEmailProfileConfigurationEmailSyncSchedule(input string) (*AndroidEasEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": AndroidEasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive,
		"basedonmyusage":   AndroidEasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage,
		"fifteenminutes":   AndroidEasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes,
		"manual":           AndroidEasEmailProfileConfigurationEmailSyncSchedulemanual,
		"sixtyminutes":     AndroidEasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes,
		"thirtyminutes":    AndroidEasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes,
		"userdefined":      AndroidEasEmailProfileConfigurationEmailSyncScheduleuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
