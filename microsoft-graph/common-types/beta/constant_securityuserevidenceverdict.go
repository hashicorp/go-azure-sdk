package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidenceVerdict string

const (
	SecurityUserEvidenceVerdict_Malicious      SecurityUserEvidenceVerdict = "malicious"
	SecurityUserEvidenceVerdict_NoThreatsFound SecurityUserEvidenceVerdict = "noThreatsFound"
	SecurityUserEvidenceVerdict_Suspicious     SecurityUserEvidenceVerdict = "suspicious"
	SecurityUserEvidenceVerdict_Unknown        SecurityUserEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityUserEvidenceVerdict() []string {
	return []string{
		string(SecurityUserEvidenceVerdict_Malicious),
		string(SecurityUserEvidenceVerdict_NoThreatsFound),
		string(SecurityUserEvidenceVerdict_Suspicious),
		string(SecurityUserEvidenceVerdict_Unknown),
	}
}

func (s *SecurityUserEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserEvidenceVerdict(input string) (*SecurityUserEvidenceVerdict, error) {
	vals := map[string]SecurityUserEvidenceVerdict{
		"malicious":      SecurityUserEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityUserEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityUserEvidenceVerdict_Suspicious,
		"unknown":        SecurityUserEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserEvidenceVerdict(input)
	return &out, nil
}
