package allowedconnections

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionType string

const (
	ConnectionTypeExternal ConnectionType = "External"
	ConnectionTypeInternal ConnectionType = "Internal"
)

func PossibleValuesForConnectionType() []string {
	return []string{
		string(ConnectionTypeExternal),
		string(ConnectionTypeInternal),
	}
}

func parseConnectionType(input string) (*ConnectionType, error) {
	vals := map[string]ConnectionType{
		"external": ConnectionTypeExternal,
		"internal": ConnectionTypeInternal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionType(input)
	return &out, nil
}
