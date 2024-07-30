package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesSecretEvidenceRemediationStatus string

const (
	SecurityKubernetesSecretEvidenceRemediationStatus_Blocked    SecurityKubernetesSecretEvidenceRemediationStatus = "blocked"
	SecurityKubernetesSecretEvidenceRemediationStatus_None       SecurityKubernetesSecretEvidenceRemediationStatus = "none"
	SecurityKubernetesSecretEvidenceRemediationStatus_NotFound   SecurityKubernetesSecretEvidenceRemediationStatus = "notFound"
	SecurityKubernetesSecretEvidenceRemediationStatus_Prevented  SecurityKubernetesSecretEvidenceRemediationStatus = "prevented"
	SecurityKubernetesSecretEvidenceRemediationStatus_Remediated SecurityKubernetesSecretEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesSecretEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesSecretEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesSecretEvidenceRemediationStatus_None),
		string(SecurityKubernetesSecretEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesSecretEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesSecretEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesSecretEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesSecretEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesSecretEvidenceRemediationStatus(input string) (*SecurityKubernetesSecretEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesSecretEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesSecretEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesSecretEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesSecretEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesSecretEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesSecretEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesSecretEvidenceRemediationStatus(input)
	return &out, nil
}
