package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatus string

const (
	ConnectorStatusactive   ConnectorStatus = "Active"
	ConnectorStatusinactive ConnectorStatus = "Inactive"
)

func PossibleValuesForConnectorStatus() []string {
	return []string{
		string(ConnectorStatusactive),
		string(ConnectorStatusinactive),
	}
}

func parseConnectorStatus(input string) (*ConnectorStatus, error) {
	vals := map[string]ConnectorStatus{
		"active":   ConnectorStatusactive,
		"inactive": ConnectorStatusinactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatus(input)
	return &out, nil
}
