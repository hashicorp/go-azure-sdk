package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceRemediationStatus string

const (
	SecurityKubernetesServiceEvidenceRemediationStatus_Blocked    SecurityKubernetesServiceEvidenceRemediationStatus = "blocked"
	SecurityKubernetesServiceEvidenceRemediationStatus_None       SecurityKubernetesServiceEvidenceRemediationStatus = "none"
	SecurityKubernetesServiceEvidenceRemediationStatus_NotFound   SecurityKubernetesServiceEvidenceRemediationStatus = "notFound"
	SecurityKubernetesServiceEvidenceRemediationStatus_Prevented  SecurityKubernetesServiceEvidenceRemediationStatus = "prevented"
	SecurityKubernetesServiceEvidenceRemediationStatus_Remediated SecurityKubernetesServiceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesServiceEvidenceRemediationStatus_None),
		string(SecurityKubernetesServiceEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesServiceEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesServiceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesServiceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceEvidenceRemediationStatus(input string) (*SecurityKubernetesServiceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesServiceEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesServiceEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesServiceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesServiceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesServiceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceRemediationStatus(input)
	return &out, nil
}
