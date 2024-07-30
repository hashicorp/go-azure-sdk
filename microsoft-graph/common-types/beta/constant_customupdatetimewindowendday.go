package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomUpdateTimeWindowEndDay string

const (
	CustomUpdateTimeWindowEndDay_Friday    CustomUpdateTimeWindowEndDay = "friday"
	CustomUpdateTimeWindowEndDay_Monday    CustomUpdateTimeWindowEndDay = "monday"
	CustomUpdateTimeWindowEndDay_Saturday  CustomUpdateTimeWindowEndDay = "saturday"
	CustomUpdateTimeWindowEndDay_Sunday    CustomUpdateTimeWindowEndDay = "sunday"
	CustomUpdateTimeWindowEndDay_Thursday  CustomUpdateTimeWindowEndDay = "thursday"
	CustomUpdateTimeWindowEndDay_Tuesday   CustomUpdateTimeWindowEndDay = "tuesday"
	CustomUpdateTimeWindowEndDay_Wednesday CustomUpdateTimeWindowEndDay = "wednesday"
)

func PossibleValuesForCustomUpdateTimeWindowEndDay() []string {
	return []string{
		string(CustomUpdateTimeWindowEndDay_Friday),
		string(CustomUpdateTimeWindowEndDay_Monday),
		string(CustomUpdateTimeWindowEndDay_Saturday),
		string(CustomUpdateTimeWindowEndDay_Sunday),
		string(CustomUpdateTimeWindowEndDay_Thursday),
		string(CustomUpdateTimeWindowEndDay_Tuesday),
		string(CustomUpdateTimeWindowEndDay_Wednesday),
	}
}

func (s *CustomUpdateTimeWindowEndDay) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomUpdateTimeWindowEndDay(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomUpdateTimeWindowEndDay(input string) (*CustomUpdateTimeWindowEndDay, error) {
	vals := map[string]CustomUpdateTimeWindowEndDay{
		"friday":    CustomUpdateTimeWindowEndDay_Friday,
		"monday":    CustomUpdateTimeWindowEndDay_Monday,
		"saturday":  CustomUpdateTimeWindowEndDay_Saturday,
		"sunday":    CustomUpdateTimeWindowEndDay_Sunday,
		"thursday":  CustomUpdateTimeWindowEndDay_Thursday,
		"tuesday":   CustomUpdateTimeWindowEndDay_Tuesday,
		"wednesday": CustomUpdateTimeWindowEndDay_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomUpdateTimeWindowEndDay(input)
	return &out, nil
}
