package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceVerdict string

const (
	SecuritySecurityGroupEvidenceVerdict_Malicious      SecuritySecurityGroupEvidenceVerdict = "malicious"
	SecuritySecurityGroupEvidenceVerdict_NoThreatsFound SecuritySecurityGroupEvidenceVerdict = "noThreatsFound"
	SecuritySecurityGroupEvidenceVerdict_Suspicious     SecuritySecurityGroupEvidenceVerdict = "suspicious"
	SecuritySecurityGroupEvidenceVerdict_Unknown        SecuritySecurityGroupEvidenceVerdict = "unknown"
)

func PossibleValuesForSecuritySecurityGroupEvidenceVerdict() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceVerdict_Malicious),
		string(SecuritySecurityGroupEvidenceVerdict_NoThreatsFound),
		string(SecuritySecurityGroupEvidenceVerdict_Suspicious),
		string(SecuritySecurityGroupEvidenceVerdict_Unknown),
	}
}

func (s *SecuritySecurityGroupEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySecurityGroupEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySecurityGroupEvidenceVerdict(input string) (*SecuritySecurityGroupEvidenceVerdict, error) {
	vals := map[string]SecuritySecurityGroupEvidenceVerdict{
		"malicious":      SecuritySecurityGroupEvidenceVerdict_Malicious,
		"nothreatsfound": SecuritySecurityGroupEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecuritySecurityGroupEvidenceVerdict_Suspicious,
		"unknown":        SecuritySecurityGroupEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceVerdict(input)
	return &out, nil
}
