package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TodoTaskListWellknownListName string

const (
	TodoTaskListWellknownListName_DefaultList   TodoTaskListWellknownListName = "defaultList"
	TodoTaskListWellknownListName_FlaggedEmails TodoTaskListWellknownListName = "flaggedEmails"
	TodoTaskListWellknownListName_None          TodoTaskListWellknownListName = "none"
)

func PossibleValuesForTodoTaskListWellknownListName() []string {
	return []string{
		string(TodoTaskListWellknownListName_DefaultList),
		string(TodoTaskListWellknownListName_FlaggedEmails),
		string(TodoTaskListWellknownListName_None),
	}
}

func (s *TodoTaskListWellknownListName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTodoTaskListWellknownListName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTodoTaskListWellknownListName(input string) (*TodoTaskListWellknownListName, error) {
	vals := map[string]TodoTaskListWellknownListName{
		"defaultlist":   TodoTaskListWellknownListName_DefaultList,
		"flaggedemails": TodoTaskListWellknownListName_FlaggedEmails,
		"none":          TodoTaskListWellknownListName_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TodoTaskListWellknownListName(input)
	return &out, nil
}
