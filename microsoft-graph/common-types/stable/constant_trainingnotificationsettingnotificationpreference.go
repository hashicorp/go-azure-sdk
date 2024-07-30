package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationSettingNotificationPreference string

const (
	TrainingNotificationSettingNotificationPreference_Custom    TrainingNotificationSettingNotificationPreference = "custom"
	TrainingNotificationSettingNotificationPreference_Microsoft TrainingNotificationSettingNotificationPreference = "microsoft"
	TrainingNotificationSettingNotificationPreference_Unknown   TrainingNotificationSettingNotificationPreference = "unknown"
)

func PossibleValuesForTrainingNotificationSettingNotificationPreference() []string {
	return []string{
		string(TrainingNotificationSettingNotificationPreference_Custom),
		string(TrainingNotificationSettingNotificationPreference_Microsoft),
		string(TrainingNotificationSettingNotificationPreference_Unknown),
	}
}

func (s *TrainingNotificationSettingNotificationPreference) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingNotificationSettingNotificationPreference(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingNotificationSettingNotificationPreference(input string) (*TrainingNotificationSettingNotificationPreference, error) {
	vals := map[string]TrainingNotificationSettingNotificationPreference{
		"custom":    TrainingNotificationSettingNotificationPreference_Custom,
		"microsoft": TrainingNotificationSettingNotificationPreference_Microsoft,
		"unknown":   TrainingNotificationSettingNotificationPreference_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingNotificationSettingNotificationPreference(input)
	return &out, nil
}
