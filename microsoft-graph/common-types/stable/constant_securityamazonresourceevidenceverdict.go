package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceVerdict string

const (
	SecurityAmazonResourceEvidenceVerdict_Malicious      SecurityAmazonResourceEvidenceVerdict = "malicious"
	SecurityAmazonResourceEvidenceVerdict_NoThreatsFound SecurityAmazonResourceEvidenceVerdict = "noThreatsFound"
	SecurityAmazonResourceEvidenceVerdict_Suspicious     SecurityAmazonResourceEvidenceVerdict = "suspicious"
	SecurityAmazonResourceEvidenceVerdict_Unknown        SecurityAmazonResourceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityAmazonResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceVerdict_Malicious),
		string(SecurityAmazonResourceEvidenceVerdict_NoThreatsFound),
		string(SecurityAmazonResourceEvidenceVerdict_Suspicious),
		string(SecurityAmazonResourceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityAmazonResourceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAmazonResourceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAmazonResourceEvidenceVerdict(input string) (*SecurityAmazonResourceEvidenceVerdict, error) {
	vals := map[string]SecurityAmazonResourceEvidenceVerdict{
		"malicious":      SecurityAmazonResourceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityAmazonResourceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityAmazonResourceEvidenceVerdict_Suspicious,
		"unknown":        SecurityAmazonResourceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceVerdict(input)
	return &out, nil
}
