package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidenceRemediationStatus string

const (
	SecurityKubernetesServiceAccountEvidenceRemediationStatus_Blocked    SecurityKubernetesServiceAccountEvidenceRemediationStatus = "blocked"
	SecurityKubernetesServiceAccountEvidenceRemediationStatus_None       SecurityKubernetesServiceAccountEvidenceRemediationStatus = "none"
	SecurityKubernetesServiceAccountEvidenceRemediationStatus_NotFound   SecurityKubernetesServiceAccountEvidenceRemediationStatus = "notFound"
	SecurityKubernetesServiceAccountEvidenceRemediationStatus_Prevented  SecurityKubernetesServiceAccountEvidenceRemediationStatus = "prevented"
	SecurityKubernetesServiceAccountEvidenceRemediationStatus_Remediated SecurityKubernetesServiceAccountEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesServiceAccountEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesServiceAccountEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesServiceAccountEvidenceRemediationStatus_None),
		string(SecurityKubernetesServiceAccountEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesServiceAccountEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesServiceAccountEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesServiceAccountEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceAccountEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceAccountEvidenceRemediationStatus(input string) (*SecurityKubernetesServiceAccountEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesServiceAccountEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesServiceAccountEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesServiceAccountEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesServiceAccountEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesServiceAccountEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesServiceAccountEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceAccountEvidenceRemediationStatus(input)
	return &out, nil
}
