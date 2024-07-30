package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidenceRemediationStatus string

const (
	SecurityKubernetesControllerEvidenceRemediationStatus_Blocked    SecurityKubernetesControllerEvidenceRemediationStatus = "blocked"
	SecurityKubernetesControllerEvidenceRemediationStatus_None       SecurityKubernetesControllerEvidenceRemediationStatus = "none"
	SecurityKubernetesControllerEvidenceRemediationStatus_NotFound   SecurityKubernetesControllerEvidenceRemediationStatus = "notFound"
	SecurityKubernetesControllerEvidenceRemediationStatus_Prevented  SecurityKubernetesControllerEvidenceRemediationStatus = "prevented"
	SecurityKubernetesControllerEvidenceRemediationStatus_Remediated SecurityKubernetesControllerEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesControllerEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesControllerEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesControllerEvidenceRemediationStatus_None),
		string(SecurityKubernetesControllerEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesControllerEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesControllerEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesControllerEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesControllerEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesControllerEvidenceRemediationStatus(input string) (*SecurityKubernetesControllerEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesControllerEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesControllerEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesControllerEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesControllerEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesControllerEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesControllerEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesControllerEvidenceRemediationStatus(input)
	return &out, nil
}
