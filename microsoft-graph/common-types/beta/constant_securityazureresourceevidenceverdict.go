package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceVerdict string

const (
	SecurityAzureResourceEvidenceVerdict_Malicious      SecurityAzureResourceEvidenceVerdict = "malicious"
	SecurityAzureResourceEvidenceVerdict_NoThreatsFound SecurityAzureResourceEvidenceVerdict = "noThreatsFound"
	SecurityAzureResourceEvidenceVerdict_Suspicious     SecurityAzureResourceEvidenceVerdict = "suspicious"
	SecurityAzureResourceEvidenceVerdict_Unknown        SecurityAzureResourceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityAzureResourceEvidenceVerdict() []string {
	return []string{
		string(SecurityAzureResourceEvidenceVerdict_Malicious),
		string(SecurityAzureResourceEvidenceVerdict_NoThreatsFound),
		string(SecurityAzureResourceEvidenceVerdict_Suspicious),
		string(SecurityAzureResourceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityAzureResourceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAzureResourceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAzureResourceEvidenceVerdict(input string) (*SecurityAzureResourceEvidenceVerdict, error) {
	vals := map[string]SecurityAzureResourceEvidenceVerdict{
		"malicious":      SecurityAzureResourceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityAzureResourceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityAzureResourceEvidenceVerdict_Suspicious,
		"unknown":        SecurityAzureResourceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceVerdict(input)
	return &out, nil
}
