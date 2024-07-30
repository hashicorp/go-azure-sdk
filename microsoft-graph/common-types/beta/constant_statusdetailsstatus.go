package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusDetailsStatus string

const (
	StatusDetailsStatus_Failure StatusDetailsStatus = "failure"
	StatusDetailsStatus_Skipped StatusDetailsStatus = "skipped"
	StatusDetailsStatus_Success StatusDetailsStatus = "success"
	StatusDetailsStatus_Warning StatusDetailsStatus = "warning"
)

func PossibleValuesForStatusDetailsStatus() []string {
	return []string{
		string(StatusDetailsStatus_Failure),
		string(StatusDetailsStatus_Skipped),
		string(StatusDetailsStatus_Success),
		string(StatusDetailsStatus_Warning),
	}
}

func (s *StatusDetailsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStatusDetailsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStatusDetailsStatus(input string) (*StatusDetailsStatus, error) {
	vals := map[string]StatusDetailsStatus{
		"failure": StatusDetailsStatus_Failure,
		"skipped": StatusDetailsStatus_Skipped,
		"success": StatusDetailsStatus_Success,
		"warning": StatusDetailsStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusDetailsStatus(input)
	return &out, nil
}
