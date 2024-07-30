package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidenceRemediationStatus string

const (
	SecurityMailClusterEvidenceRemediationStatus_Blocked    SecurityMailClusterEvidenceRemediationStatus = "blocked"
	SecurityMailClusterEvidenceRemediationStatus_None       SecurityMailClusterEvidenceRemediationStatus = "none"
	SecurityMailClusterEvidenceRemediationStatus_NotFound   SecurityMailClusterEvidenceRemediationStatus = "notFound"
	SecurityMailClusterEvidenceRemediationStatus_Prevented  SecurityMailClusterEvidenceRemediationStatus = "prevented"
	SecurityMailClusterEvidenceRemediationStatus_Remediated SecurityMailClusterEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityMailClusterEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityMailClusterEvidenceRemediationStatus_Blocked),
		string(SecurityMailClusterEvidenceRemediationStatus_None),
		string(SecurityMailClusterEvidenceRemediationStatus_NotFound),
		string(SecurityMailClusterEvidenceRemediationStatus_Prevented),
		string(SecurityMailClusterEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityMailClusterEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailClusterEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailClusterEvidenceRemediationStatus(input string) (*SecurityMailClusterEvidenceRemediationStatus, error) {
	vals := map[string]SecurityMailClusterEvidenceRemediationStatus{
		"blocked":    SecurityMailClusterEvidenceRemediationStatus_Blocked,
		"none":       SecurityMailClusterEvidenceRemediationStatus_None,
		"notfound":   SecurityMailClusterEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityMailClusterEvidenceRemediationStatus_Prevented,
		"remediated": SecurityMailClusterEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailClusterEvidenceRemediationStatus(input)
	return &out, nil
}
