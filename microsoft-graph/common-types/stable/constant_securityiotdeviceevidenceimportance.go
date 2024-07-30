package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceEvidenceImportance string

const (
	SecurityIoTDeviceEvidenceImportance_High    SecurityIoTDeviceEvidenceImportance = "high"
	SecurityIoTDeviceEvidenceImportance_Low     SecurityIoTDeviceEvidenceImportance = "low"
	SecurityIoTDeviceEvidenceImportance_Normal  SecurityIoTDeviceEvidenceImportance = "normal"
	SecurityIoTDeviceEvidenceImportance_Unknown SecurityIoTDeviceEvidenceImportance = "unknown"
)

func PossibleValuesForSecurityIoTDeviceEvidenceImportance() []string {
	return []string{
		string(SecurityIoTDeviceEvidenceImportance_High),
		string(SecurityIoTDeviceEvidenceImportance_Low),
		string(SecurityIoTDeviceEvidenceImportance_Normal),
		string(SecurityIoTDeviceEvidenceImportance_Unknown),
	}
}

func (s *SecurityIoTDeviceEvidenceImportance) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIoTDeviceEvidenceImportance(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIoTDeviceEvidenceImportance(input string) (*SecurityIoTDeviceEvidenceImportance, error) {
	vals := map[string]SecurityIoTDeviceEvidenceImportance{
		"high":    SecurityIoTDeviceEvidenceImportance_High,
		"low":     SecurityIoTDeviceEvidenceImportance_Low,
		"normal":  SecurityIoTDeviceEvidenceImportance_Normal,
		"unknown": SecurityIoTDeviceEvidenceImportance_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIoTDeviceEvidenceImportance(input)
	return &out, nil
}
