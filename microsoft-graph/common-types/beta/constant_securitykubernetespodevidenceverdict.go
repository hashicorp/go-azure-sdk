package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidenceVerdict string

const (
	SecurityKubernetesPodEvidenceVerdict_Malicious      SecurityKubernetesPodEvidenceVerdict = "malicious"
	SecurityKubernetesPodEvidenceVerdict_NoThreatsFound SecurityKubernetesPodEvidenceVerdict = "noThreatsFound"
	SecurityKubernetesPodEvidenceVerdict_Suspicious     SecurityKubernetesPodEvidenceVerdict = "suspicious"
	SecurityKubernetesPodEvidenceVerdict_Unknown        SecurityKubernetesPodEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityKubernetesPodEvidenceVerdict() []string {
	return []string{
		string(SecurityKubernetesPodEvidenceVerdict_Malicious),
		string(SecurityKubernetesPodEvidenceVerdict_NoThreatsFound),
		string(SecurityKubernetesPodEvidenceVerdict_Suspicious),
		string(SecurityKubernetesPodEvidenceVerdict_Unknown),
	}
}

func (s *SecurityKubernetesPodEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesPodEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesPodEvidenceVerdict(input string) (*SecurityKubernetesPodEvidenceVerdict, error) {
	vals := map[string]SecurityKubernetesPodEvidenceVerdict{
		"malicious":      SecurityKubernetesPodEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityKubernetesPodEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityKubernetesPodEvidenceVerdict_Suspicious,
		"unknown":        SecurityKubernetesPodEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPodEvidenceVerdict(input)
	return &out, nil
}
