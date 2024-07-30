package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityThreatSubmissionStatus string

const (
	SecurityThreatSubmissionStatus_Failed     SecurityThreatSubmissionStatus = "failed"
	SecurityThreatSubmissionStatus_NotStarted SecurityThreatSubmissionStatus = "notStarted"
	SecurityThreatSubmissionStatus_Running    SecurityThreatSubmissionStatus = "running"
	SecurityThreatSubmissionStatus_Skipped    SecurityThreatSubmissionStatus = "skipped"
	SecurityThreatSubmissionStatus_Succeeded  SecurityThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityThreatSubmissionStatus() []string {
	return []string{
		string(SecurityThreatSubmissionStatus_Failed),
		string(SecurityThreatSubmissionStatus_NotStarted),
		string(SecurityThreatSubmissionStatus_Running),
		string(SecurityThreatSubmissionStatus_Skipped),
		string(SecurityThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityThreatSubmissionStatus(input string) (*SecurityThreatSubmissionStatus, error) {
	vals := map[string]SecurityThreatSubmissionStatus{
		"failed":     SecurityThreatSubmissionStatus_Failed,
		"notstarted": SecurityThreatSubmissionStatus_NotStarted,
		"running":    SecurityThreatSubmissionStatus_Running,
		"skipped":    SecurityThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityThreatSubmissionStatus(input)
	return &out, nil
}
