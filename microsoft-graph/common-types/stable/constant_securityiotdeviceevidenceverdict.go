package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceEvidenceVerdict string

const (
	SecurityIoTDeviceEvidenceVerdict_Malicious      SecurityIoTDeviceEvidenceVerdict = "malicious"
	SecurityIoTDeviceEvidenceVerdict_NoThreatsFound SecurityIoTDeviceEvidenceVerdict = "noThreatsFound"
	SecurityIoTDeviceEvidenceVerdict_Suspicious     SecurityIoTDeviceEvidenceVerdict = "suspicious"
	SecurityIoTDeviceEvidenceVerdict_Unknown        SecurityIoTDeviceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityIoTDeviceEvidenceVerdict() []string {
	return []string{
		string(SecurityIoTDeviceEvidenceVerdict_Malicious),
		string(SecurityIoTDeviceEvidenceVerdict_NoThreatsFound),
		string(SecurityIoTDeviceEvidenceVerdict_Suspicious),
		string(SecurityIoTDeviceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityIoTDeviceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIoTDeviceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIoTDeviceEvidenceVerdict(input string) (*SecurityIoTDeviceEvidenceVerdict, error) {
	vals := map[string]SecurityIoTDeviceEvidenceVerdict{
		"malicious":      SecurityIoTDeviceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityIoTDeviceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityIoTDeviceEvidenceVerdict_Suspicious,
		"unknown":        SecurityIoTDeviceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIoTDeviceEvidenceVerdict(input)
	return &out, nil
}
