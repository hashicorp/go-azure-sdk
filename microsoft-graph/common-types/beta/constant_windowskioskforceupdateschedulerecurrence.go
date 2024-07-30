package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateScheduleRecurrence string

const (
	WindowsKioskForceUpdateScheduleRecurrence_Daily   WindowsKioskForceUpdateScheduleRecurrence = "daily"
	WindowsKioskForceUpdateScheduleRecurrence_Monthly WindowsKioskForceUpdateScheduleRecurrence = "monthly"
	WindowsKioskForceUpdateScheduleRecurrence_None    WindowsKioskForceUpdateScheduleRecurrence = "none"
	WindowsKioskForceUpdateScheduleRecurrence_Weekly  WindowsKioskForceUpdateScheduleRecurrence = "weekly"
)

func PossibleValuesForWindowsKioskForceUpdateScheduleRecurrence() []string {
	return []string{
		string(WindowsKioskForceUpdateScheduleRecurrence_Daily),
		string(WindowsKioskForceUpdateScheduleRecurrence_Monthly),
		string(WindowsKioskForceUpdateScheduleRecurrence_None),
		string(WindowsKioskForceUpdateScheduleRecurrence_Weekly),
	}
}

func (s *WindowsKioskForceUpdateScheduleRecurrence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsKioskForceUpdateScheduleRecurrence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsKioskForceUpdateScheduleRecurrence(input string) (*WindowsKioskForceUpdateScheduleRecurrence, error) {
	vals := map[string]WindowsKioskForceUpdateScheduleRecurrence{
		"daily":   WindowsKioskForceUpdateScheduleRecurrence_Daily,
		"monthly": WindowsKioskForceUpdateScheduleRecurrence_Monthly,
		"none":    WindowsKioskForceUpdateScheduleRecurrence_None,
		"weekly":  WindowsKioskForceUpdateScheduleRecurrence_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsKioskForceUpdateScheduleRecurrence(input)
	return &out, nil
}
