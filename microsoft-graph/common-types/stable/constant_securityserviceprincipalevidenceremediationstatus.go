package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalEvidenceRemediationStatus string

const (
	SecurityServicePrincipalEvidenceRemediationStatus_Blocked    SecurityServicePrincipalEvidenceRemediationStatus = "blocked"
	SecurityServicePrincipalEvidenceRemediationStatus_None       SecurityServicePrincipalEvidenceRemediationStatus = "none"
	SecurityServicePrincipalEvidenceRemediationStatus_NotFound   SecurityServicePrincipalEvidenceRemediationStatus = "notFound"
	SecurityServicePrincipalEvidenceRemediationStatus_Prevented  SecurityServicePrincipalEvidenceRemediationStatus = "prevented"
	SecurityServicePrincipalEvidenceRemediationStatus_Remediated SecurityServicePrincipalEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityServicePrincipalEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityServicePrincipalEvidenceRemediationStatus_Blocked),
		string(SecurityServicePrincipalEvidenceRemediationStatus_None),
		string(SecurityServicePrincipalEvidenceRemediationStatus_NotFound),
		string(SecurityServicePrincipalEvidenceRemediationStatus_Prevented),
		string(SecurityServicePrincipalEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityServicePrincipalEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServicePrincipalEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServicePrincipalEvidenceRemediationStatus(input string) (*SecurityServicePrincipalEvidenceRemediationStatus, error) {
	vals := map[string]SecurityServicePrincipalEvidenceRemediationStatus{
		"blocked":    SecurityServicePrincipalEvidenceRemediationStatus_Blocked,
		"none":       SecurityServicePrincipalEvidenceRemediationStatus_None,
		"notfound":   SecurityServicePrincipalEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityServicePrincipalEvidenceRemediationStatus_Prevented,
		"remediated": SecurityServicePrincipalEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServicePrincipalEvidenceRemediationStatus(input)
	return &out, nil
}
