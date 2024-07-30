package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailContentThreatSubmissionStatus string

const (
	SecurityEmailContentThreatSubmissionStatus_Failed     SecurityEmailContentThreatSubmissionStatus = "failed"
	SecurityEmailContentThreatSubmissionStatus_NotStarted SecurityEmailContentThreatSubmissionStatus = "notStarted"
	SecurityEmailContentThreatSubmissionStatus_Running    SecurityEmailContentThreatSubmissionStatus = "running"
	SecurityEmailContentThreatSubmissionStatus_Skipped    SecurityEmailContentThreatSubmissionStatus = "skipped"
	SecurityEmailContentThreatSubmissionStatus_Succeeded  SecurityEmailContentThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityEmailContentThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailContentThreatSubmissionStatus_Failed),
		string(SecurityEmailContentThreatSubmissionStatus_NotStarted),
		string(SecurityEmailContentThreatSubmissionStatus_Running),
		string(SecurityEmailContentThreatSubmissionStatus_Skipped),
		string(SecurityEmailContentThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityEmailContentThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailContentThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailContentThreatSubmissionStatus(input string) (*SecurityEmailContentThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailContentThreatSubmissionStatus{
		"failed":     SecurityEmailContentThreatSubmissionStatus_Failed,
		"notstarted": SecurityEmailContentThreatSubmissionStatus_NotStarted,
		"running":    SecurityEmailContentThreatSubmissionStatus_Running,
		"skipped":    SecurityEmailContentThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityEmailContentThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailContentThreatSubmissionStatus(input)
	return &out, nil
}
