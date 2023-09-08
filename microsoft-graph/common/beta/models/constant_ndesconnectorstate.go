package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NdesConnectorState string

const (
	NdesConnectorStateactive   NdesConnectorState = "Active"
	NdesConnectorStateinactive NdesConnectorState = "Inactive"
	NdesConnectorStatenone     NdesConnectorState = "None"
)

func PossibleValuesForNdesConnectorState() []string {
	return []string{
		string(NdesConnectorStateactive),
		string(NdesConnectorStateinactive),
		string(NdesConnectorStatenone),
	}
}

func parseNdesConnectorState(input string) (*NdesConnectorState, error) {
	vals := map[string]NdesConnectorState{
		"active":   NdesConnectorStateactive,
		"inactive": NdesConnectorStateinactive,
		"none":     NdesConnectorStatenone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NdesConnectorState(input)
	return &out, nil
}
