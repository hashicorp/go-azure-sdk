package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule string

const (
	WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleasMessagesArrive WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "AsMessagesArrive"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "BasedOnMyUsage"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulefifteenMinutes   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "FifteenMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulemanual           WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "Manual"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulesixtyMinutes     WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "SixtyMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulethirtyMinutes    WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "ThirtyMinutes"
	WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleuserDefined      WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule = "UserDefined"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleasMessagesArrive),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulefifteenMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulemanual),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulesixtyMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulethirtyMinutes),
		string(WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleuserDefined),
	}
}

func parseWindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule(input string) (*WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule{
		"asmessagesarrive": WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleasMessagesArrive,
		"basedonmyusage":   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulebasedOnMyUsage,
		"fifteenminutes":   WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulefifteenMinutes,
		"manual":           WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulemanual,
		"sixtyminutes":     WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulesixtyMinutes,
		"thirtyminutes":    WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedulethirtyMinutes,
		"userdefined":      WindowsPhoneEASEmailProfileConfigurationEmailSyncScheduleuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationEmailSyncSchedule(input)
	return &out, nil
}
