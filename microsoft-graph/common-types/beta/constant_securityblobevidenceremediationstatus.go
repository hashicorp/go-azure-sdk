package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidenceRemediationStatus string

const (
	SecurityBlobEvidenceRemediationStatus_Blocked    SecurityBlobEvidenceRemediationStatus = "blocked"
	SecurityBlobEvidenceRemediationStatus_None       SecurityBlobEvidenceRemediationStatus = "none"
	SecurityBlobEvidenceRemediationStatus_NotFound   SecurityBlobEvidenceRemediationStatus = "notFound"
	SecurityBlobEvidenceRemediationStatus_Prevented  SecurityBlobEvidenceRemediationStatus = "prevented"
	SecurityBlobEvidenceRemediationStatus_Remediated SecurityBlobEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityBlobEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityBlobEvidenceRemediationStatus_Blocked),
		string(SecurityBlobEvidenceRemediationStatus_None),
		string(SecurityBlobEvidenceRemediationStatus_NotFound),
		string(SecurityBlobEvidenceRemediationStatus_Prevented),
		string(SecurityBlobEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityBlobEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobEvidenceRemediationStatus(input string) (*SecurityBlobEvidenceRemediationStatus, error) {
	vals := map[string]SecurityBlobEvidenceRemediationStatus{
		"blocked":    SecurityBlobEvidenceRemediationStatus_Blocked,
		"none":       SecurityBlobEvidenceRemediationStatus_None,
		"notfound":   SecurityBlobEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityBlobEvidenceRemediationStatus_Prevented,
		"remediated": SecurityBlobEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobEvidenceRemediationStatus(input)
	return &out, nil
}
