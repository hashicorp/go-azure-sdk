package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingSourceType string

const (
	SettingSourceType_DeviceConfiguration SettingSourceType = "deviceConfiguration"
	SettingSourceType_DeviceIntent        SettingSourceType = "deviceIntent"
)

func PossibleValuesForSettingSourceType() []string {
	return []string{
		string(SettingSourceType_DeviceConfiguration),
		string(SettingSourceType_DeviceIntent),
	}
}

func (s *SettingSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSettingSourceType(input string) (*SettingSourceType, error) {
	vals := map[string]SettingSourceType{
		"deviceconfiguration": SettingSourceType_DeviceConfiguration,
		"deviceintent":        SettingSourceType_DeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingSourceType(input)
	return &out, nil
}
