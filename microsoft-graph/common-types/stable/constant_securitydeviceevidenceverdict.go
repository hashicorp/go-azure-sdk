package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceVerdict string

const (
	SecurityDeviceEvidenceVerdict_Malicious      SecurityDeviceEvidenceVerdict = "malicious"
	SecurityDeviceEvidenceVerdict_NoThreatsFound SecurityDeviceEvidenceVerdict = "noThreatsFound"
	SecurityDeviceEvidenceVerdict_Suspicious     SecurityDeviceEvidenceVerdict = "suspicious"
	SecurityDeviceEvidenceVerdict_Unknown        SecurityDeviceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityDeviceEvidenceVerdict() []string {
	return []string{
		string(SecurityDeviceEvidenceVerdict_Malicious),
		string(SecurityDeviceEvidenceVerdict_NoThreatsFound),
		string(SecurityDeviceEvidenceVerdict_Suspicious),
		string(SecurityDeviceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityDeviceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceVerdict(input string) (*SecurityDeviceEvidenceVerdict, error) {
	vals := map[string]SecurityDeviceEvidenceVerdict{
		"malicious":      SecurityDeviceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityDeviceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityDeviceEvidenceVerdict_Suspicious,
		"unknown":        SecurityDeviceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceVerdict(input)
	return &out, nil
}
