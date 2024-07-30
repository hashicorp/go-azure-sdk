package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubRepoEvidenceVerdict string

const (
	SecurityGitHubRepoEvidenceVerdict_Malicious      SecurityGitHubRepoEvidenceVerdict = "malicious"
	SecurityGitHubRepoEvidenceVerdict_NoThreatsFound SecurityGitHubRepoEvidenceVerdict = "noThreatsFound"
	SecurityGitHubRepoEvidenceVerdict_Suspicious     SecurityGitHubRepoEvidenceVerdict = "suspicious"
	SecurityGitHubRepoEvidenceVerdict_Unknown        SecurityGitHubRepoEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityGitHubRepoEvidenceVerdict() []string {
	return []string{
		string(SecurityGitHubRepoEvidenceVerdict_Malicious),
		string(SecurityGitHubRepoEvidenceVerdict_NoThreatsFound),
		string(SecurityGitHubRepoEvidenceVerdict_Suspicious),
		string(SecurityGitHubRepoEvidenceVerdict_Unknown),
	}
}

func (s *SecurityGitHubRepoEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubRepoEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubRepoEvidenceVerdict(input string) (*SecurityGitHubRepoEvidenceVerdict, error) {
	vals := map[string]SecurityGitHubRepoEvidenceVerdict{
		"malicious":      SecurityGitHubRepoEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityGitHubRepoEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityGitHubRepoEvidenceVerdict_Suspicious,
		"unknown":        SecurityGitHubRepoEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubRepoEvidenceVerdict(input)
	return &out, nil
}
