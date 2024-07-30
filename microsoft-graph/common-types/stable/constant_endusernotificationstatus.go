package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationStatus string

const (
	EndUserNotificationStatus_Archive EndUserNotificationStatus = "archive"
	EndUserNotificationStatus_Delete  EndUserNotificationStatus = "delete"
	EndUserNotificationStatus_Draft   EndUserNotificationStatus = "draft"
	EndUserNotificationStatus_Ready   EndUserNotificationStatus = "ready"
	EndUserNotificationStatus_Unknown EndUserNotificationStatus = "unknown"
)

func PossibleValuesForEndUserNotificationStatus() []string {
	return []string{
		string(EndUserNotificationStatus_Archive),
		string(EndUserNotificationStatus_Delete),
		string(EndUserNotificationStatus_Draft),
		string(EndUserNotificationStatus_Ready),
		string(EndUserNotificationStatus_Unknown),
	}
}

func (s *EndUserNotificationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndUserNotificationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndUserNotificationStatus(input string) (*EndUserNotificationStatus, error) {
	vals := map[string]EndUserNotificationStatus{
		"archive": EndUserNotificationStatus_Archive,
		"delete":  EndUserNotificationStatus_Delete,
		"draft":   EndUserNotificationStatus_Draft,
		"ready":   EndUserNotificationStatus_Ready,
		"unknown": EndUserNotificationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationStatus(input)
	return &out, nil
}
