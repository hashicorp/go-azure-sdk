package triggeredwebjoboperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggeredWebJobStatus string

const (
	TriggeredWebJobStatusError   TriggeredWebJobStatus = "Error"
	TriggeredWebJobStatusFailed  TriggeredWebJobStatus = "Failed"
	TriggeredWebJobStatusSuccess TriggeredWebJobStatus = "Success"
)

func PossibleValuesForTriggeredWebJobStatus() []string {
	return []string{
		string(TriggeredWebJobStatusError),
		string(TriggeredWebJobStatusFailed),
		string(TriggeredWebJobStatusSuccess),
	}
}

func (s *TriggeredWebJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTriggeredWebJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTriggeredWebJobStatus(input string) (*TriggeredWebJobStatus, error) {
	vals := map[string]TriggeredWebJobStatus{
		"error":   TriggeredWebJobStatusError,
		"failed":  TriggeredWebJobStatusFailed,
		"success": TriggeredWebJobStatusSuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggeredWebJobStatus(input)
	return &out, nil
}

type WebJobType string

const (
	WebJobTypeContinuous WebJobType = "Continuous"
	WebJobTypeTriggered  WebJobType = "Triggered"
)

func PossibleValuesForWebJobType() []string {
	return []string{
		string(WebJobTypeContinuous),
		string(WebJobTypeTriggered),
	}
}

func (s *WebJobType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWebJobType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWebJobType(input string) (*WebJobType, error) {
	vals := map[string]WebJobType{
		"continuous": WebJobTypeContinuous,
		"triggered":  WebJobTypeTriggered,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WebJobType(input)
	return &out, nil
}
