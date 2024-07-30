package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomUpdateTimeWindowStartDay string

const (
	CustomUpdateTimeWindowStartDay_Friday    CustomUpdateTimeWindowStartDay = "friday"
	CustomUpdateTimeWindowStartDay_Monday    CustomUpdateTimeWindowStartDay = "monday"
	CustomUpdateTimeWindowStartDay_Saturday  CustomUpdateTimeWindowStartDay = "saturday"
	CustomUpdateTimeWindowStartDay_Sunday    CustomUpdateTimeWindowStartDay = "sunday"
	CustomUpdateTimeWindowStartDay_Thursday  CustomUpdateTimeWindowStartDay = "thursday"
	CustomUpdateTimeWindowStartDay_Tuesday   CustomUpdateTimeWindowStartDay = "tuesday"
	CustomUpdateTimeWindowStartDay_Wednesday CustomUpdateTimeWindowStartDay = "wednesday"
)

func PossibleValuesForCustomUpdateTimeWindowStartDay() []string {
	return []string{
		string(CustomUpdateTimeWindowStartDay_Friday),
		string(CustomUpdateTimeWindowStartDay_Monday),
		string(CustomUpdateTimeWindowStartDay_Saturday),
		string(CustomUpdateTimeWindowStartDay_Sunday),
		string(CustomUpdateTimeWindowStartDay_Thursday),
		string(CustomUpdateTimeWindowStartDay_Tuesday),
		string(CustomUpdateTimeWindowStartDay_Wednesday),
	}
}

func (s *CustomUpdateTimeWindowStartDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomUpdateTimeWindowStartDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomUpdateTimeWindowStartDay(input string) (*CustomUpdateTimeWindowStartDay, error) {
	vals := map[string]CustomUpdateTimeWindowStartDay{
		"friday":    CustomUpdateTimeWindowStartDay_Friday,
		"monday":    CustomUpdateTimeWindowStartDay_Monday,
		"saturday":  CustomUpdateTimeWindowStartDay_Saturday,
		"sunday":    CustomUpdateTimeWindowStartDay_Sunday,
		"thursday":  CustomUpdateTimeWindowStartDay_Thursday,
		"tuesday":   CustomUpdateTimeWindowStartDay_Tuesday,
		"wednesday": CustomUpdateTimeWindowStartDay_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomUpdateTimeWindowStartDay(input)
	return &out, nil
}
