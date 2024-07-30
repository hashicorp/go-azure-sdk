package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySasTokenEvidenceVerdict string

const (
	SecuritySasTokenEvidenceVerdict_Malicious      SecuritySasTokenEvidenceVerdict = "malicious"
	SecuritySasTokenEvidenceVerdict_NoThreatsFound SecuritySasTokenEvidenceVerdict = "noThreatsFound"
	SecuritySasTokenEvidenceVerdict_Suspicious     SecuritySasTokenEvidenceVerdict = "suspicious"
	SecuritySasTokenEvidenceVerdict_Unknown        SecuritySasTokenEvidenceVerdict = "unknown"
)

func PossibleValuesForSecuritySasTokenEvidenceVerdict() []string {
	return []string{
		string(SecuritySasTokenEvidenceVerdict_Malicious),
		string(SecuritySasTokenEvidenceVerdict_NoThreatsFound),
		string(SecuritySasTokenEvidenceVerdict_Suspicious),
		string(SecuritySasTokenEvidenceVerdict_Unknown),
	}
}

func (s *SecuritySasTokenEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySasTokenEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySasTokenEvidenceVerdict(input string) (*SecuritySasTokenEvidenceVerdict, error) {
	vals := map[string]SecuritySasTokenEvidenceVerdict{
		"malicious":      SecuritySasTokenEvidenceVerdict_Malicious,
		"nothreatsfound": SecuritySasTokenEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecuritySasTokenEvidenceVerdict_Suspicious,
		"unknown":        SecuritySasTokenEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySasTokenEvidenceVerdict(input)
	return &out, nil
}
