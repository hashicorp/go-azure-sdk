package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationEmailSyncSchedule string

const (
	Windows10EasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive Windows10EasEmailProfileConfigurationEmailSyncSchedule = "AsMessagesArrive"
	Windows10EasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage   Windows10EasEmailProfileConfigurationEmailSyncSchedule = "BasedOnMyUsage"
	Windows10EasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes   Windows10EasEmailProfileConfigurationEmailSyncSchedule = "FifteenMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncSchedulemanual           Windows10EasEmailProfileConfigurationEmailSyncSchedule = "Manual"
	Windows10EasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes     Windows10EasEmailProfileConfigurationEmailSyncSchedule = "SixtyMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes    Windows10EasEmailProfileConfigurationEmailSyncSchedule = "ThirtyMinutes"
	Windows10EasEmailProfileConfigurationEmailSyncScheduleuserDefined      Windows10EasEmailProfileConfigurationEmailSyncSchedule = "UserDefined"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedulemanual),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes),
		string(Windows10EasEmailProfileConfigurationEmailSyncScheduleuserDefined),
	}
}

func parseWindows10EasEmailProfileConfigurationEmailSyncSchedule(input string) (*Windows10EasEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": Windows10EasEmailProfileConfigurationEmailSyncScheduleasMessagesArrive,
		"basedonmyusage":   Windows10EasEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage,
		"fifteenminutes":   Windows10EasEmailProfileConfigurationEmailSyncSchedulefifteenMinutes,
		"manual":           Windows10EasEmailProfileConfigurationEmailSyncSchedulemanual,
		"sixtyminutes":     Windows10EasEmailProfileConfigurationEmailSyncSchedulesixtyMinutes,
		"thirtyminutes":    Windows10EasEmailProfileConfigurationEmailSyncSchedulethirtyMinutes,
		"userdefined":      Windows10EasEmailProfileConfigurationEmailSyncScheduleuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
