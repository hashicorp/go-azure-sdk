package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostLogonSessionEvidenceRemediationStatus string

const (
	SecurityHostLogonSessionEvidenceRemediationStatus_Blocked    SecurityHostLogonSessionEvidenceRemediationStatus = "blocked"
	SecurityHostLogonSessionEvidenceRemediationStatus_None       SecurityHostLogonSessionEvidenceRemediationStatus = "none"
	SecurityHostLogonSessionEvidenceRemediationStatus_NotFound   SecurityHostLogonSessionEvidenceRemediationStatus = "notFound"
	SecurityHostLogonSessionEvidenceRemediationStatus_Prevented  SecurityHostLogonSessionEvidenceRemediationStatus = "prevented"
	SecurityHostLogonSessionEvidenceRemediationStatus_Remediated SecurityHostLogonSessionEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityHostLogonSessionEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityHostLogonSessionEvidenceRemediationStatus_Blocked),
		string(SecurityHostLogonSessionEvidenceRemediationStatus_None),
		string(SecurityHostLogonSessionEvidenceRemediationStatus_NotFound),
		string(SecurityHostLogonSessionEvidenceRemediationStatus_Prevented),
		string(SecurityHostLogonSessionEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityHostLogonSessionEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostLogonSessionEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostLogonSessionEvidenceRemediationStatus(input string) (*SecurityHostLogonSessionEvidenceRemediationStatus, error) {
	vals := map[string]SecurityHostLogonSessionEvidenceRemediationStatus{
		"blocked":    SecurityHostLogonSessionEvidenceRemediationStatus_Blocked,
		"none":       SecurityHostLogonSessionEvidenceRemediationStatus_None,
		"notfound":   SecurityHostLogonSessionEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityHostLogonSessionEvidenceRemediationStatus_Prevented,
		"remediated": SecurityHostLogonSessionEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostLogonSessionEvidenceRemediationStatus(input)
	return &out, nil
}
