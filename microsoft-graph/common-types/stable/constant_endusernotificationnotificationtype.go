package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationNotificationType string

const (
	EndUserNotificationNotificationType_NoTraining            EndUserNotificationNotificationType = "noTraining"
	EndUserNotificationNotificationType_PositiveReinforcement EndUserNotificationNotificationType = "positiveReinforcement"
	EndUserNotificationNotificationType_TrainingAssignment    EndUserNotificationNotificationType = "trainingAssignment"
	EndUserNotificationNotificationType_TrainingReminder      EndUserNotificationNotificationType = "trainingReminder"
	EndUserNotificationNotificationType_Unknown               EndUserNotificationNotificationType = "unknown"
)

func PossibleValuesForEndUserNotificationNotificationType() []string {
	return []string{
		string(EndUserNotificationNotificationType_NoTraining),
		string(EndUserNotificationNotificationType_PositiveReinforcement),
		string(EndUserNotificationNotificationType_TrainingAssignment),
		string(EndUserNotificationNotificationType_TrainingReminder),
		string(EndUserNotificationNotificationType_Unknown),
	}
}

func (s *EndUserNotificationNotificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationNotificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationNotificationType(input string) (*EndUserNotificationNotificationType, error) {
	vals := map[string]EndUserNotificationNotificationType{
		"notraining":            EndUserNotificationNotificationType_NoTraining,
		"positivereinforcement": EndUserNotificationNotificationType_PositiveReinforcement,
		"trainingassignment":    EndUserNotificationNotificationType_TrainingAssignment,
		"trainingreminder":      EndUserNotificationNotificationType_TrainingReminder,
		"unknown":               EndUserNotificationNotificationType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationNotificationType(input)
	return &out, nil
}
