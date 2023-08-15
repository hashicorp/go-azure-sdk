package notificationchannels

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationChannelEventType string

const (
	NotificationChannelEventTypeAutoShutdown NotificationChannelEventType = "AutoShutdown"
	NotificationChannelEventTypeCost         NotificationChannelEventType = "Cost"
)

func PossibleValuesForNotificationChannelEventType() []string {
	return []string{
		string(NotificationChannelEventTypeAutoShutdown),
		string(NotificationChannelEventTypeCost),
	}
}

func (s *NotificationChannelEventType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationChannelEventType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationChannelEventType(input string) (*NotificationChannelEventType, error) {
	vals := map[string]NotificationChannelEventType{
		"autoshutdown": NotificationChannelEventTypeAutoShutdown,
		"cost":         NotificationChannelEventTypeCost,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationChannelEventType(input)
	return &out, nil
}
