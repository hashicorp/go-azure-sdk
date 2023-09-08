package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateConfigurationScheduledInstallDays string

const (
	IosUpdateConfigurationScheduledInstallDaysfriday    IosUpdateConfigurationScheduledInstallDays = "Friday"
	IosUpdateConfigurationScheduledInstallDaysmonday    IosUpdateConfigurationScheduledInstallDays = "Monday"
	IosUpdateConfigurationScheduledInstallDayssaturday  IosUpdateConfigurationScheduledInstallDays = "Saturday"
	IosUpdateConfigurationScheduledInstallDayssunday    IosUpdateConfigurationScheduledInstallDays = "Sunday"
	IosUpdateConfigurationScheduledInstallDaysthursday  IosUpdateConfigurationScheduledInstallDays = "Thursday"
	IosUpdateConfigurationScheduledInstallDaystuesday   IosUpdateConfigurationScheduledInstallDays = "Tuesday"
	IosUpdateConfigurationScheduledInstallDayswednesday IosUpdateConfigurationScheduledInstallDays = "Wednesday"
)

func PossibleValuesForIosUpdateConfigurationScheduledInstallDays() []string {
	return []string{
		string(IosUpdateConfigurationScheduledInstallDaysfriday),
		string(IosUpdateConfigurationScheduledInstallDaysmonday),
		string(IosUpdateConfigurationScheduledInstallDayssaturday),
		string(IosUpdateConfigurationScheduledInstallDayssunday),
		string(IosUpdateConfigurationScheduledInstallDaysthursday),
		string(IosUpdateConfigurationScheduledInstallDaystuesday),
		string(IosUpdateConfigurationScheduledInstallDayswednesday),
	}
}

func parseIosUpdateConfigurationScheduledInstallDays(input string) (*IosUpdateConfigurationScheduledInstallDays, error) {
	vals := map[string]IosUpdateConfigurationScheduledInstallDays{
		"friday":    IosUpdateConfigurationScheduledInstallDaysfriday,
		"monday":    IosUpdateConfigurationScheduledInstallDaysmonday,
		"saturday":  IosUpdateConfigurationScheduledInstallDayssaturday,
		"sunday":    IosUpdateConfigurationScheduledInstallDayssunday,
		"thursday":  IosUpdateConfigurationScheduledInstallDaysthursday,
		"tuesday":   IosUpdateConfigurationScheduledInstallDaystuesday,
		"wednesday": IosUpdateConfigurationScheduledInstallDayswednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateConfigurationScheduledInstallDays(input)
	return &out, nil
}
