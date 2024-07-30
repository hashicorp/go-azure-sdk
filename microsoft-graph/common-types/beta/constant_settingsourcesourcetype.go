package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingSourceSourceType string

const (
	SettingSourceSourceType_DeviceConfiguration SettingSourceSourceType = "deviceConfiguration"
	SettingSourceSourceType_DeviceIntent        SettingSourceSourceType = "deviceIntent"
)

func PossibleValuesForSettingSourceSourceType() []string {
	return []string{
		string(SettingSourceSourceType_DeviceConfiguration),
		string(SettingSourceSourceType_DeviceIntent),
	}
}

func (s *SettingSourceSourceType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSettingSourceSourceType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSettingSourceSourceType(input string) (*SettingSourceSourceType, error) {
	vals := map[string]SettingSourceSourceType{
		"deviceconfiguration": SettingSourceSourceType_DeviceConfiguration,
		"deviceintent":        SettingSourceSourceType_DeviceIntent,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SettingSourceSourceType(input)
	return &out, nil
}
