package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileUrlThreatSubmissionStatus string

const (
	SecurityFileUrlThreatSubmissionStatus_Failed     SecurityFileUrlThreatSubmissionStatus = "failed"
	SecurityFileUrlThreatSubmissionStatus_NotStarted SecurityFileUrlThreatSubmissionStatus = "notStarted"
	SecurityFileUrlThreatSubmissionStatus_Running    SecurityFileUrlThreatSubmissionStatus = "running"
	SecurityFileUrlThreatSubmissionStatus_Skipped    SecurityFileUrlThreatSubmissionStatus = "skipped"
	SecurityFileUrlThreatSubmissionStatus_Succeeded  SecurityFileUrlThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityFileUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityFileUrlThreatSubmissionStatus_Failed),
		string(SecurityFileUrlThreatSubmissionStatus_NotStarted),
		string(SecurityFileUrlThreatSubmissionStatus_Running),
		string(SecurityFileUrlThreatSubmissionStatus_Skipped),
		string(SecurityFileUrlThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityFileUrlThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileUrlThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileUrlThreatSubmissionStatus(input string) (*SecurityFileUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityFileUrlThreatSubmissionStatus{
		"failed":     SecurityFileUrlThreatSubmissionStatus_Failed,
		"notstarted": SecurityFileUrlThreatSubmissionStatus_NotStarted,
		"running":    SecurityFileUrlThreatSubmissionStatus_Running,
		"skipped":    SecurityFileUrlThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityFileUrlThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileUrlThreatSubmissionStatus(input)
	return &out, nil
}
