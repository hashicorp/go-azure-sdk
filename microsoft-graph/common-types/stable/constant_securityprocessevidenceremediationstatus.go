package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceRemediationStatus string

const (
	SecurityProcessEvidenceRemediationStatus_Blocked    SecurityProcessEvidenceRemediationStatus = "blocked"
	SecurityProcessEvidenceRemediationStatus_None       SecurityProcessEvidenceRemediationStatus = "none"
	SecurityProcessEvidenceRemediationStatus_NotFound   SecurityProcessEvidenceRemediationStatus = "notFound"
	SecurityProcessEvidenceRemediationStatus_Prevented  SecurityProcessEvidenceRemediationStatus = "prevented"
	SecurityProcessEvidenceRemediationStatus_Remediated SecurityProcessEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityProcessEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityProcessEvidenceRemediationStatus_Blocked),
		string(SecurityProcessEvidenceRemediationStatus_None),
		string(SecurityProcessEvidenceRemediationStatus_NotFound),
		string(SecurityProcessEvidenceRemediationStatus_Prevented),
		string(SecurityProcessEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityProcessEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityProcessEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityProcessEvidenceRemediationStatus(input string) (*SecurityProcessEvidenceRemediationStatus, error) {
	vals := map[string]SecurityProcessEvidenceRemediationStatus{
		"blocked":    SecurityProcessEvidenceRemediationStatus_Blocked,
		"none":       SecurityProcessEvidenceRemediationStatus_None,
		"notfound":   SecurityProcessEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityProcessEvidenceRemediationStatus_Prevented,
		"remediated": SecurityProcessEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceRemediationStatus(input)
	return &out, nil
}
