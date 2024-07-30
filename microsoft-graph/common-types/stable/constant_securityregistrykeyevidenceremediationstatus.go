package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceRemediationStatus string

const (
	SecurityRegistryKeyEvidenceRemediationStatus_Blocked    SecurityRegistryKeyEvidenceRemediationStatus = "blocked"
	SecurityRegistryKeyEvidenceRemediationStatus_None       SecurityRegistryKeyEvidenceRemediationStatus = "none"
	SecurityRegistryKeyEvidenceRemediationStatus_NotFound   SecurityRegistryKeyEvidenceRemediationStatus = "notFound"
	SecurityRegistryKeyEvidenceRemediationStatus_Prevented  SecurityRegistryKeyEvidenceRemediationStatus = "prevented"
	SecurityRegistryKeyEvidenceRemediationStatus_Remediated SecurityRegistryKeyEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityRegistryKeyEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceRemediationStatus_Blocked),
		string(SecurityRegistryKeyEvidenceRemediationStatus_None),
		string(SecurityRegistryKeyEvidenceRemediationStatus_NotFound),
		string(SecurityRegistryKeyEvidenceRemediationStatus_Prevented),
		string(SecurityRegistryKeyEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityRegistryKeyEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryKeyEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryKeyEvidenceRemediationStatus(input string) (*SecurityRegistryKeyEvidenceRemediationStatus, error) {
	vals := map[string]SecurityRegistryKeyEvidenceRemediationStatus{
		"blocked":    SecurityRegistryKeyEvidenceRemediationStatus_Blocked,
		"none":       SecurityRegistryKeyEvidenceRemediationStatus_None,
		"notfound":   SecurityRegistryKeyEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityRegistryKeyEvidenceRemediationStatus_Prevented,
		"remediated": SecurityRegistryKeyEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceRemediationStatus(input)
	return &out, nil
}
