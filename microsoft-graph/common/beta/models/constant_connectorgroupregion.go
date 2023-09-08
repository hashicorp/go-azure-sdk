package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorGroupRegion string

const (
	ConnectorGroupRegionasia ConnectorGroupRegion = "Asia"
	ConnectorGroupRegionaus  ConnectorGroupRegion = "Aus"
	ConnectorGroupRegioneur  ConnectorGroupRegion = "Eur"
	ConnectorGroupRegionind  ConnectorGroupRegion = "Ind"
	ConnectorGroupRegionnam  ConnectorGroupRegion = "Nam"
)

func PossibleValuesForConnectorGroupRegion() []string {
	return []string{
		string(ConnectorGroupRegionasia),
		string(ConnectorGroupRegionaus),
		string(ConnectorGroupRegioneur),
		string(ConnectorGroupRegionind),
		string(ConnectorGroupRegionnam),
	}
}

func parseConnectorGroupRegion(input string) (*ConnectorGroupRegion, error) {
	vals := map[string]ConnectorGroupRegion{
		"asia": ConnectorGroupRegionasia,
		"aus":  ConnectorGroupRegionaus,
		"eur":  ConnectorGroupRegioneur,
		"ind":  ConnectorGroupRegionind,
		"nam":  ConnectorGroupRegionnam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorGroupRegion(input)
	return &out, nil
}
