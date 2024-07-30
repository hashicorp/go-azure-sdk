package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceRemediationStatus string

const (
	SecurityRegistryValueEvidenceRemediationStatus_Blocked    SecurityRegistryValueEvidenceRemediationStatus = "blocked"
	SecurityRegistryValueEvidenceRemediationStatus_None       SecurityRegistryValueEvidenceRemediationStatus = "none"
	SecurityRegistryValueEvidenceRemediationStatus_NotFound   SecurityRegistryValueEvidenceRemediationStatus = "notFound"
	SecurityRegistryValueEvidenceRemediationStatus_Prevented  SecurityRegistryValueEvidenceRemediationStatus = "prevented"
	SecurityRegistryValueEvidenceRemediationStatus_Remediated SecurityRegistryValueEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityRegistryValueEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityRegistryValueEvidenceRemediationStatus_Blocked),
		string(SecurityRegistryValueEvidenceRemediationStatus_None),
		string(SecurityRegistryValueEvidenceRemediationStatus_NotFound),
		string(SecurityRegistryValueEvidenceRemediationStatus_Prevented),
		string(SecurityRegistryValueEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityRegistryValueEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityRegistryValueEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityRegistryValueEvidenceRemediationStatus(input string) (*SecurityRegistryValueEvidenceRemediationStatus, error) {
	vals := map[string]SecurityRegistryValueEvidenceRemediationStatus{
		"blocked":    SecurityRegistryValueEvidenceRemediationStatus_Blocked,
		"none":       SecurityRegistryValueEvidenceRemediationStatus_None,
		"notfound":   SecurityRegistryValueEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityRegistryValueEvidenceRemediationStatus_Prevented,
		"remediated": SecurityRegistryValueEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceRemediationStatus(input)
	return &out, nil
}
