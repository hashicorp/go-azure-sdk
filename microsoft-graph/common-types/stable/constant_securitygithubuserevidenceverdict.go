package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubUserEvidenceVerdict string

const (
	SecurityGitHubUserEvidenceVerdict_Malicious      SecurityGitHubUserEvidenceVerdict = "malicious"
	SecurityGitHubUserEvidenceVerdict_NoThreatsFound SecurityGitHubUserEvidenceVerdict = "noThreatsFound"
	SecurityGitHubUserEvidenceVerdict_Suspicious     SecurityGitHubUserEvidenceVerdict = "suspicious"
	SecurityGitHubUserEvidenceVerdict_Unknown        SecurityGitHubUserEvidenceVerdict = "unknown"
)

func PossibleValuesForSecurityGitHubUserEvidenceVerdict() []string {
	return []string{
		string(SecurityGitHubUserEvidenceVerdict_Malicious),
		string(SecurityGitHubUserEvidenceVerdict_NoThreatsFound),
		string(SecurityGitHubUserEvidenceVerdict_Suspicious),
		string(SecurityGitHubUserEvidenceVerdict_Unknown),
	}
}

func (s *SecurityGitHubUserEvidenceVerdict) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubUserEvidenceVerdict(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubUserEvidenceVerdict(input string) (*SecurityGitHubUserEvidenceVerdict, error) {
	vals := map[string]SecurityGitHubUserEvidenceVerdict{
		"malicious":      SecurityGitHubUserEvidenceVerdict_Malicious,
		"nothreatsfound": SecurityGitHubUserEvidenceVerdict_NoThreatsFound,
		"suspicious":     SecurityGitHubUserEvidenceVerdict_Suspicious,
		"unknown":        SecurityGitHubUserEvidenceVerdict_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubUserEvidenceVerdict(input)
	return &out, nil
}
