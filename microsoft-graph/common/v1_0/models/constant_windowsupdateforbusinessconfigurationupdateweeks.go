package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUpdateWeeks string

const (
	WindowsUpdateForBusinessConfigurationUpdateWeekseveryWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "EveryWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeksfirstWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "FirstWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeksfourthWeek  WindowsUpdateForBusinessConfigurationUpdateWeeks = "FourthWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeekssecondWeek  WindowsUpdateForBusinessConfigurationUpdateWeeks = "SecondWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeksthirdWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "ThirdWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeksuserDefined WindowsUpdateForBusinessConfigurationUpdateWeeks = "UserDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUpdateWeeks() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUpdateWeekseveryWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeksfirstWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeksfourthWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeekssecondWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeksthirdWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeksuserDefined),
	}
}

func parseWindowsUpdateForBusinessConfigurationUpdateWeeks(input string) (*WindowsUpdateForBusinessConfigurationUpdateWeeks, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUpdateWeeks{
		"everyweek":   WindowsUpdateForBusinessConfigurationUpdateWeekseveryWeek,
		"firstweek":   WindowsUpdateForBusinessConfigurationUpdateWeeksfirstWeek,
		"fourthweek":  WindowsUpdateForBusinessConfigurationUpdateWeeksfourthWeek,
		"secondweek":  WindowsUpdateForBusinessConfigurationUpdateWeekssecondWeek,
		"thirdweek":   WindowsUpdateForBusinessConfigurationUpdateWeeksthirdWeek,
		"userdefined": WindowsUpdateForBusinessConfigurationUpdateWeeksuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUpdateWeeks(input)
	return &out, nil
}
