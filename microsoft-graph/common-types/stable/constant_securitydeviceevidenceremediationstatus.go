package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRemediationStatus string

const (
	SecurityDeviceEvidenceRemediationStatus_Blocked    SecurityDeviceEvidenceRemediationStatus = "blocked"
	SecurityDeviceEvidenceRemediationStatus_None       SecurityDeviceEvidenceRemediationStatus = "none"
	SecurityDeviceEvidenceRemediationStatus_NotFound   SecurityDeviceEvidenceRemediationStatus = "notFound"
	SecurityDeviceEvidenceRemediationStatus_Prevented  SecurityDeviceEvidenceRemediationStatus = "prevented"
	SecurityDeviceEvidenceRemediationStatus_Remediated SecurityDeviceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityDeviceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityDeviceEvidenceRemediationStatus_Blocked),
		string(SecurityDeviceEvidenceRemediationStatus_None),
		string(SecurityDeviceEvidenceRemediationStatus_NotFound),
		string(SecurityDeviceEvidenceRemediationStatus_Prevented),
		string(SecurityDeviceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityDeviceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceRemediationStatus(input string) (*SecurityDeviceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityDeviceEvidenceRemediationStatus{
		"blocked":    SecurityDeviceEvidenceRemediationStatus_Blocked,
		"none":       SecurityDeviceEvidenceRemediationStatus_None,
		"notfound":   SecurityDeviceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityDeviceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityDeviceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRemediationStatus(input)
	return &out, nil
}
