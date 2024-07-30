package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceRemediationStatus string

const (
	SecurityContainerImageEvidenceRemediationStatus_Blocked    SecurityContainerImageEvidenceRemediationStatus = "blocked"
	SecurityContainerImageEvidenceRemediationStatus_None       SecurityContainerImageEvidenceRemediationStatus = "none"
	SecurityContainerImageEvidenceRemediationStatus_NotFound   SecurityContainerImageEvidenceRemediationStatus = "notFound"
	SecurityContainerImageEvidenceRemediationStatus_Prevented  SecurityContainerImageEvidenceRemediationStatus = "prevented"
	SecurityContainerImageEvidenceRemediationStatus_Remediated SecurityContainerImageEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityContainerImageEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityContainerImageEvidenceRemediationStatus_Blocked),
		string(SecurityContainerImageEvidenceRemediationStatus_None),
		string(SecurityContainerImageEvidenceRemediationStatus_NotFound),
		string(SecurityContainerImageEvidenceRemediationStatus_Prevented),
		string(SecurityContainerImageEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityContainerImageEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerImageEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerImageEvidenceRemediationStatus(input string) (*SecurityContainerImageEvidenceRemediationStatus, error) {
	vals := map[string]SecurityContainerImageEvidenceRemediationStatus{
		"blocked":    SecurityContainerImageEvidenceRemediationStatus_Blocked,
		"none":       SecurityContainerImageEvidenceRemediationStatus_None,
		"notfound":   SecurityContainerImageEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityContainerImageEvidenceRemediationStatus_Prevented,
		"remediated": SecurityContainerImageEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceRemediationStatus(input)
	return &out, nil
}
