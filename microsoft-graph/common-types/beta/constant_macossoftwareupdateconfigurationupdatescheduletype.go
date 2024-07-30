package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfigurationUpdateScheduleType string

const (
	MacOSSoftwareUpdateConfigurationUpdateScheduleType_AlwaysUpdate               MacOSSoftwareUpdateConfigurationUpdateScheduleType = "alwaysUpdate"
	MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows    MacOSSoftwareUpdateConfigurationUpdateScheduleType = "updateDuringTimeWindows"
	MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows MacOSSoftwareUpdateConfigurationUpdateScheduleType = "updateOutsideOfTimeWindows"
)

func PossibleValuesForMacOSSoftwareUpdateConfigurationUpdateScheduleType() []string {
	return []string{
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleType_AlwaysUpdate),
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows),
		string(MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows),
	}
}

func (s *MacOSSoftwareUpdateConfigurationUpdateScheduleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSSoftwareUpdateConfigurationUpdateScheduleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSSoftwareUpdateConfigurationUpdateScheduleType(input string) (*MacOSSoftwareUpdateConfigurationUpdateScheduleType, error) {
	vals := map[string]MacOSSoftwareUpdateConfigurationUpdateScheduleType{
		"alwaysupdate":               MacOSSoftwareUpdateConfigurationUpdateScheduleType_AlwaysUpdate,
		"updateduringtimewindows":    MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateDuringTimeWindows,
		"updateoutsideoftimewindows": MacOSSoftwareUpdateConfigurationUpdateScheduleType_UpdateOutsideOfTimeWindows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSSoftwareUpdateConfigurationUpdateScheduleType(input)
	return &out, nil
}
