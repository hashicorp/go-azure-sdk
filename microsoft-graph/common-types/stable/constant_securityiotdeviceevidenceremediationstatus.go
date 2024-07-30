package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceEvidenceRemediationStatus string

const (
	SecurityIoTDeviceEvidenceRemediationStatus_Blocked    SecurityIoTDeviceEvidenceRemediationStatus = "blocked"
	SecurityIoTDeviceEvidenceRemediationStatus_None       SecurityIoTDeviceEvidenceRemediationStatus = "none"
	SecurityIoTDeviceEvidenceRemediationStatus_NotFound   SecurityIoTDeviceEvidenceRemediationStatus = "notFound"
	SecurityIoTDeviceEvidenceRemediationStatus_Prevented  SecurityIoTDeviceEvidenceRemediationStatus = "prevented"
	SecurityIoTDeviceEvidenceRemediationStatus_Remediated SecurityIoTDeviceEvidenceRemediationStatus = "remediated"
)

func PossibleValuesForSecurityIoTDeviceEvidenceRemediationStatus() []string {
	return []string{
		string(SecurityIoTDeviceEvidenceRemediationStatus_Blocked),
		string(SecurityIoTDeviceEvidenceRemediationStatus_None),
		string(SecurityIoTDeviceEvidenceRemediationStatus_NotFound),
		string(SecurityIoTDeviceEvidenceRemediationStatus_Prevented),
		string(SecurityIoTDeviceEvidenceRemediationStatus_Remediated),
	}
}

func (s *SecurityIoTDeviceEvidenceRemediationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIoTDeviceEvidenceRemediationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIoTDeviceEvidenceRemediationStatus(input string) (*SecurityIoTDeviceEvidenceRemediationStatus, error) {
	vals := map[string]SecurityIoTDeviceEvidenceRemediationStatus{
		"blocked":    SecurityIoTDeviceEvidenceRemediationStatus_Blocked,
		"none":       SecurityIoTDeviceEvidenceRemediationStatus_None,
		"notfound":   SecurityIoTDeviceEvidenceRemediationStatus_NotFound,
		"prevented":  SecurityIoTDeviceEvidenceRemediationStatus_Prevented,
		"remediated": SecurityIoTDeviceEvidenceRemediationStatus_Remediated,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIoTDeviceEvidenceRemediationStatus(input)
	return &out, nil
}
