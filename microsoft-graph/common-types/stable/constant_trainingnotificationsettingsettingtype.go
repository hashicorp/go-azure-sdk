package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationSettingSettingType string

const (
	TrainingNotificationSettingSettingType_NoNotification   TrainingNotificationSettingSettingType = "noNotification"
	TrainingNotificationSettingSettingType_NoTraining       TrainingNotificationSettingSettingType = "noTraining"
	TrainingNotificationSettingSettingType_TrainingSelected TrainingNotificationSettingSettingType = "trainingSelected"
	TrainingNotificationSettingSettingType_Unknown          TrainingNotificationSettingSettingType = "unknown"
)

func PossibleValuesForTrainingNotificationSettingSettingType() []string {
	return []string{
		string(TrainingNotificationSettingSettingType_NoNotification),
		string(TrainingNotificationSettingSettingType_NoTraining),
		string(TrainingNotificationSettingSettingType_TrainingSelected),
		string(TrainingNotificationSettingSettingType_Unknown),
	}
}

func (s *TrainingNotificationSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingNotificationSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingNotificationSettingSettingType(input string) (*TrainingNotificationSettingSettingType, error) {
	vals := map[string]TrainingNotificationSettingSettingType{
		"nonotification":   TrainingNotificationSettingSettingType_NoNotification,
		"notraining":       TrainingNotificationSettingSettingType_NoTraining,
		"trainingselected": TrainingNotificationSettingSettingType_TrainingSelected,
		"unknown":          TrainingNotificationSettingSettingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingNotificationSettingSettingType(input)
	return &out, nil
}
