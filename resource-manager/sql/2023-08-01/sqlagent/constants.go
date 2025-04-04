package sqlagent

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlAgentState string

const (
	SqlAgentStateDisabled SqlAgentState = "Disabled"
	SqlAgentStateEnabled  SqlAgentState = "Enabled"
)

func PossibleValuesForSqlAgentState() []string {
	return []string{
		string(SqlAgentStateDisabled),
		string(SqlAgentStateEnabled),
	}
}

func (s *SqlAgentState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlAgentState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlAgentState(input string) (*SqlAgentState, error) {
	vals := map[string]SqlAgentState{
		"disabled": SqlAgentStateDisabled,
		"enabled":  SqlAgentStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlAgentState(input)
	return &out, nil
}
