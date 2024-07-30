package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceVerdict string

const (
	SecurityAlertEvidenceVerdict_Malicious      SecurityAlertEvidenceVerdict = "malicious"
	SecurityAlertEvidenceVerdict_NoThreatsFound SecurityAlertEvidenceVerdict = "noThreatsFound"
	SecurityAlertEvidenceVerdict_Suspicious     SecurityAlertEvidenceVerdict = "suspicious"
	SecurityAlertEvidenceVerdict_Unknown        SecurityAlertEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityAlertEvidenceVerdict() []string {
	return []string{
		string(SecurityAlertEvidenceVerdict_Malicious),
		string(SecurityAlertEvidenceVerdict_NoThreatsFound),
		string(SecurityAlertEvidenceVerdict_Suspicious),
		string(SecurityAlertEvidenceVerdict_Unknown),
	}
}

func (s *SecurityAlertEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertEvidenceVerdict(input string) (*SecurityAlertEvidenceVerdict, error) {
	vals := map[string]SecurityAlertEvidenceVerdict{
		"malicious":      SecurityAlertEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityAlertEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityAlertEvidenceVerdict_Suspicious,
		"unknown":        SecurityAlertEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceVerdict(input)
	return &out, nil
}
