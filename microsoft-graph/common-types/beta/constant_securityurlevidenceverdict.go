package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceVerdict string

const (
	SecurityUrlEvidenceVerdict_Malicious      SecurityUrlEvidenceVerdict = "malicious"
	SecurityUrlEvidenceVerdict_NoThreatsFound SecurityUrlEvidenceVerdict = "noThreatsFound"
	SecurityUrlEvidenceVerdict_Suspicious     SecurityUrlEvidenceVerdict = "suspicious"
	SecurityUrlEvidenceVerdict_Unknown        SecurityUrlEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityUrlEvidenceVerdict() []string {
	return []string{
		string(SecurityUrlEvidenceVerdict_Malicious),
		string(SecurityUrlEvidenceVerdict_NoThreatsFound),
		string(SecurityUrlEvidenceVerdict_Suspicious),
		string(SecurityUrlEvidenceVerdict_Unknown),
	}
}

func (s *SecurityUrlEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlEvidenceVerdict(input string) (*SecurityUrlEvidenceVerdict, error) {
	vals := map[string]SecurityUrlEvidenceVerdict{
		"malicious":      SecurityUrlEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityUrlEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityUrlEvidenceVerdict_Suspicious,
		"unknown":        SecurityUrlEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlEvidenceVerdict(input)
	return &out, nil
}
