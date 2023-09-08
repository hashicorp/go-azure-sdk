package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateConfigurationUpdateScheduleType string

const (
	IosUpdateConfigurationUpdateScheduleTypealwaysUpdate               IosUpdateConfigurationUpdateScheduleType = "AlwaysUpdate"
	IosUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows    IosUpdateConfigurationUpdateScheduleType = "UpdateDuringTimeWindows"
	IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfActiveHours IosUpdateConfigurationUpdateScheduleType = "UpdateOutsideOfActiveHours"
	IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows IosUpdateConfigurationUpdateScheduleType = "UpdateOutsideOfTimeWindows"
)

func PossibleValuesForIosUpdateConfigurationUpdateScheduleType() []string {
	return []string{
		string(IosUpdateConfigurationUpdateScheduleTypealwaysUpdate),
		string(IosUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows),
		string(IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfActiveHours),
		string(IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows),
	}
}

func parseIosUpdateConfigurationUpdateScheduleType(input string) (*IosUpdateConfigurationUpdateScheduleType, error) {
	vals := map[string]IosUpdateConfigurationUpdateScheduleType{
		"alwaysupdate":               IosUpdateConfigurationUpdateScheduleTypealwaysUpdate,
		"updateduringtimewindows":    IosUpdateConfigurationUpdateScheduleTypeupdateDuringTimeWindows,
		"updateoutsideofactivehours": IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfActiveHours,
		"updateoutsideoftimewindows": IosUpdateConfigurationUpdateScheduleTypeupdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateConfigurationUpdateScheduleType(input)
	return &out, nil
}
