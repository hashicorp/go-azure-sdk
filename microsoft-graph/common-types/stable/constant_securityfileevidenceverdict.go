package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceVerdict string

const (
	SecurityFileEvidenceVerdict_Malicious      SecurityFileEvidenceVerdict = "malicious"
	SecurityFileEvidenceVerdict_NoThreatsFound SecurityFileEvidenceVerdict = "noThreatsFound"
	SecurityFileEvidenceVerdict_Suspicious     SecurityFileEvidenceVerdict = "suspicious"
	SecurityFileEvidenceVerdict_Unknown        SecurityFileEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityFileEvidenceVerdict() []string {
	return []string{
		string(SecurityFileEvidenceVerdict_Malicious),
		string(SecurityFileEvidenceVerdict_NoThreatsFound),
		string(SecurityFileEvidenceVerdict_Suspicious),
		string(SecurityFileEvidenceVerdict_Unknown),
	}
}

func (s *SecurityFileEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileEvidenceVerdict(input string) (*SecurityFileEvidenceVerdict, error) {
	vals := map[string]SecurityFileEvidenceVerdict{
		"malicious":      SecurityFileEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityFileEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityFileEvidenceVerdict_Suspicious,
		"unknown":        SecurityFileEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceVerdict(input)
	return &out, nil
}
