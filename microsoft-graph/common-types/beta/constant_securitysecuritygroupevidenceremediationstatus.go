package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceRemediationStatus string

const (
	SecuritySecurityGroupEvidenceRemediationStatus_Blocked    SecuritySecurityGroupEvidenceRemediationStatus = "blocked"
	SecuritySecurityGroupEvidenceRemediationStatus_None       SecuritySecurityGroupEvidenceRemediationStatus = "none"
	SecuritySecurityGroupEvidenceRemediationStatus_NotFound   SecuritySecurityGroupEvidenceRemediationStatus = "notFound"
	SecuritySecurityGroupEvidenceRemediationStatus_Prevented  SecuritySecurityGroupEvidenceRemediationStatus = "prevented"
	SecuritySecurityGroupEvidenceRemediationStatus_Remediated SecuritySecurityGroupEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecuritySecurityGroupEvidenceRemediationStatus() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceRemediationStatus_Blocked),
		string(SecuritySecurityGroupEvidenceRemediationStatus_None),
		string(SecuritySecurityGroupEvidenceRemediationStatus_NotFound),
		string(SecuritySecurityGroupEvidenceRemediationStatus_Prevented),
		string(SecuritySecurityGroupEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecuritySecurityGroupEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySecurityGroupEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySecurityGroupEvidenceRemediationStatus(input string) (*SecuritySecurityGroupEvidenceRemediationStatus, error) {
	vals := map[string]SecuritySecurityGroupEvidenceRemediationStatus{
		"blocked":    SecuritySecurityGroupEvidenceRemediationStatus_Blocked,
		"none":       SecuritySecurityGroupEvidenceRemediationStatus_None,
		"notfound":   SecuritySecurityGroupEvidenceRemediationStatus_NotFound,
		"prevented":  SecuritySecurityGroupEvidenceRemediationStatus_Prevented,
		"remediated": SecuritySecurityGroupEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceRemediationStatus(input)
	return &out, nil
}
