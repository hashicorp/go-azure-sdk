package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateScheduleDayofWeek string

const (
	WindowsKioskForceUpdateScheduleDayofWeek_Friday    WindowsKioskForceUpdateScheduleDayofWeek = "friday"
	WindowsKioskForceUpdateScheduleDayofWeek_Monday    WindowsKioskForceUpdateScheduleDayofWeek = "monday"
	WindowsKioskForceUpdateScheduleDayofWeek_Saturday  WindowsKioskForceUpdateScheduleDayofWeek = "saturday"
	WindowsKioskForceUpdateScheduleDayofWeek_Sunday    WindowsKioskForceUpdateScheduleDayofWeek = "sunday"
	WindowsKioskForceUpdateScheduleDayofWeek_Thursday  WindowsKioskForceUpdateScheduleDayofWeek = "thursday"
	WindowsKioskForceUpdateScheduleDayofWeek_Tuesday   WindowsKioskForceUpdateScheduleDayofWeek = "tuesday"
	WindowsKioskForceUpdateScheduleDayofWeek_Wednesday WindowsKioskForceUpdateScheduleDayofWeek = "wednesday"
)

func PossibleValuesForWindowsKioskForceUpdateScheduleDayofWeek() []string {
	return []string{
		string(WindowsKioskForceUpdateScheduleDayofWeek_Friday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Monday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Saturday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Sunday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Thursday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Tuesday),
		string(WindowsKioskForceUpdateScheduleDayofWeek_Wednesday),
	}
}

func (s *WindowsKioskForceUpdateScheduleDayofWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskForceUpdateScheduleDayofWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskForceUpdateScheduleDayofWeek(input string) (*WindowsKioskForceUpdateScheduleDayofWeek, error) {
	vals := map[string]WindowsKioskForceUpdateScheduleDayofWeek{
		"friday":    WindowsKioskForceUpdateScheduleDayofWeek_Friday,
		"monday":    WindowsKioskForceUpdateScheduleDayofWeek_Monday,
		"saturday":  WindowsKioskForceUpdateScheduleDayofWeek_Saturday,
		"sunday":    WindowsKioskForceUpdateScheduleDayofWeek_Sunday,
		"thursday":  WindowsKioskForceUpdateScheduleDayofWeek_Thursday,
		"tuesday":   WindowsKioskForceUpdateScheduleDayofWeek_Tuesday,
		"wednesday": WindowsKioskForceUpdateScheduleDayofWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskForceUpdateScheduleDayofWeek(input)
	return &out, nil
}
