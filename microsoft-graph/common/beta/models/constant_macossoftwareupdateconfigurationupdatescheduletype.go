package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationUpdateScheduleType string

const (
	MacOSSoftwareUpdateConfigurationUpdateScheduleTypealwaysUpdate               MacOSSoftwareUpdateConfigurationUpdateScheduleType = "AlwaysUpdate"
	MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows    MacOSSoftwareUpdateConfigurationUpdateScheduleType = "UpdateDuringTimeWindows"
	MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows MacOSSoftwareUpdateConfigurationUpdateScheduleType = "UpdateOutsideOfTimeWindows"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationUpdateScheduleType() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleTypealwaysUpdate),
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows),
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows),
	}
}

func parseMacOSSoftwareUpdateConfigurationUpdateScheduleType(input string) (*MacOSSoftwareUpdateConfigurationUpdateScheduleType, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationUpdateScheduleType{
		"alwaysupdate":               MacOSSoftwareUpdateConfigurationUpdateScheduleTypealwaysUpdate,
		"updateduringtimewindows":    MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows,
		"updateoutsideoftimewindows": MacOSSoftwareUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationUpdateScheduleType(input)
	return &out, nil
}
