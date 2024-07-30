package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidenceRemediationStatus string

const (
	SecurityKubernetesClusterEvidenceRemediationStatus_Blocked    SecurityKubernetesClusterEvidenceRemediationStatus = "blocked"
	SecurityKubernetesClusterEvidenceRemediationStatus_None       SecurityKubernetesClusterEvidenceRemediationStatus = "none"
	SecurityKubernetesClusterEvidenceRemediationStatus_NotFound   SecurityKubernetesClusterEvidenceRemediationStatus = "notFound"
	SecurityKubernetesClusterEvidenceRemediationStatus_Prevented  SecurityKubernetesClusterEvidenceRemediationStatus = "prevented"
	SecurityKubernetesClusterEvidenceRemediationStatus_Remediated SecurityKubernetesClusterEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesClusterEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesClusterEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesClusterEvidenceRemediationStatus_None),
		string(SecurityKubernetesClusterEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesClusterEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesClusterEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesClusterEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesClusterEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesClusterEvidenceRemediationStatus(input string) (*SecurityKubernetesClusterEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesClusterEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesClusterEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesClusterEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesClusterEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesClusterEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesClusterEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidenceRemediationStatus(input)
	return &out, nil
}
