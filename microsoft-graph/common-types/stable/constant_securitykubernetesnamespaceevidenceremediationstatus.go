package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceRemediationStatus string

const (
	SecurityKubernetesNamespaceEvidenceRemediationStatus_Blocked    SecurityKubernetesNamespaceEvidenceRemediationStatus = "blocked"
	SecurityKubernetesNamespaceEvidenceRemediationStatus_None       SecurityKubernetesNamespaceEvidenceRemediationStatus = "none"
	SecurityKubernetesNamespaceEvidenceRemediationStatus_NotFound   SecurityKubernetesNamespaceEvidenceRemediationStatus = "notFound"
	SecurityKubernetesNamespaceEvidenceRemediationStatus_Prevented  SecurityKubernetesNamespaceEvidenceRemediationStatus = "prevented"
	SecurityKubernetesNamespaceEvidenceRemediationStatus_Remediated SecurityKubernetesNamespaceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatus_None),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesNamespaceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesNamespaceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesNamespaceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesNamespaceEvidenceRemediationStatus(input string) (*SecurityKubernetesNamespaceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesNamespaceEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesNamespaceEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesNamespaceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesNamespaceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesNamespaceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceRemediationStatus(input)
	return &out, nil
}
