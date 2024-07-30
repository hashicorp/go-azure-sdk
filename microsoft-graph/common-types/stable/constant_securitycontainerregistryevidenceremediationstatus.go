package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidenceRemediationStatus string

const (
	SecurityContainerRegistryEvidenceRemediationStatus_Blocked    SecurityContainerRegistryEvidenceRemediationStatus = "blocked"
	SecurityContainerRegistryEvidenceRemediationStatus_None       SecurityContainerRegistryEvidenceRemediationStatus = "none"
	SecurityContainerRegistryEvidenceRemediationStatus_NotFound   SecurityContainerRegistryEvidenceRemediationStatus = "notFound"
	SecurityContainerRegistryEvidenceRemediationStatus_Prevented  SecurityContainerRegistryEvidenceRemediationStatus = "prevented"
	SecurityContainerRegistryEvidenceRemediationStatus_Remediated SecurityContainerRegistryEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityContainerRegistryEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityContainerRegistryEvidenceRemediationStatus_Blocked),
		string(SecurityContainerRegistryEvidenceRemediationStatus_None),
		string(SecurityContainerRegistryEvidenceRemediationStatus_NotFound),
		string(SecurityContainerRegistryEvidenceRemediationStatus_Prevented),
		string(SecurityContainerRegistryEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityContainerRegistryEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerRegistryEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerRegistryEvidenceRemediationStatus(input string) (*SecurityContainerRegistryEvidenceRemediationStatus, error) {
	vals := map[string]SecurityContainerRegistryEvidenceRemediationStatus{
		"blocked":    SecurityContainerRegistryEvidenceRemediationStatus_Blocked,
		"none":       SecurityContainerRegistryEvidenceRemediationStatus_None,
		"notfound":   SecurityContainerRegistryEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityContainerRegistryEvidenceRemediationStatus_Prevented,
		"remediated": SecurityContainerRegistryEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerRegistryEvidenceRemediationStatus(input)
	return &out, nil
}
