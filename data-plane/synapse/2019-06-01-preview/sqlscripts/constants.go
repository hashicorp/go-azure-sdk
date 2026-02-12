package sqlscripts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlConnectionType string

const (
	SqlConnectionTypeSqlOnDemand SqlConnectionType = "SqlOnDemand"
	SqlConnectionTypeSqlPool     SqlConnectionType = "SqlPool"
)

func PossibleValuesForSqlConnectionType() []string {
	return []string{
		string(SqlConnectionTypeSqlOnDemand),
		string(SqlConnectionTypeSqlPool),
	}
}

func (s *SqlConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlConnectionType(input string) (*SqlConnectionType, error) {
	vals := map[string]SqlConnectionType{
		"sqlondemand": SqlConnectionTypeSqlOnDemand,
		"sqlpool":     SqlConnectionTypeSqlPool,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlConnectionType(input)
	return &out, nil
}

type SqlScriptType string

const (
	SqlScriptTypeSqlQuery SqlScriptType = "SqlQuery"
)

func PossibleValuesForSqlScriptType() []string {
	return []string{
		string(SqlScriptTypeSqlQuery),
	}
}

func (s *SqlScriptType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSqlScriptType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSqlScriptType(input string) (*SqlScriptType, error) {
	vals := map[string]SqlScriptType{
		"sqlquery": SqlScriptTypeSqlQuery,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SqlScriptType(input)
	return &out, nil
}
