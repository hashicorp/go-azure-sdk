package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderSystemScanSchedule string

const (
	Windows10GeneralConfigurationDefenderSystemScanScheduleeveryday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "Everyday"
	Windows10GeneralConfigurationDefenderSystemScanSchedulefriday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "Friday"
	Windows10GeneralConfigurationDefenderSystemScanSchedulemonday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "Monday"
	Windows10GeneralConfigurationDefenderSystemScanSchedulesaturday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "Saturday"
	Windows10GeneralConfigurationDefenderSystemScanSchedulesunday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "Sunday"
	Windows10GeneralConfigurationDefenderSystemScanSchedulethursday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "Thursday"
	Windows10GeneralConfigurationDefenderSystemScanScheduletuesday     Windows10GeneralConfigurationDefenderSystemScanSchedule = "Tuesday"
	Windows10GeneralConfigurationDefenderSystemScanScheduleuserDefined Windows10GeneralConfigurationDefenderSystemScanSchedule = "UserDefined"
	Windows10GeneralConfigurationDefenderSystemScanSchedulewednesday   Windows10GeneralConfigurationDefenderSystemScanSchedule = "Wednesday"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderSystemScanSchedule() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderSystemScanScheduleeveryday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulefriday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulemonday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulesaturday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulesunday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulethursday),
		string(Windows10GeneralConfigurationDefenderSystemScanScheduletuesday),
		string(Windows10GeneralConfigurationDefenderSystemScanScheduleuserDefined),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedulewednesday),
	}
}

func parseWindows10GeneralConfigurationDefenderSystemScanSchedule(input string) (*Windows10GeneralConfigurationDefenderSystemScanSchedule, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderSystemScanSchedule{
		"everyday":    Windows10GeneralConfigurationDefenderSystemScanScheduleeveryday,
		"friday":      Windows10GeneralConfigurationDefenderSystemScanSchedulefriday,
		"monday":      Windows10GeneralConfigurationDefenderSystemScanSchedulemonday,
		"saturday":    Windows10GeneralConfigurationDefenderSystemScanSchedulesaturday,
		"sunday":      Windows10GeneralConfigurationDefenderSystemScanSchedulesunday,
		"thursday":    Windows10GeneralConfigurationDefenderSystemScanSchedulethursday,
		"tuesday":     Windows10GeneralConfigurationDefenderSystemScanScheduletuesday,
		"userdefined": Windows10GeneralConfigurationDefenderSystemScanScheduleuserDefined,
		"wednesday":   Windows10GeneralConfigurationDefenderSystemScanSchedulewednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderSystemScanSchedule(input)
	return &out, nil
}
