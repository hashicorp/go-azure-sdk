package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceRemediationStatus string

const (
	SecurityUrlEvidenceRemediationStatus_Blocked    SecurityUrlEvidenceRemediationStatus = "blocked"
	SecurityUrlEvidenceRemediationStatus_None       SecurityUrlEvidenceRemediationStatus = "none"
	SecurityUrlEvidenceRemediationStatus_NotFound   SecurityUrlEvidenceRemediationStatus = "notFound"
	SecurityUrlEvidenceRemediationStatus_Prevented  SecurityUrlEvidenceRemediationStatus = "prevented"
	SecurityUrlEvidenceRemediationStatus_Remediated SecurityUrlEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityUrlEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityUrlEvidenceRemediationStatus_Blocked),
		string(SecurityUrlEvidenceRemediationStatus_None),
		string(SecurityUrlEvidenceRemediationStatus_NotFound),
		string(SecurityUrlEvidenceRemediationStatus_Prevented),
		string(SecurityUrlEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityUrlEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlEvidenceRemediationStatus(input string) (*SecurityUrlEvidenceRemediationStatus, error) {
	vals := map[string]SecurityUrlEvidenceRemediationStatus{
		"blocked":    SecurityUrlEvidenceRemediationStatus_Blocked,
		"none":       SecurityUrlEvidenceRemediationStatus_None,
		"notfound":   SecurityUrlEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityUrlEvidenceRemediationStatus_Prevented,
		"remediated": SecurityUrlEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlEvidenceRemediationStatus(input)
	return &out, nil
}
