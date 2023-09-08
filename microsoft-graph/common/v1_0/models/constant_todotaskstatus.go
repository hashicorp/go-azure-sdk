package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskStatus string

const (
	TodoTaskStatuscompleted       TodoTaskStatus = "Completed"
	TodoTaskStatusdeferred        TodoTaskStatus = "Deferred"
	TodoTaskStatusinProgress      TodoTaskStatus = "InProgress"
	TodoTaskStatusnotStarted      TodoTaskStatus = "NotStarted"
	TodoTaskStatuswaitingOnOthers TodoTaskStatus = "WaitingOnOthers"
)

func PossibleValuesForTodoTaskStatus() []string {
	return []string{
		string(TodoTaskStatuscompleted),
		string(TodoTaskStatusdeferred),
		string(TodoTaskStatusinProgress),
		string(TodoTaskStatusnotStarted),
		string(TodoTaskStatuswaitingOnOthers),
	}
}

func parseTodoTaskStatus(input string) (*TodoTaskStatus, error) {
	vals := map[string]TodoTaskStatus{
		"completed":       TodoTaskStatuscompleted,
		"deferred":        TodoTaskStatusdeferred,
		"inprogress":      TodoTaskStatusinProgress,
		"notstarted":      TodoTaskStatusnotStarted,
		"waitingonothers": TodoTaskStatuswaitingOnOthers,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskStatus(input)
	return &out, nil
}
