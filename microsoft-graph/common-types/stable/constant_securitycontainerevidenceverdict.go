package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceVerdict string

const (
	SecurityContainerEvidenceVerdict_Malicious      SecurityContainerEvidenceVerdict = "malicious"
	SecurityContainerEvidenceVerdict_NoThreatsFound SecurityContainerEvidenceVerdict = "noThreatsFound"
	SecurityContainerEvidenceVerdict_Suspicious     SecurityContainerEvidenceVerdict = "suspicious"
	SecurityContainerEvidenceVerdict_Unknown        SecurityContainerEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityContainerEvidenceVerdict() []string {
	return []string{
		string(SecurityContainerEvidenceVerdict_Malicious),
		string(SecurityContainerEvidenceVerdict_NoThreatsFound),
		string(SecurityContainerEvidenceVerdict_Suspicious),
		string(SecurityContainerEvidenceVerdict_Unknown),
	}
}

func (s *SecurityContainerEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerEvidenceVerdict(input string) (*SecurityContainerEvidenceVerdict, error) {
	vals := map[string]SecurityContainerEvidenceVerdict{
		"malicious":      SecurityContainerEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityContainerEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityContainerEvidenceVerdict_Suspicious,
		"unknown":        SecurityContainerEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerEvidenceVerdict(input)
	return &out, nil
}
