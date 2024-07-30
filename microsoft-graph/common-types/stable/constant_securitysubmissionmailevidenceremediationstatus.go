package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionMailEvidenceRemediationStatus string

const (
	SecuritySubmissionMailEvidenceRemediationStatus_Blocked    SecuritySubmissionMailEvidenceRemediationStatus = "blocked"
	SecuritySubmissionMailEvidenceRemediationStatus_None       SecuritySubmissionMailEvidenceRemediationStatus = "none"
	SecuritySubmissionMailEvidenceRemediationStatus_NotFound   SecuritySubmissionMailEvidenceRemediationStatus = "notFound"
	SecuritySubmissionMailEvidenceRemediationStatus_Prevented  SecuritySubmissionMailEvidenceRemediationStatus = "prevented"
	SecuritySubmissionMailEvidenceRemediationStatus_Remediated SecuritySubmissionMailEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecuritySubmissionMailEvidenceRemediationStatus() []string {
	return []string{
		string(SecuritySubmissionMailEvidenceRemediationStatus_Blocked),
		string(SecuritySubmissionMailEvidenceRemediationStatus_None),
		string(SecuritySubmissionMailEvidenceRemediationStatus_NotFound),
		string(SecuritySubmissionMailEvidenceRemediationStatus_Prevented),
		string(SecuritySubmissionMailEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecuritySubmissionMailEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionMailEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionMailEvidenceRemediationStatus(input string) (*SecuritySubmissionMailEvidenceRemediationStatus, error) {
	vals := map[string]SecuritySubmissionMailEvidenceRemediationStatus{
		"blocked":    SecuritySubmissionMailEvidenceRemediationStatus_Blocked,
		"none":       SecuritySubmissionMailEvidenceRemediationStatus_None,
		"notfound":   SecuritySubmissionMailEvidenceRemediationStatus_NotFound,
		"prevented":  SecuritySubmissionMailEvidenceRemediationStatus_Prevented,
		"remediated": SecuritySubmissionMailEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionMailEvidenceRemediationStatus(input)
	return &out, nil
}
