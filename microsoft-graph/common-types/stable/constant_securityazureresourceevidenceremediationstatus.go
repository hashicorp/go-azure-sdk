package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceRemediationStatus string

const (
	SecurityAzureResourceEvidenceRemediationStatus_Blocked    SecurityAzureResourceEvidenceRemediationStatus = "blocked"
	SecurityAzureResourceEvidenceRemediationStatus_None       SecurityAzureResourceEvidenceRemediationStatus = "none"
	SecurityAzureResourceEvidenceRemediationStatus_NotFound   SecurityAzureResourceEvidenceRemediationStatus = "notFound"
	SecurityAzureResourceEvidenceRemediationStatus_Prevented  SecurityAzureResourceEvidenceRemediationStatus = "prevented"
	SecurityAzureResourceEvidenceRemediationStatus_Remediated SecurityAzureResourceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityAzureResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAzureResourceEvidenceRemediationStatus_Blocked),
		string(SecurityAzureResourceEvidenceRemediationStatus_None),
		string(SecurityAzureResourceEvidenceRemediationStatus_NotFound),
		string(SecurityAzureResourceEvidenceRemediationStatus_Prevented),
		string(SecurityAzureResourceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityAzureResourceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAzureResourceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAzureResourceEvidenceRemediationStatus(input string) (*SecurityAzureResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAzureResourceEvidenceRemediationStatus{
		"blocked":    SecurityAzureResourceEvidenceRemediationStatus_Blocked,
		"none":       SecurityAzureResourceEvidenceRemediationStatus_None,
		"notfound":   SecurityAzureResourceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityAzureResourceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityAzureResourceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceRemediationStatus(input)
	return &out, nil
}
