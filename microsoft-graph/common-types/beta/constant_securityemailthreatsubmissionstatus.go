package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailThreatSubmissionStatus string

const (
	SecurityEmailThreatSubmissionStatus_Failed     SecurityEmailThreatSubmissionStatus = "failed"
	SecurityEmailThreatSubmissionStatus_NotStarted SecurityEmailThreatSubmissionStatus = "notStarted"
	SecurityEmailThreatSubmissionStatus_Running    SecurityEmailThreatSubmissionStatus = "running"
	SecurityEmailThreatSubmissionStatus_Skipped    SecurityEmailThreatSubmissionStatus = "skipped"
	SecurityEmailThreatSubmissionStatus_Succeeded  SecurityEmailThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityEmailThreatSubmissionStatus() []string {
	return []string{
		string(SecurityEmailThreatSubmissionStatus_Failed),
		string(SecurityEmailThreatSubmissionStatus_NotStarted),
		string(SecurityEmailThreatSubmissionStatus_Running),
		string(SecurityEmailThreatSubmissionStatus_Skipped),
		string(SecurityEmailThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityEmailThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEmailThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEmailThreatSubmissionStatus(input string) (*SecurityEmailThreatSubmissionStatus, error) {
	vals := map[string]SecurityEmailThreatSubmissionStatus{
		"failed":     SecurityEmailThreatSubmissionStatus_Failed,
		"notstarted": SecurityEmailThreatSubmissionStatus_NotStarted,
		"running":    SecurityEmailThreatSubmissionStatus_Running,
		"skipped":    SecurityEmailThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityEmailThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEmailThreatSubmissionStatus(input)
	return &out, nil
}
