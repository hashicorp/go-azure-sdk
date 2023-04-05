package softwareupdateconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinuxUpdateClasses string

const (
	LinuxUpdateClassesCritical     LinuxUpdateClasses = "Critical"
	LinuxUpdateClassesOther        LinuxUpdateClasses = "Other"
	LinuxUpdateClassesSecurity     LinuxUpdateClasses = "Security"
	LinuxUpdateClassesUnclassified LinuxUpdateClasses = "Unclassified"
)

func PossibleValuesForLinuxUpdateClasses() []string {
	return []string{
		string(LinuxUpdateClassesCritical),
		string(LinuxUpdateClassesOther),
		string(LinuxUpdateClassesSecurity),
		string(LinuxUpdateClassesUnclassified),
	}
}

type OperatingSystemType string

const (
	OperatingSystemTypeLinux   OperatingSystemType = "Linux"
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
)

func PossibleValuesForOperatingSystemType() []string {
	return []string{
		string(OperatingSystemTypeLinux),
		string(OperatingSystemTypeWindows),
	}
}

type ScheduleDay string

const (
	ScheduleDayFriday    ScheduleDay = "Friday"
	ScheduleDayMonday    ScheduleDay = "Monday"
	ScheduleDaySaturday  ScheduleDay = "Saturday"
	ScheduleDaySunday    ScheduleDay = "Sunday"
	ScheduleDayThursday  ScheduleDay = "Thursday"
	ScheduleDayTuesday   ScheduleDay = "Tuesday"
	ScheduleDayWednesday ScheduleDay = "Wednesday"
)

func PossibleValuesForScheduleDay() []string {
	return []string{
		string(ScheduleDayFriday),
		string(ScheduleDayMonday),
		string(ScheduleDaySaturday),
		string(ScheduleDaySunday),
		string(ScheduleDayThursday),
		string(ScheduleDayTuesday),
		string(ScheduleDayWednesday),
	}
}

type ScheduleFrequency string

const (
	ScheduleFrequencyDay     ScheduleFrequency = "Day"
	ScheduleFrequencyHour    ScheduleFrequency = "Hour"
	ScheduleFrequencyMinute  ScheduleFrequency = "Minute"
	ScheduleFrequencyMonth   ScheduleFrequency = "Month"
	ScheduleFrequencyOneTime ScheduleFrequency = "OneTime"
	ScheduleFrequencyWeek    ScheduleFrequency = "Week"
)

func PossibleValuesForScheduleFrequency() []string {
	return []string{
		string(ScheduleFrequencyDay),
		string(ScheduleFrequencyHour),
		string(ScheduleFrequencyMinute),
		string(ScheduleFrequencyMonth),
		string(ScheduleFrequencyOneTime),
		string(ScheduleFrequencyWeek),
	}
}

type TagOperators string

const (
	TagOperatorsAll TagOperators = "All"
	TagOperatorsAny TagOperators = "Any"
)

func PossibleValuesForTagOperators() []string {
	return []string{
		string(TagOperatorsAll),
		string(TagOperatorsAny),
	}
}

type WindowsUpdateClasses string

const (
	WindowsUpdateClassesCritical     WindowsUpdateClasses = "Critical"
	WindowsUpdateClassesDefinition   WindowsUpdateClasses = "Definition"
	WindowsUpdateClassesFeaturePack  WindowsUpdateClasses = "FeaturePack"
	WindowsUpdateClassesSecurity     WindowsUpdateClasses = "Security"
	WindowsUpdateClassesServicePack  WindowsUpdateClasses = "ServicePack"
	WindowsUpdateClassesTools        WindowsUpdateClasses = "Tools"
	WindowsUpdateClassesUnclassified WindowsUpdateClasses = "Unclassified"
	WindowsUpdateClassesUpdateRollup WindowsUpdateClasses = "UpdateRollup"
	WindowsUpdateClassesUpdates      WindowsUpdateClasses = "Updates"
)

func PossibleValuesForWindowsUpdateClasses() []string {
	return []string{
		string(WindowsUpdateClassesCritical),
		string(WindowsUpdateClassesDefinition),
		string(WindowsUpdateClassesFeaturePack),
		string(WindowsUpdateClassesSecurity),
		string(WindowsUpdateClassesServicePack),
		string(WindowsUpdateClassesTools),
		string(WindowsUpdateClassesUnclassified),
		string(WindowsUpdateClassesUpdateRollup),
		string(WindowsUpdateClassesUpdates),
	}
}
