package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingReminderNotificationDeliveryFrequency string

const (
	TrainingReminderNotificationDeliveryFrequency_BiWeekly TrainingReminderNotificationDeliveryFrequency = "biWeekly"
	TrainingReminderNotificationDeliveryFrequency_Unknown  TrainingReminderNotificationDeliveryFrequency = "unknown"
	TrainingReminderNotificationDeliveryFrequency_Weekly   TrainingReminderNotificationDeliveryFrequency = "weekly"
)

func PossibleValuesForTrainingReminderNotificationDeliveryFrequency() []string {
	return []string{
		string(TrainingReminderNotificationDeliveryFrequency_BiWeekly),
		string(TrainingReminderNotificationDeliveryFrequency_Unknown),
		string(TrainingReminderNotificationDeliveryFrequency_Weekly),
	}
}

func (s *TrainingReminderNotificationDeliveryFrequency) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingReminderNotificationDeliveryFrequency(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingReminderNotificationDeliveryFrequency(input string) (*TrainingReminderNotificationDeliveryFrequency, error) {
	vals := map[string]TrainingReminderNotificationDeliveryFrequency{
		"biweekly": TrainingReminderNotificationDeliveryFrequency_BiWeekly,
		"unknown":  TrainingReminderNotificationDeliveryFrequency_Unknown,
		"weekly":   TrainingReminderNotificationDeliveryFrequency_Weekly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingReminderNotificationDeliveryFrequency(input)
	return &out, nil
}
