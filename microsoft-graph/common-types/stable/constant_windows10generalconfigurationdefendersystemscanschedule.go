package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderSystemScanSchedule string

const (
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Everyday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "everyday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Friday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "friday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Monday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "monday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Saturday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "saturday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Sunday      Windows10GeneralConfigurationDefenderSystemScanSchedule = "sunday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Thursday    Windows10GeneralConfigurationDefenderSystemScanSchedule = "thursday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Tuesday     Windows10GeneralConfigurationDefenderSystemScanSchedule = "tuesday"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_UserDefined Windows10GeneralConfigurationDefenderSystemScanSchedule = "userDefined"
	Windows10GeneralConfigurationDefenderSystemScanSchedule_Wednesday   Windows10GeneralConfigurationDefenderSystemScanSchedule = "wednesday"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderSystemScanSchedule() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Everyday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Friday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Monday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Saturday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Sunday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Thursday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Tuesday),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_UserDefined),
		string(Windows10GeneralConfigurationDefenderSystemScanSchedule_Wednesday),
	}
}

func (s *Windows10GeneralConfigurationDefenderSystemScanSchedule) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderSystemScanSchedule(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderSystemScanSchedule(input string) (*Windows10GeneralConfigurationDefenderSystemScanSchedule, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderSystemScanSchedule{
		"everyday":    Windows10GeneralConfigurationDefenderSystemScanSchedule_Everyday,
		"friday":      Windows10GeneralConfigurationDefenderSystemScanSchedule_Friday,
		"monday":      Windows10GeneralConfigurationDefenderSystemScanSchedule_Monday,
		"saturday":    Windows10GeneralConfigurationDefenderSystemScanSchedule_Saturday,
		"sunday":      Windows10GeneralConfigurationDefenderSystemScanSchedule_Sunday,
		"thursday":    Windows10GeneralConfigurationDefenderSystemScanSchedule_Thursday,
		"tuesday":     Windows10GeneralConfigurationDefenderSystemScanSchedule_Tuesday,
		"userdefined": Windows10GeneralConfigurationDefenderSystemScanSchedule_UserDefined,
		"wednesday":   Windows10GeneralConfigurationDefenderSystemScanSchedule_Wednesday,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderSystemScanSchedule(input)
	return &out, nil
}
