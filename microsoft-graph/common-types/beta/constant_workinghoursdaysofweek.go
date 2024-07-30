package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkingHoursDaysOfWeek string

const (
	WorkingHoursDaysOfWeek_Friday    WorkingHoursDaysOfWeek = "friday"
	WorkingHoursDaysOfWeek_Monday    WorkingHoursDaysOfWeek = "monday"
	WorkingHoursDaysOfWeek_Saturday  WorkingHoursDaysOfWeek = "saturday"
	WorkingHoursDaysOfWeek_Sunday    WorkingHoursDaysOfWeek = "sunday"
	WorkingHoursDaysOfWeek_Thursday  WorkingHoursDaysOfWeek = "thursday"
	WorkingHoursDaysOfWeek_Tuesday   WorkingHoursDaysOfWeek = "tuesday"
	WorkingHoursDaysOfWeek_Wednesday WorkingHoursDaysOfWeek = "wednesday"
)

func PossibleValuesForWorkingHoursDaysOfWeek() []string {
	return []string{
		string(WorkingHoursDaysOfWeek_Friday),
		string(WorkingHoursDaysOfWeek_Monday),
		string(WorkingHoursDaysOfWeek_Saturday),
		string(WorkingHoursDaysOfWeek_Sunday),
		string(WorkingHoursDaysOfWeek_Thursday),
		string(WorkingHoursDaysOfWeek_Tuesday),
		string(WorkingHoursDaysOfWeek_Wednesday),
	}
}

func (s *WorkingHoursDaysOfWeek) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWorkingHoursDaysOfWeek(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWorkingHoursDaysOfWeek(input string) (*WorkingHoursDaysOfWeek, error) {
	vals := map[string]WorkingHoursDaysOfWeek{
		"friday":    WorkingHoursDaysOfWeek_Friday,
		"monday":    WorkingHoursDaysOfWeek_Monday,
		"saturday":  WorkingHoursDaysOfWeek_Saturday,
		"sunday":    WorkingHoursDaysOfWeek_Sunday,
		"thursday":  WorkingHoursDaysOfWeek_Thursday,
		"tuesday":   WorkingHoursDaysOfWeek_Tuesday,
		"wednesday": WorkingHoursDaysOfWeek_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkingHoursDaysOfWeek(input)
	return &out, nil
}
