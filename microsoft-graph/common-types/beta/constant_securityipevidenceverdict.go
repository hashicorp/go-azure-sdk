package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceVerdict string

const (
	SecurityIpEvidenceVerdict_Malicious      SecurityIpEvidenceVerdict = "malicious"
	SecurityIpEvidenceVerdict_NoThreatsFound SecurityIpEvidenceVerdict = "noThreatsFound"
	SecurityIpEvidenceVerdict_Suspicious     SecurityIpEvidenceVerdict = "suspicious"
	SecurityIpEvidenceVerdict_Unknown        SecurityIpEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityIpEvidenceVerdict() []string {
	return []string{
		string(SecurityIpEvidenceVerdict_Malicious),
		string(SecurityIpEvidenceVerdict_NoThreatsFound),
		string(SecurityIpEvidenceVerdict_Suspicious),
		string(SecurityIpEvidenceVerdict_Unknown),
	}
}

func (s *SecurityIpEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIpEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIpEvidenceVerdict(input string) (*SecurityIpEvidenceVerdict, error) {
	vals := map[string]SecurityIpEvidenceVerdict{
		"malicious":      SecurityIpEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityIpEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityIpEvidenceVerdict_Suspicious,
		"unknown":        SecurityIpEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceVerdict(input)
	return &out, nil
}
