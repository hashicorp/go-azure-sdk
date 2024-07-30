package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileThreatSubmissionStatus string

const (
	SecurityFileThreatSubmissionStatus_Failed     SecurityFileThreatSubmissionStatus = "failed"
	SecurityFileThreatSubmissionStatus_NotStarted SecurityFileThreatSubmissionStatus = "notStarted"
	SecurityFileThreatSubmissionStatus_Running    SecurityFileThreatSubmissionStatus = "running"
	SecurityFileThreatSubmissionStatus_Skipped    SecurityFileThreatSubmissionStatus = "skipped"
	SecurityFileThreatSubmissionStatus_Succeeded  SecurityFileThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityFileThreatSubmissionStatus() []string {
	return []string{
		string(SecurityFileThreatSubmissionStatus_Failed),
		string(SecurityFileThreatSubmissionStatus_NotStarted),
		string(SecurityFileThreatSubmissionStatus_Running),
		string(SecurityFileThreatSubmissionStatus_Skipped),
		string(SecurityFileThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityFileThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileThreatSubmissionStatus(input string) (*SecurityFileThreatSubmissionStatus, error) {
	vals := map[string]SecurityFileThreatSubmissionStatus{
		"failed":     SecurityFileThreatSubmissionStatus_Failed,
		"notstarted": SecurityFileThreatSubmissionStatus_NotStarted,
		"running":    SecurityFileThreatSubmissionStatus_Running,
		"skipped":    SecurityFileThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityFileThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileThreatSubmissionStatus(input)
	return &out, nil
}
