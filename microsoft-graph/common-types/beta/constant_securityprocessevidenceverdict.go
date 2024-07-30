package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceVerdict string

const (
	SecurityProcessEvidenceVerdict_Malicious      SecurityProcessEvidenceVerdict = "malicious"
	SecurityProcessEvidenceVerdict_NoThreatsFound SecurityProcessEvidenceVerdict = "noThreatsFound"
	SecurityProcessEvidenceVerdict_Suspicious     SecurityProcessEvidenceVerdict = "suspicious"
	SecurityProcessEvidenceVerdict_Unknown        SecurityProcessEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityProcessEvidenceVerdict() []string {
	return []string{
		string(SecurityProcessEvidenceVerdict_Malicious),
		string(SecurityProcessEvidenceVerdict_NoThreatsFound),
		string(SecurityProcessEvidenceVerdict_Suspicious),
		string(SecurityProcessEvidenceVerdict_Unknown),
	}
}

func (s *SecurityProcessEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityProcessEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityProcessEvidenceVerdict(input string) (*SecurityProcessEvidenceVerdict, error) {
	vals := map[string]SecurityProcessEvidenceVerdict{
		"malicious":      SecurityProcessEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityProcessEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityProcessEvidenceVerdict_Suspicious,
		"unknown":        SecurityProcessEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceVerdict(input)
	return &out, nil
}
