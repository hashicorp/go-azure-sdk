package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSettingNotificationPreference string

const (
	EndUserNotificationSettingNotificationPreference_Custom    EndUserNotificationSettingNotificationPreference = "custom"
	EndUserNotificationSettingNotificationPreference_Microsoft EndUserNotificationSettingNotificationPreference = "microsoft"
	EndUserNotificationSettingNotificationPreference_Unknown   EndUserNotificationSettingNotificationPreference = "unknown"
)

func PossibleValuesForEndUserNotificationSettingNotificationPreference() []string {
	return []string{
		string(EndUserNotificationSettingNotificationPreference_Custom),
		string(EndUserNotificationSettingNotificationPreference_Microsoft),
		string(EndUserNotificationSettingNotificationPreference_Unknown),
	}
}

func (s *EndUserNotificationSettingNotificationPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationSettingNotificationPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationSettingNotificationPreference(input string) (*EndUserNotificationSettingNotificationPreference, error) {
	vals := map[string]EndUserNotificationSettingNotificationPreference{
		"custom":    EndUserNotificationSettingNotificationPreference_Custom,
		"microsoft": EndUserNotificationSettingNotificationPreference_Microsoft,
		"unknown":   EndUserNotificationSettingNotificationPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSettingNotificationPreference(input)
	return &out, nil
}
