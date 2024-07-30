package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationPriority string

const (
	NotificationPriority_High NotificationPriority = "High"
	NotificationPriority_Low  NotificationPriority = "Low"
	NotificationPriority_None NotificationPriority = "None"
)

func PossibleValuesForNotificationPriority() []string {
	return []string{
		string(NotificationPriority_High),
		string(NotificationPriority_Low),
		string(NotificationPriority_None),
	}
}

func (s *NotificationPriority) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNotificationPriority(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNotificationPriority(input string) (*NotificationPriority, error) {
	vals := map[string]NotificationPriority{
		"high": NotificationPriority_High,
		"low":  NotificationPriority_Low,
		"none": NotificationPriority_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NotificationPriority(input)
	return &out, nil
}
