package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOauthApplicationEvidenceVerdict string

const (
	SecurityOauthApplicationEvidenceVerdict_Malicious      SecurityOauthApplicationEvidenceVerdict = "malicious"
	SecurityOauthApplicationEvidenceVerdict_NoThreatsFound SecurityOauthApplicationEvidenceVerdict = "noThreatsFound"
	SecurityOauthApplicationEvidenceVerdict_Suspicious     SecurityOauthApplicationEvidenceVerdict = "suspicious"
	SecurityOauthApplicationEvidenceVerdict_Unknown        SecurityOauthApplicationEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityOauthApplicationEvidenceVerdict() []string {
	return []string{
		string(SecurityOauthApplicationEvidenceVerdict_Malicious),
		string(SecurityOauthApplicationEvidenceVerdict_NoThreatsFound),
		string(SecurityOauthApplicationEvidenceVerdict_Suspicious),
		string(SecurityOauthApplicationEvidenceVerdict_Unknown),
	}
}

func (s *SecurityOauthApplicationEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityOauthApplicationEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityOauthApplicationEvidenceVerdict(input string) (*SecurityOauthApplicationEvidenceVerdict, error) {
	vals := map[string]SecurityOauthApplicationEvidenceVerdict{
		"malicious":      SecurityOauthApplicationEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityOauthApplicationEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityOauthApplicationEvidenceVerdict_Suspicious,
		"unknown":        SecurityOauthApplicationEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityOauthApplicationEvidenceVerdict(input)
	return &out, nil
}
