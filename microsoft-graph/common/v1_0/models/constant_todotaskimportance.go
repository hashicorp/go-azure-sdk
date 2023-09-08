package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskImportance string

const (
	TodoTaskImportancehigh   TodoTaskImportance = "High"
	TodoTaskImportancelow    TodoTaskImportance = "Low"
	TodoTaskImportancenormal TodoTaskImportance = "Normal"
)

func PossibleValuesForTodoTaskImportance() []string {
	return []string{
		string(TodoTaskImportancehigh),
		string(TodoTaskImportancelow),
		string(TodoTaskImportancenormal),
	}
}

func parseTodoTaskImportance(input string) (*TodoTaskImportance, error) {
	vals := map[string]TodoTaskImportance{
		"high":   TodoTaskImportancehigh,
		"low":    TodoTaskImportancelow,
		"normal": TodoTaskImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskImportance(input)
	return &out, nil
}
