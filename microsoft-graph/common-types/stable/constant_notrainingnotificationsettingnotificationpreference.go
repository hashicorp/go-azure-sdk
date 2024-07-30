package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingNotificationSettingNotificationPreference string

const (
	NoTrainingNotificationSettingNotificationPreference_Custom    NoTrainingNotificationSettingNotificationPreference = "custom"
	NoTrainingNotificationSettingNotificationPreference_Microsoft NoTrainingNotificationSettingNotificationPreference = "microsoft"
	NoTrainingNotificationSettingNotificationPreference_Unknown   NoTrainingNotificationSettingNotificationPreference = "unknown"
)

func PossibleValuesForNoTrainingNotificationSettingNotificationPreference() []string {
	return []string{
		string(NoTrainingNotificationSettingNotificationPreference_Custom),
		string(NoTrainingNotificationSettingNotificationPreference_Microsoft),
		string(NoTrainingNotificationSettingNotificationPreference_Unknown),
	}
}

func (s *NoTrainingNotificationSettingNotificationPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNoTrainingNotificationSettingNotificationPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNoTrainingNotificationSettingNotificationPreference(input string) (*NoTrainingNotificationSettingNotificationPreference, error) {
	vals := map[string]NoTrainingNotificationSettingNotificationPreference{
		"custom":    NoTrainingNotificationSettingNotificationPreference_Custom,
		"microsoft": NoTrainingNotificationSettingNotificationPreference_Microsoft,
		"unknown":   NoTrainingNotificationSettingNotificationPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingNotificationSettingNotificationPreference(input)
	return &out, nil
}
