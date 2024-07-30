package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNetworkConnectionEvidenceRemediationStatus string

const (
	SecurityNetworkConnectionEvidenceRemediationStatus_Blocked    SecurityNetworkConnectionEvidenceRemediationStatus = "blocked"
	SecurityNetworkConnectionEvidenceRemediationStatus_None       SecurityNetworkConnectionEvidenceRemediationStatus = "none"
	SecurityNetworkConnectionEvidenceRemediationStatus_NotFound   SecurityNetworkConnectionEvidenceRemediationStatus = "notFound"
	SecurityNetworkConnectionEvidenceRemediationStatus_Prevented  SecurityNetworkConnectionEvidenceRemediationStatus = "prevented"
	SecurityNetworkConnectionEvidenceRemediationStatus_Remediated SecurityNetworkConnectionEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityNetworkConnectionEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityNetworkConnectionEvidenceRemediationStatus_Blocked),
		string(SecurityNetworkConnectionEvidenceRemediationStatus_None),
		string(SecurityNetworkConnectionEvidenceRemediationStatus_NotFound),
		string(SecurityNetworkConnectionEvidenceRemediationStatus_Prevented),
		string(SecurityNetworkConnectionEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityNetworkConnectionEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNetworkConnectionEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNetworkConnectionEvidenceRemediationStatus(input string) (*SecurityNetworkConnectionEvidenceRemediationStatus, error) {
	vals := map[string]SecurityNetworkConnectionEvidenceRemediationStatus{
		"blocked":    SecurityNetworkConnectionEvidenceRemediationStatus_Blocked,
		"none":       SecurityNetworkConnectionEvidenceRemediationStatus_None,
		"notfound":   SecurityNetworkConnectionEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityNetworkConnectionEvidenceRemediationStatus_Prevented,
		"remediated": SecurityNetworkConnectionEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNetworkConnectionEvidenceRemediationStatus(input)
	return &out, nil
}
