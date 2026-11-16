package logfiles

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogFileType string

const (
	LogFileTypePostgreSqlServer LogFileType = "PostgreSqlServer"
)

func PossibleValuesForLogFileType() []string {
	return []string{
		string(LogFileTypePostgreSqlServer),
	}
}

func (s *LogFileType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogFileType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogFileType(input string) (*LogFileType, error) {
	vals := map[string]LogFileType{
		"postgresqlserver": LogFileTypePostgreSqlServer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogFileType(input)
	return &out, nil
}

type NodeRole string

const (
	NodeRoleReadOnly   NodeRole = "ReadOnly"
	NodeRoleReadWrite  NodeRole = "ReadWrite"
	NodeRoleUnassigned NodeRole = "Unassigned"
)

func PossibleValuesForNodeRole() []string {
	return []string{
		string(NodeRoleReadOnly),
		string(NodeRoleReadWrite),
		string(NodeRoleUnassigned),
	}
}

func (s *NodeRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNodeRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNodeRole(input string) (*NodeRole, error) {
	vals := map[string]NodeRole{
		"readonly":   NodeRoleReadOnly,
		"readwrite":  NodeRoleReadWrite,
		"unassigned": NodeRoleUnassigned,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NodeRole(input)
	return &out, nil
}
