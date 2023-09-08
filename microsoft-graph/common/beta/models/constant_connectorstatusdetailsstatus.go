package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetailsStatus string

const (
	ConnectorStatusDetailsStatushealthy   ConnectorStatusDetailsStatus = "Healthy"
	ConnectorStatusDetailsStatusunhealthy ConnectorStatusDetailsStatus = "Unhealthy"
	ConnectorStatusDetailsStatusunknown   ConnectorStatusDetailsStatus = "Unknown"
	ConnectorStatusDetailsStatuswarning   ConnectorStatusDetailsStatus = "Warning"
)

func PossibleValuesForConnectorStatusDetailsStatus() []string {
	return []string{
		string(ConnectorStatusDetailsStatushealthy),
		string(ConnectorStatusDetailsStatusunhealthy),
		string(ConnectorStatusDetailsStatusunknown),
		string(ConnectorStatusDetailsStatuswarning),
	}
}

func parseConnectorStatusDetailsStatus(input string) (*ConnectorStatusDetailsStatus, error) {
	vals := map[string]ConnectorStatusDetailsStatus{
		"healthy":   ConnectorStatusDetailsStatushealthy,
		"unhealthy": ConnectorStatusDetailsStatusunhealthy,
		"unknown":   ConnectorStatusDetailsStatusunknown,
		"warning":   ConnectorStatusDetailsStatuswarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatusDetailsStatus(input)
	return &out, nil
}
