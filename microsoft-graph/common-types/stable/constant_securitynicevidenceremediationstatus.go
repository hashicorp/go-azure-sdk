package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNicEvidenceRemediationStatus string

const (
	SecurityNicEvidenceRemediationStatus_Blocked    SecurityNicEvidenceRemediationStatus = "blocked"
	SecurityNicEvidenceRemediationStatus_None       SecurityNicEvidenceRemediationStatus = "none"
	SecurityNicEvidenceRemediationStatus_NotFound   SecurityNicEvidenceRemediationStatus = "notFound"
	SecurityNicEvidenceRemediationStatus_Prevented  SecurityNicEvidenceRemediationStatus = "prevented"
	SecurityNicEvidenceRemediationStatus_Remediated SecurityNicEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityNicEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityNicEvidenceRemediationStatus_Blocked),
		string(SecurityNicEvidenceRemediationStatus_None),
		string(SecurityNicEvidenceRemediationStatus_NotFound),
		string(SecurityNicEvidenceRemediationStatus_Prevented),
		string(SecurityNicEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityNicEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNicEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNicEvidenceRemediationStatus(input string) (*SecurityNicEvidenceRemediationStatus, error) {
	vals := map[string]SecurityNicEvidenceRemediationStatus{
		"blocked":    SecurityNicEvidenceRemediationStatus_Blocked,
		"none":       SecurityNicEvidenceRemediationStatus_None,
		"notfound":   SecurityNicEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityNicEvidenceRemediationStatus_Prevented,
		"remediated": SecurityNicEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNicEvidenceRemediationStatus(input)
	return &out, nil
}
