package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSettingSettingType string

const (
	EndUserNotificationSettingSettingType_NoNotification   EndUserNotificationSettingSettingType = "noNotification"
	EndUserNotificationSettingSettingType_NoTraining       EndUserNotificationSettingSettingType = "noTraining"
	EndUserNotificationSettingSettingType_TrainingSelected EndUserNotificationSettingSettingType = "trainingSelected"
	EndUserNotificationSettingSettingType_Unknown          EndUserNotificationSettingSettingType = "unknown"
)

func PossibleValuesForEndUserNotificationSettingSettingType() []string {
	return []string{
		string(EndUserNotificationSettingSettingType_NoNotification),
		string(EndUserNotificationSettingSettingType_NoTraining),
		string(EndUserNotificationSettingSettingType_TrainingSelected),
		string(EndUserNotificationSettingSettingType_Unknown),
	}
}

func (s *EndUserNotificationSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationSettingSettingType(input string) (*EndUserNotificationSettingSettingType, error) {
	vals := map[string]EndUserNotificationSettingSettingType{
		"nonotification":   EndUserNotificationSettingSettingType_NoNotification,
		"notraining":       EndUserNotificationSettingSettingType_NoTraining,
		"trainingselected": EndUserNotificationSettingSettingType_TrainingSelected,
		"unknown":          EndUserNotificationSettingSettingType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSettingSettingType(input)
	return &out, nil
}
