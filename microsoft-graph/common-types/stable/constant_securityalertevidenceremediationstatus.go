package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceRemediationStatus string

const (
	SecurityAlertEvidenceRemediationStatus_Blocked    SecurityAlertEvidenceRemediationStatus = "blocked"
	SecurityAlertEvidenceRemediationStatus_None       SecurityAlertEvidenceRemediationStatus = "none"
	SecurityAlertEvidenceRemediationStatus_NotFound   SecurityAlertEvidenceRemediationStatus = "notFound"
	SecurityAlertEvidenceRemediationStatus_Prevented  SecurityAlertEvidenceRemediationStatus = "prevented"
	SecurityAlertEvidenceRemediationStatus_Remediated SecurityAlertEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityAlertEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAlertEvidenceRemediationStatus_Blocked),
		string(SecurityAlertEvidenceRemediationStatus_None),
		string(SecurityAlertEvidenceRemediationStatus_NotFound),
		string(SecurityAlertEvidenceRemediationStatus_Prevented),
		string(SecurityAlertEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityAlertEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertEvidenceRemediationStatus(input string) (*SecurityAlertEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAlertEvidenceRemediationStatus{
		"blocked":    SecurityAlertEvidenceRemediationStatus_Blocked,
		"none":       SecurityAlertEvidenceRemediationStatus_None,
		"notfound":   SecurityAlertEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityAlertEvidenceRemediationStatus_Prevented,
		"remediated": SecurityAlertEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceRemediationStatus(input)
	return &out, nil
}
