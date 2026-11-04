package sqlagent

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlAgentConfigurationPropertiesState string

const (
	SqlAgentConfigurationPropertiesStateDisabled SqlAgentConfigurationPropertiesState = "Disabled"
	SqlAgentConfigurationPropertiesStateEnabled  SqlAgentConfigurationPropertiesState = "Enabled"
)

func PossibleValuesForSqlAgentConfigurationPropertiesState() []string {
	return []string{
		string(SqlAgentConfigurationPropertiesStateDisabled),
		string(SqlAgentConfigurationPropertiesStateEnabled),
	}
}

func (s *SqlAgentConfigurationPropertiesState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlAgentConfigurationPropertiesState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlAgentConfigurationPropertiesState(input string) (*SqlAgentConfigurationPropertiesState, error) {
	vals := map[string]SqlAgentConfigurationPropertiesState{
		"disabled": SqlAgentConfigurationPropertiesStateDisabled,
		"enabled":  SqlAgentConfigurationPropertiesStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlAgentConfigurationPropertiesState(input)
	return &out, nil
}
