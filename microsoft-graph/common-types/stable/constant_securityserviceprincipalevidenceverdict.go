package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityServicePrincipalEvidenceVerdict string

const (
	SecurityServicePrincipalEvidenceVerdict_Malicious      SecurityServicePrincipalEvidenceVerdict = "malicious"
	SecurityServicePrincipalEvidenceVerdict_NoThreatsFound SecurityServicePrincipalEvidenceVerdict = "noThreatsFound"
	SecurityServicePrincipalEvidenceVerdict_Suspicious     SecurityServicePrincipalEvidenceVerdict = "suspicious"
	SecurityServicePrincipalEvidenceVerdict_Unknown        SecurityServicePrincipalEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityServicePrincipalEvidenceVerdict() []string {
	return []string{
		string(SecurityServicePrincipalEvidenceVerdict_Malicious),
		string(SecurityServicePrincipalEvidenceVerdict_NoThreatsFound),
		string(SecurityServicePrincipalEvidenceVerdict_Suspicious),
		string(SecurityServicePrincipalEvidenceVerdict_Unknown),
	}
}

func (s *SecurityServicePrincipalEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityServicePrincipalEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityServicePrincipalEvidenceVerdict(input string) (*SecurityServicePrincipalEvidenceVerdict, error) {
	vals := map[string]SecurityServicePrincipalEvidenceVerdict{
		"malicious":      SecurityServicePrincipalEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityServicePrincipalEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityServicePrincipalEvidenceVerdict_Suspicious,
		"unknown":        SecurityServicePrincipalEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityServicePrincipalEvidenceVerdict(input)
	return &out, nil
}
