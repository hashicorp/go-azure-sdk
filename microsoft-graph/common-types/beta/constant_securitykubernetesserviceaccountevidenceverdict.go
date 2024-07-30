package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidenceVerdict string

const (
	SecurityKubernetesServiceAccountEvidenceVerdict_Malicious      SecurityKubernetesServiceAccountEvidenceVerdict = "malicious"
	SecurityKubernetesServiceAccountEvidenceVerdict_NoThreatsFound SecurityKubernetesServiceAccountEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesServiceAccountEvidenceVerdict_Suspicious     SecurityKubernetesServiceAccountEvidenceVerdict = "suspicious"
	SecurityKubernetesServiceAccountEvidenceVerdict_Unknown        SecurityKubernetesServiceAccountEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesServiceAccountEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesServiceAccountEvidenceVerdict_Malicious),
		string(SecurityKubernetesServiceAccountEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesServiceAccountEvidenceVerdict_Suspicious),
		string(SecurityKubernetesServiceAccountEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesServiceAccountEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesServiceAccountEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesServiceAccountEvidenceVerdict(input string) (*SecurityKubernetesServiceAccountEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesServiceAccountEvidenceVerdict{
		"malicious":      SecurityKubernetesServiceAccountEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesServiceAccountEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesServiceAccountEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesServiceAccountEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceAccountEvidenceVerdict(input)
	return &out, nil
}
