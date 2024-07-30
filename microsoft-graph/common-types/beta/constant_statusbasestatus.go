package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusBaseStatus string

const (
	StatusBaseStatus_Failure StatusBaseStatus = "failure"
	StatusBaseStatus_Skipped StatusBaseStatus = "skipped"
	StatusBaseStatus_Success StatusBaseStatus = "success"
	StatusBaseStatus_Warning StatusBaseStatus = "warning"
)

func PossibleValuesForStatusBaseStatus() []string {
	return []string{
		string(StatusBaseStatus_Failure),
		string(StatusBaseStatus_Skipped),
		string(StatusBaseStatus_Success),
		string(StatusBaseStatus_Warning),
	}
}

func (s *StatusBaseStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusBaseStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusBaseStatus(input string) (*StatusBaseStatus, error) {
	vals := map[string]StatusBaseStatus{
		"failure": StatusBaseStatus_Failure,
		"skipped": StatusBaseStatus_Skipped,
		"success": StatusBaseStatus_Success,
		"warning": StatusBaseStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusBaseStatus(input)
	return &out, nil
}
