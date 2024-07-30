package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationUpdateWeeks string

const (
	WindowsUpdateForBusinessConfigurationUpdateWeeks_EveryWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "everyWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeks_FirstWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "firstWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeks_FourthWeek  WindowsUpdateForBusinessConfigurationUpdateWeeks = "fourthWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeks_SecondWeek  WindowsUpdateForBusinessConfigurationUpdateWeeks = "secondWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeks_ThirdWeek   WindowsUpdateForBusinessConfigurationUpdateWeeks = "thirdWeek"
	WindowsUpdateForBusinessConfigurationUpdateWeeks_UserDefined WindowsUpdateForBusinessConfigurationUpdateWeeks = "userDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationUpdateWeeks() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_EveryWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_FirstWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_FourthWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_SecondWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_ThirdWeek),
		string(WindowsUpdateForBusinessConfigurationUpdateWeeks_UserDefined),
	}
}

func (s *WindowsUpdateForBusinessConfigurationUpdateWeeks) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationUpdateWeeks(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationUpdateWeeks(input string) (*WindowsUpdateForBusinessConfigurationUpdateWeeks, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationUpdateWeeks{
		"everyweek":   WindowsUpdateForBusinessConfigurationUpdateWeeks_EveryWeek,
		"firstweek":   WindowsUpdateForBusinessConfigurationUpdateWeeks_FirstWeek,
		"fourthweek":  WindowsUpdateForBusinessConfigurationUpdateWeeks_FourthWeek,
		"secondweek":  WindowsUpdateForBusinessConfigurationUpdateWeeks_SecondWeek,
		"thirdweek":   WindowsUpdateForBusinessConfigurationUpdateWeeks_ThirdWeek,
		"userdefined": WindowsUpdateForBusinessConfigurationUpdateWeeks_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationUpdateWeeks(input)
	return &out, nil
}
