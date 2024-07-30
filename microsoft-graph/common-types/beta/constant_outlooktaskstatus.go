package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskStatus string

const (
	OutlookTaskStatus_Completed       OutlookTaskStatus = "completed"
	OutlookTaskStatus_Deferred        OutlookTaskStatus = "deferred"
	OutlookTaskStatus_InProgress      OutlookTaskStatus = "inProgress"
	OutlookTaskStatus_NotStarted      OutlookTaskStatus = "notStarted"
	OutlookTaskStatus_WaitingOnOthers OutlookTaskStatus = "waitingOnOthers"
)

func PossibleValuesForOutlookTaskStatus() []string {
	return []string{
		string(OutlookTaskStatus_Completed),
		string(OutlookTaskStatus_Deferred),
		string(OutlookTaskStatus_InProgress),
		string(OutlookTaskStatus_NotStarted),
		string(OutlookTaskStatus_WaitingOnOthers),
	}
}

func (s *OutlookTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOutlookTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOutlookTaskStatus(input string) (*OutlookTaskStatus, error) {
	vals := map[string]OutlookTaskStatus{
		"completed":       OutlookTaskStatus_Completed,
		"deferred":        OutlookTaskStatus_Deferred,
		"inprogress":      OutlookTaskStatus_InProgress,
		"notstarted":      OutlookTaskStatus_NotStarted,
		"waitingonothers": OutlookTaskStatus_WaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskStatus(input)
	return &out, nil
}
