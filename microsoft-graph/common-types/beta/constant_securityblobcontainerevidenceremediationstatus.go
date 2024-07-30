package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceRemediationStatus string

const (
	SecurityBlobContainerEvidenceRemediationStatus_Blocked    SecurityBlobContainerEvidenceRemediationStatus = "blocked"
	SecurityBlobContainerEvidenceRemediationStatus_None       SecurityBlobContainerEvidenceRemediationStatus = "none"
	SecurityBlobContainerEvidenceRemediationStatus_NotFound   SecurityBlobContainerEvidenceRemediationStatus = "notFound"
	SecurityBlobContainerEvidenceRemediationStatus_Prevented  SecurityBlobContainerEvidenceRemediationStatus = "prevented"
	SecurityBlobContainerEvidenceRemediationStatus_Remediated SecurityBlobContainerEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityBlobContainerEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityBlobContainerEvidenceRemediationStatus_Blocked),
		string(SecurityBlobContainerEvidenceRemediationStatus_None),
		string(SecurityBlobContainerEvidenceRemediationStatus_NotFound),
		string(SecurityBlobContainerEvidenceRemediationStatus_Prevented),
		string(SecurityBlobContainerEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityBlobContainerEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobContainerEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobContainerEvidenceRemediationStatus(input string) (*SecurityBlobContainerEvidenceRemediationStatus, error) {
	vals := map[string]SecurityBlobContainerEvidenceRemediationStatus{
		"blocked":    SecurityBlobContainerEvidenceRemediationStatus_Blocked,
		"none":       SecurityBlobContainerEvidenceRemediationStatus_None,
		"notfound":   SecurityBlobContainerEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityBlobContainerEvidenceRemediationStatus_Prevented,
		"remediated": SecurityBlobContainerEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceRemediationStatus(input)
	return &out, nil
}
