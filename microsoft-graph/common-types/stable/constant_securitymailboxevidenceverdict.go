package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceVerdict string

const (
	SecurityMailboxEvidenceVerdict_Malicious      SecurityMailboxEvidenceVerdict = "malicious"
	SecurityMailboxEvidenceVerdict_NoThreatsFound SecurityMailboxEvidenceVerdict = "noThreatsFound"
	SecurityMailboxEvidenceVerdict_Suspicious     SecurityMailboxEvidenceVerdict = "suspicious"
	SecurityMailboxEvidenceVerdict_Unknown        SecurityMailboxEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityMailboxEvidenceVerdict() []string {
	return []string{
		string(SecurityMailboxEvidenceVerdict_Malicious),
		string(SecurityMailboxEvidenceVerdict_NoThreatsFound),
		string(SecurityMailboxEvidenceVerdict_Suspicious),
		string(SecurityMailboxEvidenceVerdict_Unknown),
	}
}

func (s *SecurityMailboxEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailboxEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailboxEvidenceVerdict(input string) (*SecurityMailboxEvidenceVerdict, error) {
	vals := map[string]SecurityMailboxEvidenceVerdict{
		"malicious":      SecurityMailboxEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityMailboxEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityMailboxEvidenceVerdict_Suspicious,
		"unknown":        SecurityMailboxEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxEvidenceVerdict(input)
	return &out, nil
}
