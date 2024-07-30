package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileContentThreatSubmissionStatus string

const (
	SecurityFileContentThreatSubmissionStatus_Failed     SecurityFileContentThreatSubmissionStatus = "failed"
	SecurityFileContentThreatSubmissionStatus_NotStarted SecurityFileContentThreatSubmissionStatus = "notStarted"
	SecurityFileContentThreatSubmissionStatus_Running    SecurityFileContentThreatSubmissionStatus = "running"
	SecurityFileContentThreatSubmissionStatus_Skipped    SecurityFileContentThreatSubmissionStatus = "skipped"
	SecurityFileContentThreatSubmissionStatus_Succeeded  SecurityFileContentThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityFileContentThreatSubmissionStatus() []string {
	return []string{
		string(SecurityFileContentThreatSubmissionStatus_Failed),
		string(SecurityFileContentThreatSubmissionStatus_NotStarted),
		string(SecurityFileContentThreatSubmissionStatus_Running),
		string(SecurityFileContentThreatSubmissionStatus_Skipped),
		string(SecurityFileContentThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityFileContentThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileContentThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileContentThreatSubmissionStatus(input string) (*SecurityFileContentThreatSubmissionStatus, error) {
	vals := map[string]SecurityFileContentThreatSubmissionStatus{
		"failed":     SecurityFileContentThreatSubmissionStatus_Failed,
		"notstarted": SecurityFileContentThreatSubmissionStatus_NotStarted,
		"running":    SecurityFileContentThreatSubmissionStatus_Running,
		"skipped":    SecurityFileContentThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityFileContentThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileContentThreatSubmissionStatus(input)
	return &out, nil
}
