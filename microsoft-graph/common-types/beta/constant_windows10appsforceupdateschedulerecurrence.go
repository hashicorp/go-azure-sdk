package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10AppsForceUpdateScheduleRecurrence string

const (
	Windows10AppsForceUpdateScheduleRecurrence_Daily   Windows10AppsForceUpdateScheduleRecurrence = "daily"
	Windows10AppsForceUpdateScheduleRecurrence_Monthly Windows10AppsForceUpdateScheduleRecurrence = "monthly"
	Windows10AppsForceUpdateScheduleRecurrence_None    Windows10AppsForceUpdateScheduleRecurrence = "none"
	Windows10AppsForceUpdateScheduleRecurrence_Weekly  Windows10AppsForceUpdateScheduleRecurrence = "weekly"
)

func PossibleValuesForWindows10AppsForceUpdateScheduleRecurrence() []string {
	return []string{
		string(Windows10AppsForceUpdateScheduleRecurrence_Daily),
		string(Windows10AppsForceUpdateScheduleRecurrence_Monthly),
		string(Windows10AppsForceUpdateScheduleRecurrence_None),
		string(Windows10AppsForceUpdateScheduleRecurrence_Weekly),
	}
}

func (s *Windows10AppsForceUpdateScheduleRecurrence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10AppsForceUpdateScheduleRecurrence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10AppsForceUpdateScheduleRecurrence(input string) (*Windows10AppsForceUpdateScheduleRecurrence, error) {
	vals := map[string]Windows10AppsForceUpdateScheduleRecurrence{
		"daily":   Windows10AppsForceUpdateScheduleRecurrence_Daily,
		"monthly": Windows10AppsForceUpdateScheduleRecurrence_Monthly,
		"none":    Windows10AppsForceUpdateScheduleRecurrence_None,
		"weekly":  Windows10AppsForceUpdateScheduleRecurrence_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10AppsForceUpdateScheduleRecurrence(input)
	return &out, nil
}
