package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidenceVerdict string

const (
	SecurityBlobEvidenceVerdict_Malicious      SecurityBlobEvidenceVerdict = "malicious"
	SecurityBlobEvidenceVerdict_NoThreatsFound SecurityBlobEvidenceVerdict = "noThreatsFound"
	SecurityBlobEvidenceVerdict_Suspicious     SecurityBlobEvidenceVerdict = "suspicious"
	SecurityBlobEvidenceVerdict_Unknown        SecurityBlobEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityBlobEvidenceVerdict() []string {
	return []string{
		string(SecurityBlobEvidenceVerdict_Malicious),
		string(SecurityBlobEvidenceVerdict_NoThreatsFound),
		string(SecurityBlobEvidenceVerdict_Suspicious),
		string(SecurityBlobEvidenceVerdict_Unknown),
	}
}

func (s *SecurityBlobEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobEvidenceVerdict(input string) (*SecurityBlobEvidenceVerdict, error) {
	vals := map[string]SecurityBlobEvidenceVerdict{
		"malicious":      SecurityBlobEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityBlobEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityBlobEvidenceVerdict_Suspicious,
		"unknown":        SecurityBlobEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobEvidenceVerdict(input)
	return &out, nil
}
