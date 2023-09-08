package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorGroupConnectorGroupType string

const (
	ConnectorGroupConnectorGroupTypeapplicationProxy ConnectorGroupConnectorGroupType = "ApplicationProxy"
)

func PossibleValuesForConnectorGroupConnectorGroupType() []string {
	return []string{
		string(ConnectorGroupConnectorGroupTypeapplicationProxy),
	}
}

func parseConnectorGroupConnectorGroupType(input string) (*ConnectorGroupConnectorGroupType, error) {
	vals := map[string]ConnectorGroupConnectorGroupType{
		"applicationproxy": ConnectorGroupConnectorGroupTypeapplicationProxy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorGroupConnectorGroupType(input)
	return &out, nil
}
