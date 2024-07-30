package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertHistoryStateStatus string

const (
	AlertHistoryStateStatus_Dismissed  AlertHistoryStateStatus = "dismissed"
	AlertHistoryStateStatus_InProgress AlertHistoryStateStatus = "inProgress"
	AlertHistoryStateStatus_NewAlert   AlertHistoryStateStatus = "newAlert"
	AlertHistoryStateStatus_Resolved   AlertHistoryStateStatus = "resolved"
	AlertHistoryStateStatus_Unknown    AlertHistoryStateStatus = "unknown"
)

func PossibleValuesForAlertHistoryStateStatus() []string {
	return []string{
		string(AlertHistoryStateStatus_Dismissed),
		string(AlertHistoryStateStatus_InProgress),
		string(AlertHistoryStateStatus_NewAlert),
		string(AlertHistoryStateStatus_Resolved),
		string(AlertHistoryStateStatus_Unknown),
	}
}

func (s *AlertHistoryStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAlertHistoryStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAlertHistoryStateStatus(input string) (*AlertHistoryStateStatus, error) {
	vals := map[string]AlertHistoryStateStatus{
		"dismissed":  AlertHistoryStateStatus_Dismissed,
		"inprogress": AlertHistoryStateStatus_InProgress,
		"newalert":   AlertHistoryStateStatus_NewAlert,
		"resolved":   AlertHistoryStateStatus_Resolved,
		"unknown":    AlertHistoryStateStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AlertHistoryStateStatus(input)
	return &out, nil
}
