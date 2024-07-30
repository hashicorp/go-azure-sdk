package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceVerdict string

const (
	SecurityKubernetesNamespaceEvidenceVerdict_Malicious      SecurityKubernetesNamespaceEvidenceVerdict = "malicious"
	SecurityKubernetesNamespaceEvidenceVerdict_NoThreatsFound SecurityKubernetesNamespaceEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesNamespaceEvidenceVerdict_Suspicious     SecurityKubernetesNamespaceEvidenceVerdict = "suspicious"
	SecurityKubernetesNamespaceEvidenceVerdict_Unknown        SecurityKubernetesNamespaceEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceVerdict_Malicious),
		string(SecurityKubernetesNamespaceEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesNamespaceEvidenceVerdict_Suspicious),
		string(SecurityKubernetesNamespaceEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesNamespaceEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesNamespaceEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesNamespaceEvidenceVerdict(input string) (*SecurityKubernetesNamespaceEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceVerdict{
		"malicious":      SecurityKubernetesNamespaceEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesNamespaceEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesNamespaceEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesNamespaceEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceVerdict(input)
	return &out, nil
}
