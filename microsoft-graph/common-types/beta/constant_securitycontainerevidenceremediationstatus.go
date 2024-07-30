package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceRemediationStatus string

const (
	SecurityContainerEvidenceRemediationStatus_Blocked    SecurityContainerEvidenceRemediationStatus = "blocked"
	SecurityContainerEvidenceRemediationStatus_None       SecurityContainerEvidenceRemediationStatus = "none"
	SecurityContainerEvidenceRemediationStatus_NotFound   SecurityContainerEvidenceRemediationStatus = "notFound"
	SecurityContainerEvidenceRemediationStatus_Prevented  SecurityContainerEvidenceRemediationStatus = "prevented"
	SecurityContainerEvidenceRemediationStatus_Remediated SecurityContainerEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityContainerEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityContainerEvidenceRemediationStatus_Blocked),
		string(SecurityContainerEvidenceRemediationStatus_None),
		string(SecurityContainerEvidenceRemediationStatus_NotFound),
		string(SecurityContainerEvidenceRemediationStatus_Prevented),
		string(SecurityContainerEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityContainerEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerEvidenceRemediationStatus(input string) (*SecurityContainerEvidenceRemediationStatus, error) {
	vals := map[string]SecurityContainerEvidenceRemediationStatus{
		"blocked":    SecurityContainerEvidenceRemediationStatus_Blocked,
		"none":       SecurityContainerEvidenceRemediationStatus_None,
		"notfound":   SecurityContainerEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityContainerEvidenceRemediationStatus_Prevented,
		"remediated": SecurityContainerEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerEvidenceRemediationStatus(input)
	return &out, nil
}
