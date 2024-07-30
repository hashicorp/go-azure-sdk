package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidenceRemediationStatus string

const (
	SecurityUserEvidenceRemediationStatus_Blocked    SecurityUserEvidenceRemediationStatus = "blocked"
	SecurityUserEvidenceRemediationStatus_None       SecurityUserEvidenceRemediationStatus = "none"
	SecurityUserEvidenceRemediationStatus_NotFound   SecurityUserEvidenceRemediationStatus = "notFound"
	SecurityUserEvidenceRemediationStatus_Prevented  SecurityUserEvidenceRemediationStatus = "prevented"
	SecurityUserEvidenceRemediationStatus_Remediated SecurityUserEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityUserEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityUserEvidenceRemediationStatus_Blocked),
		string(SecurityUserEvidenceRemediationStatus_None),
		string(SecurityUserEvidenceRemediationStatus_NotFound),
		string(SecurityUserEvidenceRemediationStatus_Prevented),
		string(SecurityUserEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityUserEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserEvidenceRemediationStatus(input string) (*SecurityUserEvidenceRemediationStatus, error) {
	vals := map[string]SecurityUserEvidenceRemediationStatus{
		"blocked":    SecurityUserEvidenceRemediationStatus_Blocked,
		"none":       SecurityUserEvidenceRemediationStatus_None,
		"notfound":   SecurityUserEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityUserEvidenceRemediationStatus_Prevented,
		"remediated": SecurityUserEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserEvidenceRemediationStatus(input)
	return &out, nil
}
