package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingNotificationSettingSettingType string

const (
	NoTrainingNotificationSettingSettingType_NoNotification   NoTrainingNotificationSettingSettingType = "noNotification"
	NoTrainingNotificationSettingSettingType_NoTraining       NoTrainingNotificationSettingSettingType = "noTraining"
	NoTrainingNotificationSettingSettingType_TrainingSelected NoTrainingNotificationSettingSettingType = "trainingSelected"
	NoTrainingNotificationSettingSettingType_Unknown          NoTrainingNotificationSettingSettingType = "unknown"
)

func PossibleValuesForNoTrainingNotificationSettingSettingType() []string {
	return []string{
		string(NoTrainingNotificationSettingSettingType_NoNotification),
		string(NoTrainingNotificationSettingSettingType_NoTraining),
		string(NoTrainingNotificationSettingSettingType_TrainingSelected),
		string(NoTrainingNotificationSettingSettingType_Unknown),
	}
}

func (s *NoTrainingNotificationSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNoTrainingNotificationSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNoTrainingNotificationSettingSettingType(input string) (*NoTrainingNotificationSettingSettingType, error) {
	vals := map[string]NoTrainingNotificationSettingSettingType{
		"nonotification":   NoTrainingNotificationSettingSettingType_NoNotification,
		"notraining":       NoTrainingNotificationSettingSettingType_NoTraining,
		"trainingselected": NoTrainingNotificationSettingSettingType_TrainingSelected,
		"unknown":          NoTrainingNotificationSettingSettingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingNotificationSettingSettingType(input)
	return &out, nil
}
