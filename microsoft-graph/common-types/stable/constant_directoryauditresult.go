package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryAuditResult string

const (
	DirectoryAuditResult_Failure DirectoryAuditResult = "failure"
	DirectoryAuditResult_Success DirectoryAuditResult = "success"
	DirectoryAuditResult_Timeout DirectoryAuditResult = "timeout"
)

func PossibleValuesForDirectoryAuditResult() []string {
	return []string{
		string(DirectoryAuditResult_Failure),
		string(DirectoryAuditResult_Success),
		string(DirectoryAuditResult_Timeout),
	}
}

func (s *DirectoryAuditResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDirectoryAuditResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDirectoryAuditResult(input string) (*DirectoryAuditResult, error) {
	vals := map[string]DirectoryAuditResult{
		"failure": DirectoryAuditResult_Failure,
		"success": DirectoryAuditResult_Success,
		"timeout": DirectoryAuditResult_Timeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DirectoryAuditResult(input)
	return &out, nil
}
