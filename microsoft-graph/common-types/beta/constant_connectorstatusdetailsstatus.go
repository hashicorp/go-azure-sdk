package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectorStatusDetailsStatus string

const (
	ConnectorStatusDetailsStatus_Healthy   ConnectorStatusDetailsStatus = "healthy"
	ConnectorStatusDetailsStatus_Unhealthy ConnectorStatusDetailsStatus = "unhealthy"
	ConnectorStatusDetailsStatus_Unknown   ConnectorStatusDetailsStatus = "unknown"
	ConnectorStatusDetailsStatus_Warning   ConnectorStatusDetailsStatus = "warning"
)

func PossibleValuesForConnectorStatusDetailsStatus() []string {
	return []string{
		string(ConnectorStatusDetailsStatus_Healthy),
		string(ConnectorStatusDetailsStatus_Unhealthy),
		string(ConnectorStatusDetailsStatus_Unknown),
		string(ConnectorStatusDetailsStatus_Warning),
	}
}

func (s *ConnectorStatusDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectorStatusDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectorStatusDetailsStatus(input string) (*ConnectorStatusDetailsStatus, error) {
	vals := map[string]ConnectorStatusDetailsStatus{
		"healthy":   ConnectorStatusDetailsStatus_Healthy,
		"unhealthy": ConnectorStatusDetailsStatus_Unhealthy,
		"unknown":   ConnectorStatusDetailsStatus_Unknown,
		"warning":   ConnectorStatusDetailsStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectorStatusDetailsStatus(input)
	return &out, nil
}
