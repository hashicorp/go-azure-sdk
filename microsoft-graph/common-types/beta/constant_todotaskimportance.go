package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskImportance string

const (
	TodoTaskImportance_High   TodoTaskImportance = "high"
	TodoTaskImportance_Low    TodoTaskImportance = "low"
	TodoTaskImportance_Normal TodoTaskImportance = "normal"
)

func PossibleValuesForTodoTaskImportance() []string {
	return []string{
		string(TodoTaskImportance_High),
		string(TodoTaskImportance_Low),
		string(TodoTaskImportance_Normal),
	}
}

func (s *TodoTaskImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskImportance(input string) (*TodoTaskImportance, error) {
	vals := map[string]TodoTaskImportance{
		"high":   TodoTaskImportance_High,
		"low":    TodoTaskImportance_Low,
		"normal": TodoTaskImportance_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskImportance(input)
	return &out, nil
}
