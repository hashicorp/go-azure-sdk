package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailUrlThreatSubmissionStatus string

const (
	SecurityEmailUrlThreatSubmissionStatus_Failed     SecurityEmailUrlThreatSubmissionStatus = "failed"
	SecurityEmailUrlThreatSubmissionStatus_NotStarted SecurityEmailUrlThreatSubmissionStatus = "notStarted"
	SecurityEmailUrlThreatSubmissionStatus_Running    SecurityEmailUrlThreatSubmissionStatus = "running"
	SecurityEmailUrlThreatSubmissionStatus_Skipped    SecurityEmailUrlThreatSubmissionStatus = "skipped"
	SecurityEmailUrlThreatSubmissionStatus_Succeeded  SecurityEmailUrlThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityEmailUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailUrlThreatSubmissionStatus_Failed),
		string(SecurityEmailUrlThreatSubmissionStatus_NotStarted),
		string(SecurityEmailUrlThreatSubmissionStatus_Running),
		string(SecurityEmailUrlThreatSubmissionStatus_Skipped),
		string(SecurityEmailUrlThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityEmailUrlThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailUrlThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailUrlThreatSubmissionStatus(input string) (*SecurityEmailUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailUrlThreatSubmissionStatus{
		"failed":     SecurityEmailUrlThreatSubmissionStatus_Failed,
		"notstarted": SecurityEmailUrlThreatSubmissionStatus_NotStarted,
		"running":    SecurityEmailUrlThreatSubmissionStatus_Running,
		"skipped":    SecurityEmailUrlThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityEmailUrlThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailUrlThreatSubmissionStatus(input)
	return &out, nil
}
