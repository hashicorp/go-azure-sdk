package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateConfigurationUpdateScheduleType string

const (
	IosUpdateConfigurationUpdateScheduleType_AlwaysUpdate               IosUpdateConfigurationUpdateScheduleType = "alwaysUpdate"
	IosUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows    IosUpdateConfigurationUpdateScheduleType = "updateDuringTimeWindows"
	IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfActiveHours IosUpdateConfigurationUpdateScheduleType = "updateOutsideOfActiveHours"
	IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows IosUpdateConfigurationUpdateScheduleType = "updateOutsideOfTimeWindows"
)

func PossibleValuesForIosUpdateConfigurationUpdateScheduleType() []string {
	return []string{
		string(IosUpdateConfigurationUpdateScheduleType_AlwaysUpdate),
		string(IosUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows),
		string(IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfActiveHours),
		string(IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows),
	}
}

func (s *IosUpdateConfigurationUpdateScheduleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosUpdateConfigurationUpdateScheduleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosUpdateConfigurationUpdateScheduleType(input string) (*IosUpdateConfigurationUpdateScheduleType, error) {
	vals := map[string]IosUpdateConfigurationUpdateScheduleType{
		"alwaysupdate":               IosUpdateConfigurationUpdateScheduleType_AlwaysUpdate,
		"updateduringtimewindows":    IosUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows,
		"updateoutsideofactivehours": IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfActiveHours,
		"updateoutsideoftimewindows": IosUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosUpdateConfigurationUpdateScheduleType(input)
	return &out, nil
}
