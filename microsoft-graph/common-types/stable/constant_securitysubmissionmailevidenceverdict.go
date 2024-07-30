package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionMailEvidenceVerdict string

const (
	SecuritySubmissionMailEvidenceVerdict_Malicious      SecuritySubmissionMailEvidenceVerdict = "malicious"
	SecuritySubmissionMailEvidenceVerdict_NoThreatsFound SecuritySubmissionMailEvidenceVerdict = "noThreatsFound"
	SecuritySubmissionMailEvidenceVerdict_Suspicious     SecuritySubmissionMailEvidenceVerdict = "suspicious"
	SecuritySubmissionMailEvidenceVerdict_Unknown        SecuritySubmissionMailEvidenceVerdict = "unknown"
)

func PossibleValuesForSecuritySubmissionMailEvidenceVerdict() []string {
	return []string{
		string(SecuritySubmissionMailEvidenceVerdict_Malicious),
		string(SecuritySubmissionMailEvidenceVerdict_NoThreatsFound),
		string(SecuritySubmissionMailEvidenceVerdict_Suspicious),
		string(SecuritySubmissionMailEvidenceVerdict_Unknown),
	}
}

func (s *SecuritySubmissionMailEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionMailEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionMailEvidenceVerdict(input string) (*SecuritySubmissionMailEvidenceVerdict, error) {
	vals := map[string]SecuritySubmissionMailEvidenceVerdict{
		"malicious":      SecuritySubmissionMailEvidenceVerdict_Malicious,
		"nothreatsfound": SecuritySubmissionMailEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecuritySubmissionMailEvidenceVerdict_Suspicious,
		"unknown":        SecuritySubmissionMailEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionMailEvidenceVerdict(input)
	return &out, nil
}
