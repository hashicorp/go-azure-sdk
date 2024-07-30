package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceRemediationStatus string

const (
	SecurityIpEvidenceRemediationStatus_Blocked    SecurityIpEvidenceRemediationStatus = "blocked"
	SecurityIpEvidenceRemediationStatus_None       SecurityIpEvidenceRemediationStatus = "none"
	SecurityIpEvidenceRemediationStatus_NotFound   SecurityIpEvidenceRemediationStatus = "notFound"
	SecurityIpEvidenceRemediationStatus_Prevented  SecurityIpEvidenceRemediationStatus = "prevented"
	SecurityIpEvidenceRemediationStatus_Remediated SecurityIpEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityIpEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityIpEvidenceRemediationStatus_Blocked),
		string(SecurityIpEvidenceRemediationStatus_None),
		string(SecurityIpEvidenceRemediationStatus_NotFound),
		string(SecurityIpEvidenceRemediationStatus_Prevented),
		string(SecurityIpEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityIpEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIpEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIpEvidenceRemediationStatus(input string) (*SecurityIpEvidenceRemediationStatus, error) {
	vals := map[string]SecurityIpEvidenceRemediationStatus{
		"blocked":    SecurityIpEvidenceRemediationStatus_Blocked,
		"none":       SecurityIpEvidenceRemediationStatus_None,
		"notfound":   SecurityIpEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityIpEvidenceRemediationStatus_Prevented,
		"remediated": SecurityIpEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceRemediationStatus(input)
	return &out, nil
}
