package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUserAccountType string

const (
	LogonUserAccountType_Administrator LogonUserAccountType = "administrator"
	LogonUserAccountType_Power         LogonUserAccountType = "power"
	LogonUserAccountType_Standard      LogonUserAccountType = "standard"
	LogonUserAccountType_Unknown       LogonUserAccountType = "unknown"
)

func PossibleValuesForLogonUserAccountType() []string {
	return []string{
		string(LogonUserAccountType_Administrator),
		string(LogonUserAccountType_Power),
		string(LogonUserAccountType_Standard),
		string(LogonUserAccountType_Unknown),
	}
}

func (s *LogonUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLogonUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLogonUserAccountType(input string) (*LogonUserAccountType, error) {
	vals := map[string]LogonUserAccountType{
		"administrator": LogonUserAccountType_Administrator,
		"power":         LogonUserAccountType_Power,
		"standard":      LogonUserAccountType_Standard,
		"unknown":       LogonUserAccountType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LogonUserAccountType(input)
	return &out, nil
}
