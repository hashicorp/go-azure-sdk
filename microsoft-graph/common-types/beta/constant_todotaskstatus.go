package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskStatus string

const (
	TodoTaskStatus_Completed       TodoTaskStatus = "completed"
	TodoTaskStatus_Deferred        TodoTaskStatus = "deferred"
	TodoTaskStatus_InProgress      TodoTaskStatus = "inProgress"
	TodoTaskStatus_NotStarted      TodoTaskStatus = "notStarted"
	TodoTaskStatus_WaitingOnOthers TodoTaskStatus = "waitingOnOthers"
)

func PossibleValuesForTodoTaskStatus() []string {
	return []string{
		string(TodoTaskStatus_Completed),
		string(TodoTaskStatus_Deferred),
		string(TodoTaskStatus_InProgress),
		string(TodoTaskStatus_NotStarted),
		string(TodoTaskStatus_WaitingOnOthers),
	}
}

func (s *TodoTaskStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskStatus(input string) (*TodoTaskStatus, error) {
	vals := map[string]TodoTaskStatus{
		"completed":       TodoTaskStatus_Completed,
		"deferred":        TodoTaskStatus_Deferred,
		"inprogress":      TodoTaskStatus_InProgress,
		"notstarted":      TodoTaskStatus_NotStarted,
		"waitingonothers": TodoTaskStatus_WaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskStatus(input)
	return &out, nil
}
