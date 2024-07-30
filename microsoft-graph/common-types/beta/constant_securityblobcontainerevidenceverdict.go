package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceVerdict string

const (
	SecurityBlobContainerEvidenceVerdict_Malicious      SecurityBlobContainerEvidenceVerdict = "malicious"
	SecurityBlobContainerEvidenceVerdict_NoThreatsFound SecurityBlobContainerEvidenceVerdict = "noThreatsFound"
	SecurityBlobContainerEvidenceVerdict_Suspicious     SecurityBlobContainerEvidenceVerdict = "suspicious"
	SecurityBlobContainerEvidenceVerdict_Unknown        SecurityBlobContainerEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityBlobContainerEvidenceVerdict() []string {
	return []string{
		string(SecurityBlobContainerEvidenceVerdict_Malicious),
		string(SecurityBlobContainerEvidenceVerdict_NoThreatsFound),
		string(SecurityBlobContainerEvidenceVerdict_Suspicious),
		string(SecurityBlobContainerEvidenceVerdict_Unknown),
	}
}

func (s *SecurityBlobContainerEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobContainerEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobContainerEvidenceVerdict(input string) (*SecurityBlobContainerEvidenceVerdict, error) {
	vals := map[string]SecurityBlobContainerEvidenceVerdict{
		"malicious":      SecurityBlobContainerEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityBlobContainerEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityBlobContainerEvidenceVerdict_Suspicious,
		"unknown":        SecurityBlobContainerEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceVerdict(input)
	return &out, nil
}
