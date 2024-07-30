package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidenceRemediationStatus string

const (
	SecurityKubernetesPodEvidenceRemediationStatus_Blocked    SecurityKubernetesPodEvidenceRemediationStatus = "blocked"
	SecurityKubernetesPodEvidenceRemediationStatus_None       SecurityKubernetesPodEvidenceRemediationStatus = "none"
	SecurityKubernetesPodEvidenceRemediationStatus_NotFound   SecurityKubernetesPodEvidenceRemediationStatus = "notFound"
	SecurityKubernetesPodEvidenceRemediationStatus_Prevented  SecurityKubernetesPodEvidenceRemediationStatus = "prevented"
	SecurityKubernetesPodEvidenceRemediationStatus_Remediated SecurityKubernetesPodEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityKubernetesPodEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityKubernetesPodEvidenceRemediationStatus_Blocked),
		string(SecurityKubernetesPodEvidenceRemediationStatus_None),
		string(SecurityKubernetesPodEvidenceRemediationStatus_NotFound),
		string(SecurityKubernetesPodEvidenceRemediationStatus_Prevented),
		string(SecurityKubernetesPodEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityKubernetesPodEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesPodEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesPodEvidenceRemediationStatus(input string) (*SecurityKubernetesPodEvidenceRemediationStatus, error) {
	vals := map[string]SecurityKubernetesPodEvidenceRemediationStatus{
		"blocked":    SecurityKubernetesPodEvidenceRemediationStatus_Blocked,
		"none":       SecurityKubernetesPodEvidenceRemediationStatus_None,
		"notfound":   SecurityKubernetesPodEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityKubernetesPodEvidenceRemediationStatus_Prevented,
		"remediated": SecurityKubernetesPodEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPodEvidenceRemediationStatus(input)
	return &out, nil
}
