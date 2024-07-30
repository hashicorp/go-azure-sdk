package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceVerdict string

const (
	SecurityAnalyzedMessageEvidenceVerdict_Malicious      SecurityAnalyzedMessageEvidenceVerdict = "malicious"
	SecurityAnalyzedMessageEvidenceVerdict_NoThreatsFound SecurityAnalyzedMessageEvidenceVerdict = "noThreatsFound"
	SecurityAnalyzedMessageEvidenceVerdict_Suspicious     SecurityAnalyzedMessageEvidenceVerdict = "suspicious"
	SecurityAnalyzedMessageEvidenceVerdict_Unknown        SecurityAnalyzedMessageEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceVerdict() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceVerdict_Malicious),
		string(SecurityAnalyzedMessageEvidenceVerdict_NoThreatsFound),
		string(SecurityAnalyzedMessageEvidenceVerdict_Suspicious),
		string(SecurityAnalyzedMessageEvidenceVerdict_Unknown),
	}
}

func (s *SecurityAnalyzedMessageEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAnalyzedMessageEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAnalyzedMessageEvidenceVerdict(input string) (*SecurityAnalyzedMessageEvidenceVerdict, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceVerdict{
		"malicious":      SecurityAnalyzedMessageEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityAnalyzedMessageEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityAnalyzedMessageEvidenceVerdict_Suspicious,
		"unknown":        SecurityAnalyzedMessageEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceVerdict(input)
	return &out, nil
}
