package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutOfBoxExperienceSettingsDeviceUsageType string

const (
	OutOfBoxExperienceSettingsDeviceUsageType_Shared     OutOfBoxExperienceSettingsDeviceUsageType = "shared"
	OutOfBoxExperienceSettingsDeviceUsageType_SingleUser OutOfBoxExperienceSettingsDeviceUsageType = "singleUser"
)

func PossibleValuesForOutOfBoxExperienceSettingsDeviceUsageType() []string {
	return []string{
		string(OutOfBoxExperienceSettingsDeviceUsageType_Shared),
		string(OutOfBoxExperienceSettingsDeviceUsageType_SingleUser),
	}
}

func (s *OutOfBoxExperienceSettingsDeviceUsageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutOfBoxExperienceSettingsDeviceUsageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutOfBoxExperienceSettingsDeviceUsageType(input string) (*OutOfBoxExperienceSettingsDeviceUsageType, error) {
	vals := map[string]OutOfBoxExperienceSettingsDeviceUsageType{
		"shared":     OutOfBoxExperienceSettingsDeviceUsageType_Shared,
		"singleuser": OutOfBoxExperienceSettingsDeviceUsageType_SingleUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutOfBoxExperienceSettingsDeviceUsageType(input)
	return &out, nil
}
