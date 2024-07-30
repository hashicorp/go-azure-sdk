package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateScheduledInstallScheduledInstallDay string

const (
	WindowsUpdateScheduledInstallScheduledInstallDay_Everyday    WindowsUpdateScheduledInstallScheduledInstallDay = "everyday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Friday      WindowsUpdateScheduledInstallScheduledInstallDay = "friday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Monday      WindowsUpdateScheduledInstallScheduledInstallDay = "monday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Saturday    WindowsUpdateScheduledInstallScheduledInstallDay = "saturday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Sunday      WindowsUpdateScheduledInstallScheduledInstallDay = "sunday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Thursday    WindowsUpdateScheduledInstallScheduledInstallDay = "thursday"
	WindowsUpdateScheduledInstallScheduledInstallDay_Tuesday     WindowsUpdateScheduledInstallScheduledInstallDay = "tuesday"
	WindowsUpdateScheduledInstallScheduledInstallDay_UserDefined WindowsUpdateScheduledInstallScheduledInstallDay = "userDefined"
	WindowsUpdateScheduledInstallScheduledInstallDay_Wednesday   WindowsUpdateScheduledInstallScheduledInstallDay = "wednesday"
)

func PossibleValuesForWindowsUpdateScheduledInstallScheduledInstallDay() []string {
	return []string{
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Everyday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Friday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Monday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Saturday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Sunday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Thursday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Tuesday),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_UserDefined),
		string(WindowsUpdateScheduledInstallScheduledInstallDay_Wednesday),
	}
}

func (s *WindowsUpdateScheduledInstallScheduledInstallDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateScheduledInstallScheduledInstallDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateScheduledInstallScheduledInstallDay(input string) (*WindowsUpdateScheduledInstallScheduledInstallDay, error) {
	vals := map[string]WindowsUpdateScheduledInstallScheduledInstallDay{
		"everyday":    WindowsUpdateScheduledInstallScheduledInstallDay_Everyday,
		"friday":      WindowsUpdateScheduledInstallScheduledInstallDay_Friday,
		"monday":      WindowsUpdateScheduledInstallScheduledInstallDay_Monday,
		"saturday":    WindowsUpdateScheduledInstallScheduledInstallDay_Saturday,
		"sunday":      WindowsUpdateScheduledInstallScheduledInstallDay_Sunday,
		"thursday":    WindowsUpdateScheduledInstallScheduledInstallDay_Thursday,
		"tuesday":     WindowsUpdateScheduledInstallScheduledInstallDay_Tuesday,
		"userdefined": WindowsUpdateScheduledInstallScheduledInstallDay_UserDefined,
		"wednesday":   WindowsUpdateScheduledInstallScheduledInstallDay_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateScheduledInstallScheduledInstallDay(input)
	return &out, nil
}
