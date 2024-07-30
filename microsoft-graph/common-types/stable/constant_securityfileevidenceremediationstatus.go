package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceRemediationStatus string

const (
	SecurityFileEvidenceRemediationStatus_Blocked    SecurityFileEvidenceRemediationStatus = "blocked"
	SecurityFileEvidenceRemediationStatus_None       SecurityFileEvidenceRemediationStatus = "none"
	SecurityFileEvidenceRemediationStatus_NotFound   SecurityFileEvidenceRemediationStatus = "notFound"
	SecurityFileEvidenceRemediationStatus_Prevented  SecurityFileEvidenceRemediationStatus = "prevented"
	SecurityFileEvidenceRemediationStatus_Remediated SecurityFileEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityFileEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityFileEvidenceRemediationStatus_Blocked),
		string(SecurityFileEvidenceRemediationStatus_None),
		string(SecurityFileEvidenceRemediationStatus_NotFound),
		string(SecurityFileEvidenceRemediationStatus_Prevented),
		string(SecurityFileEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityFileEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileEvidenceRemediationStatus(input string) (*SecurityFileEvidenceRemediationStatus, error) {
	vals := map[string]SecurityFileEvidenceRemediationStatus{
		"blocked":    SecurityFileEvidenceRemediationStatus_Blocked,
		"none":       SecurityFileEvidenceRemediationStatus_None,
		"notfound":   SecurityFileEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityFileEvidenceRemediationStatus_Prevented,
		"remediated": SecurityFileEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceRemediationStatus(input)
	return &out, nil
}
