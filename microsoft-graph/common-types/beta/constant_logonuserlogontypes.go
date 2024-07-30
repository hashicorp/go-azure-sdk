package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUserLogonTypes string

const (
	LogonUserLogonTypes_Batch             LogonUserLogonTypes = "batch"
	LogonUserLogonTypes_Interactive       LogonUserLogonTypes = "interactive"
	LogonUserLogonTypes_Network           LogonUserLogonTypes = "network"
	LogonUserLogonTypes_RemoteInteractive LogonUserLogonTypes = "remoteInteractive"
	LogonUserLogonTypes_Service           LogonUserLogonTypes = "service"
	LogonUserLogonTypes_Unknown           LogonUserLogonTypes = "unknown"
)

func PossibleValuesForLogonUserLogonTypes() []string {
	return []string{
		string(LogonUserLogonTypes_Batch),
		string(LogonUserLogonTypes_Interactive),
		string(LogonUserLogonTypes_Network),
		string(LogonUserLogonTypes_RemoteInteractive),
		string(LogonUserLogonTypes_Service),
		string(LogonUserLogonTypes_Unknown),
	}
}

func (s *LogonUserLogonTypes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogonUserLogonTypes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogonUserLogonTypes(input string) (*LogonUserLogonTypes, error) {
	vals := map[string]LogonUserLogonTypes{
		"batch":             LogonUserLogonTypes_Batch,
		"interactive":       LogonUserLogonTypes_Interactive,
		"network":           LogonUserLogonTypes_Network,
		"remoteinteractive": LogonUserLogonTypes_RemoteInteractive,
		"service":           LogonUserLogonTypes_Service,
		"unknown":           LogonUserLogonTypes_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogonUserLogonTypes(input)
	return &out, nil
}
