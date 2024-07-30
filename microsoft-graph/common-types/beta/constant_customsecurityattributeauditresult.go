package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeAuditResult string

const (
	CustomSecurityAttributeAuditResult_Failure CustomSecurityAttributeAuditResult = "failure"
	CustomSecurityAttributeAuditResult_Success CustomSecurityAttributeAuditResult = "success"
	CustomSecurityAttributeAuditResult_Timeout CustomSecurityAttributeAuditResult = "timeout"
)

func PossibleValuesForCustomSecurityAttributeAuditResult() []string {
	return []string{
		string(CustomSecurityAttributeAuditResult_Failure),
		string(CustomSecurityAttributeAuditResult_Success),
		string(CustomSecurityAttributeAuditResult_Timeout),
	}
}

func (s *CustomSecurityAttributeAuditResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomSecurityAttributeAuditResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomSecurityAttributeAuditResult(input string) (*CustomSecurityAttributeAuditResult, error) {
	vals := map[string]CustomSecurityAttributeAuditResult{
		"failure": CustomSecurityAttributeAuditResult_Failure,
		"success": CustomSecurityAttributeAuditResult_Success,
		"timeout": CustomSecurityAttributeAuditResult_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomSecurityAttributeAuditResult(input)
	return &out, nil
}
