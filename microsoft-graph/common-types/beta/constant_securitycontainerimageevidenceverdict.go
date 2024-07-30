package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceVerdict string

const (
	SecurityContainerImageEvidenceVerdict_Malicious      SecurityContainerImageEvidenceVerdict = "malicious"
	SecurityContainerImageEvidenceVerdict_NoThreatsFound SecurityContainerImageEvidenceVerdict = "noThreatsFound"
	SecurityContainerImageEvidenceVerdict_Suspicious     SecurityContainerImageEvidenceVerdict = "suspicious"
	SecurityContainerImageEvidenceVerdict_Unknown        SecurityContainerImageEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityContainerImageEvidenceVerdict() []string {
	return []string{
		string(SecurityContainerImageEvidenceVerdict_Malicious),
		string(SecurityContainerImageEvidenceVerdict_NoThreatsFound),
		string(SecurityContainerImageEvidenceVerdict_Suspicious),
		string(SecurityContainerImageEvidenceVerdict_Unknown),
	}
}

func (s *SecurityContainerImageEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerImageEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerImageEvidenceVerdict(input string) (*SecurityContainerImageEvidenceVerdict, error) {
	vals := map[string]SecurityContainerImageEvidenceVerdict{
		"malicious":      SecurityContainerImageEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityContainerImageEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityContainerImageEvidenceVerdict_Suspicious,
		"unknown":        SecurityContainerImageEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceVerdict(input)
	return &out, nil
}
