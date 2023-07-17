package applyupdate

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateStatus string

const (
	UpdateStatusCompleted  UpdateStatus = "Completed"
	UpdateStatusInProgress UpdateStatus = "InProgress"
	UpdateStatusPending    UpdateStatus = "Pending"
	UpdateStatusRetryLater UpdateStatus = "RetryLater"
	UpdateStatusRetryNow   UpdateStatus = "RetryNow"
)

func PossibleValuesForUpdateStatus() []string {
	return []string{
		string(UpdateStatusCompleted),
		string(UpdateStatusInProgress),
		string(UpdateStatusPending),
		string(UpdateStatusRetryLater),
		string(UpdateStatusRetryNow),
	}
}

func (s *UpdateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpdateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpdateStatus(input string) (*UpdateStatus, error) {
	vals := map[string]UpdateStatus{
		"completed":  UpdateStatusCompleted,
		"inprogress": UpdateStatusInProgress,
		"pending":    UpdateStatusPending,
		"retrylater": UpdateStatusRetryLater,
		"retrynow":   UpdateStatusRetryNow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateStatus(input)
	return &out, nil
}
