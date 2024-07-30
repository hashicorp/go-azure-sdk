package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostLogonSessionEvidenceVerdict string

const (
	SecurityHostLogonSessionEvidenceVerdict_Malicious      SecurityHostLogonSessionEvidenceVerdict = "malicious"
	SecurityHostLogonSessionEvidenceVerdict_NoThreatsFound SecurityHostLogonSessionEvidenceVerdict = "noThreatsFound"
	SecurityHostLogonSessionEvidenceVerdict_Suspicious     SecurityHostLogonSessionEvidenceVerdict = "suspicious"
	SecurityHostLogonSessionEvidenceVerdict_Unknown        SecurityHostLogonSessionEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityHostLogonSessionEvidenceVerdict() []string {
	return []string{
		string(SecurityHostLogonSessionEvidenceVerdict_Malicious),
		string(SecurityHostLogonSessionEvidenceVerdict_NoThreatsFound),
		string(SecurityHostLogonSessionEvidenceVerdict_Suspicious),
		string(SecurityHostLogonSessionEvidenceVerdict_Unknown),
	}
}

func (s *SecurityHostLogonSessionEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityHostLogonSessionEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityHostLogonSessionEvidenceVerdict(input string) (*SecurityHostLogonSessionEvidenceVerdict, error) {
	vals := map[string]SecurityHostLogonSessionEvidenceVerdict{
		"malicious":      SecurityHostLogonSessionEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityHostLogonSessionEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityHostLogonSessionEvidenceVerdict_Suspicious,
		"unknown":        SecurityHostLogonSessionEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityHostLogonSessionEvidenceVerdict(input)
	return &out, nil
}
