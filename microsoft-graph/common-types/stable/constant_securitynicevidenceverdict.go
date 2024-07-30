package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityNicEvidenceVerdict string

const (
	SecurityNicEvidenceVerdict_Malicious      SecurityNicEvidenceVerdict = "malicious"
	SecurityNicEvidenceVerdict_NoThreatsFound SecurityNicEvidenceVerdict = "noThreatsFound"
	SecurityNicEvidenceVerdict_Suspicious     SecurityNicEvidenceVerdict = "suspicious"
	SecurityNicEvidenceVerdict_Unknown        SecurityNicEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityNicEvidenceVerdict() []string {
	return []string{
		string(SecurityNicEvidenceVerdict_Malicious),
		string(SecurityNicEvidenceVerdict_NoThreatsFound),
		string(SecurityNicEvidenceVerdict_Suspicious),
		string(SecurityNicEvidenceVerdict_Unknown),
	}
}

func (s *SecurityNicEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityNicEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityNicEvidenceVerdict(input string) (*SecurityNicEvidenceVerdict, error) {
	vals := map[string]SecurityNicEvidenceVerdict{
		"malicious":      SecurityNicEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityNicEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityNicEvidenceVerdict_Suspicious,
		"unknown":        SecurityNicEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityNicEvidenceVerdict(input)
	return &out, nil
}
