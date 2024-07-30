package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlThreatSubmissionStatus string

const (
	SecurityUrlThreatSubmissionStatus_Failed     SecurityUrlThreatSubmissionStatus = "failed"
	SecurityUrlThreatSubmissionStatus_NotStarted SecurityUrlThreatSubmissionStatus = "notStarted"
	SecurityUrlThreatSubmissionStatus_Running    SecurityUrlThreatSubmissionStatus = "running"
	SecurityUrlThreatSubmissionStatus_Skipped    SecurityUrlThreatSubmissionStatus = "skipped"
	SecurityUrlThreatSubmissionStatus_Succeeded  SecurityUrlThreatSubmissionStatus = "succeeded"
)

func PossibleValuesForSecurityUrlThreatSubmissionStatus() []string {
	return []string{
		string(SecurityUrlThreatSubmissionStatus_Failed),
		string(SecurityUrlThreatSubmissionStatus_NotStarted),
		string(SecurityUrlThreatSubmissionStatus_Running),
		string(SecurityUrlThreatSubmissionStatus_Skipped),
		string(SecurityUrlThreatSubmissionStatus_Succeeded),
	}
}

func (s *SecurityUrlThreatSubmissionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlThreatSubmissionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlThreatSubmissionStatus(input string) (*SecurityUrlThreatSubmissionStatus, error) {
	vals := map[string]SecurityUrlThreatSubmissionStatus{
		"failed":     SecurityUrlThreatSubmissionStatus_Failed,
		"notstarted": SecurityUrlThreatSubmissionStatus_NotStarted,
		"running":    SecurityUrlThreatSubmissionStatus_Running,
		"skipped":    SecurityUrlThreatSubmissionStatus_Skipped,
		"succeeded":  SecurityUrlThreatSubmissionStatus_Succeeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlThreatSubmissionStatus(input)
	return &out, nil
}
