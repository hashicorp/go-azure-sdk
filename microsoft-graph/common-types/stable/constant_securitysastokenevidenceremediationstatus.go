package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySasTokenEvidenceRemediationStatus string

const (
	SecuritySasTokenEvidenceRemediationStatus_Blocked    SecuritySasTokenEvidenceRemediationStatus = "blocked"
	SecuritySasTokenEvidenceRemediationStatus_None       SecuritySasTokenEvidenceRemediationStatus = "none"
	SecuritySasTokenEvidenceRemediationStatus_NotFound   SecuritySasTokenEvidenceRemediationStatus = "notFound"
	SecuritySasTokenEvidenceRemediationStatus_Prevented  SecuritySasTokenEvidenceRemediationStatus = "prevented"
	SecuritySasTokenEvidenceRemediationStatus_Remediated SecuritySasTokenEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecuritySasTokenEvidenceRemediationStatus() []string {
	return []string{
		string(SecuritySasTokenEvidenceRemediationStatus_Blocked),
		string(SecuritySasTokenEvidenceRemediationStatus_None),
		string(SecuritySasTokenEvidenceRemediationStatus_NotFound),
		string(SecuritySasTokenEvidenceRemediationStatus_Prevented),
		string(SecuritySasTokenEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecuritySasTokenEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySasTokenEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySasTokenEvidenceRemediationStatus(input string) (*SecuritySasTokenEvidenceRemediationStatus, error) {
	vals := map[string]SecuritySasTokenEvidenceRemediationStatus{
		"blocked":    SecuritySasTokenEvidenceRemediationStatus_Blocked,
		"none":       SecuritySasTokenEvidenceRemediationStatus_None,
		"notfound":   SecuritySasTokenEvidenceRemediationStatus_NotFound,
		"prevented":  SecuritySasTokenEvidenceRemediationStatus_Prevented,
		"remediated": SecuritySasTokenEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySasTokenEvidenceRemediationStatus(input)
	return &out, nil
}
