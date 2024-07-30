package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceRemediationStatus string

const (
	SecurityAmazonResourceEvidenceRemediationStatus_Blocked    SecurityAmazonResourceEvidenceRemediationStatus = "blocked"
	SecurityAmazonResourceEvidenceRemediationStatus_None       SecurityAmazonResourceEvidenceRemediationStatus = "none"
	SecurityAmazonResourceEvidenceRemediationStatus_NotFound   SecurityAmazonResourceEvidenceRemediationStatus = "notFound"
	SecurityAmazonResourceEvidenceRemediationStatus_Prevented  SecurityAmazonResourceEvidenceRemediationStatus = "prevented"
	SecurityAmazonResourceEvidenceRemediationStatus_Remediated SecurityAmazonResourceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityAmazonResourceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceRemediationStatus_Blocked),
		string(SecurityAmazonResourceEvidenceRemediationStatus_None),
		string(SecurityAmazonResourceEvidenceRemediationStatus_NotFound),
		string(SecurityAmazonResourceEvidenceRemediationStatus_Prevented),
		string(SecurityAmazonResourceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityAmazonResourceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAmazonResourceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAmazonResourceEvidenceRemediationStatus(input string) (*SecurityAmazonResourceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityAmazonResourceEvidenceRemediationStatus{
		"blocked":    SecurityAmazonResourceEvidenceRemediationStatus_Blocked,
		"none":       SecurityAmazonResourceEvidenceRemediationStatus_None,
		"notfound":   SecurityAmazonResourceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityAmazonResourceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityAmazonResourceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceRemediationStatus(input)
	return &out, nil
}
