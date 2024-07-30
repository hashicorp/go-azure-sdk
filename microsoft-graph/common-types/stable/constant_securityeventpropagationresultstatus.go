package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEventPropagationResultStatus string

const (
	SecurityEventPropagationResultStatus_Failed       SecurityEventPropagationResultStatus = "failed"
	SecurityEventPropagationResultStatus_InProcessing SecurityEventPropagationResultStatus = "inProcessing"
	SecurityEventPropagationResultStatus_None         SecurityEventPropagationResultStatus = "none"
	SecurityEventPropagationResultStatus_Success      SecurityEventPropagationResultStatus = "success"
)

func PossibleValuesForSecurityEventPropagationResultStatus() []string {
	return []string{
		string(SecurityEventPropagationResultStatus_Failed),
		string(SecurityEventPropagationResultStatus_InProcessing),
		string(SecurityEventPropagationResultStatus_None),
		string(SecurityEventPropagationResultStatus_Success),
	}
}

func (s *SecurityEventPropagationResultStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEventPropagationResultStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEventPropagationResultStatus(input string) (*SecurityEventPropagationResultStatus, error) {
	vals := map[string]SecurityEventPropagationResultStatus{
		"failed":       SecurityEventPropagationResultStatus_Failed,
		"inprocessing": SecurityEventPropagationResultStatus_InProcessing,
		"none":         SecurityEventPropagationResultStatus_None,
		"success":      SecurityEventPropagationResultStatus_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEventPropagationResultStatus(input)
	return &out, nil
}
