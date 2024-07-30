package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidenceVerdict string

const (
	SecurityContainerRegistryEvidenceVerdict_Malicious      SecurityContainerRegistryEvidenceVerdict = "malicious"
	SecurityContainerRegistryEvidenceVerdict_NoThreatsFound SecurityContainerRegistryEvidenceVerdict = "noThreatsFound"
	SecurityContainerRegistryEvidenceVerdict_Suspicious     SecurityContainerRegistryEvidenceVerdict = "suspicious"
	SecurityContainerRegistryEvidenceVerdict_Unknown        SecurityContainerRegistryEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityContainerRegistryEvidenceVerdict() []string {
	return []string{
		string(SecurityContainerRegistryEvidenceVerdict_Malicious),
		string(SecurityContainerRegistryEvidenceVerdict_NoThreatsFound),
		string(SecurityContainerRegistryEvidenceVerdict_Suspicious),
		string(SecurityContainerRegistryEvidenceVerdict_Unknown),
	}
}

func (s *SecurityContainerRegistryEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerRegistryEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerRegistryEvidenceVerdict(input string) (*SecurityContainerRegistryEvidenceVerdict, error) {
	vals := map[string]SecurityContainerRegistryEvidenceVerdict{
		"malicious":      SecurityContainerRegistryEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityContainerRegistryEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityContainerRegistryEvidenceVerdict_Suspicious,
		"unknown":        SecurityContainerRegistryEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerRegistryEvidenceVerdict(input)
	return &out, nil
}
