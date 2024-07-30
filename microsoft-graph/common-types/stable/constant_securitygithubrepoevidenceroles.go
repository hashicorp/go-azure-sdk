package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGitHubRepoEvidenceRoles string

const (
	SecurityGitHubRepoEvidenceRoles_Added             SecurityGitHubRepoEvidenceRoles = "added"
	SecurityGitHubRepoEvidenceRoles_Attacked          SecurityGitHubRepoEvidenceRoles = "attacked"
	SecurityGitHubRepoEvidenceRoles_Attacker          SecurityGitHubRepoEvidenceRoles = "attacker"
	SecurityGitHubRepoEvidenceRoles_CommandAndControl SecurityGitHubRepoEvidenceRoles = "commandAndControl"
	SecurityGitHubRepoEvidenceRoles_Compromised       SecurityGitHubRepoEvidenceRoles = "compromised"
	SecurityGitHubRepoEvidenceRoles_Contextual        SecurityGitHubRepoEvidenceRoles = "contextual"
	SecurityGitHubRepoEvidenceRoles_Created           SecurityGitHubRepoEvidenceRoles = "created"
	SecurityGitHubRepoEvidenceRoles_Destination       SecurityGitHubRepoEvidenceRoles = "destination"
	SecurityGitHubRepoEvidenceRoles_Edited            SecurityGitHubRepoEvidenceRoles = "edited"
	SecurityGitHubRepoEvidenceRoles_Loaded            SecurityGitHubRepoEvidenceRoles = "loaded"
	SecurityGitHubRepoEvidenceRoles_PolicyViolator    SecurityGitHubRepoEvidenceRoles = "policyViolator"
	SecurityGitHubRepoEvidenceRoles_Scanned           SecurityGitHubRepoEvidenceRoles = "scanned"
	SecurityGitHubRepoEvidenceRoles_Source            SecurityGitHubRepoEvidenceRoles = "source"
	SecurityGitHubRepoEvidenceRoles_Suspicious        SecurityGitHubRepoEvidenceRoles = "suspicious"
	SecurityGitHubRepoEvidenceRoles_Unknown           SecurityGitHubRepoEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityGitHubRepoEvidenceRoles() []string {
	return []string{
		string(SecurityGitHubRepoEvidenceRoles_Added),
		string(SecurityGitHubRepoEvidenceRoles_Attacked),
		string(SecurityGitHubRepoEvidenceRoles_Attacker),
		string(SecurityGitHubRepoEvidenceRoles_CommandAndControl),
		string(SecurityGitHubRepoEvidenceRoles_Compromised),
		string(SecurityGitHubRepoEvidenceRoles_Contextual),
		string(SecurityGitHubRepoEvidenceRoles_Created),
		string(SecurityGitHubRepoEvidenceRoles_Destination),
		string(SecurityGitHubRepoEvidenceRoles_Edited),
		string(SecurityGitHubRepoEvidenceRoles_Loaded),
		string(SecurityGitHubRepoEvidenceRoles_PolicyViolator),
		string(SecurityGitHubRepoEvidenceRoles_Scanned),
		string(SecurityGitHubRepoEvidenceRoles_Source),
		string(SecurityGitHubRepoEvidenceRoles_Suspicious),
		string(SecurityGitHubRepoEvidenceRoles_Unknown),
	}
}

func (s *SecurityGitHubRepoEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGitHubRepoEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGitHubRepoEvidenceRoles(input string) (*SecurityGitHubRepoEvidenceRoles, error) {
	vals := map[string]SecurityGitHubRepoEvidenceRoles{
		"added":             SecurityGitHubRepoEvidenceRoles_Added,
		"attacked":          SecurityGitHubRepoEvidenceRoles_Attacked,
		"attacker":          SecurityGitHubRepoEvidenceRoles_Attacker,
		"commandandcontrol": SecurityGitHubRepoEvidenceRoles_CommandAndControl,
		"compromised":       SecurityGitHubRepoEvidenceRoles_Compromised,
		"contextual":        SecurityGitHubRepoEvidenceRoles_Contextual,
		"created":           SecurityGitHubRepoEvidenceRoles_Created,
		"destination":       SecurityGitHubRepoEvidenceRoles_Destination,
		"edited":            SecurityGitHubRepoEvidenceRoles_Edited,
		"loaded":            SecurityGitHubRepoEvidenceRoles_Loaded,
		"policyviolator":    SecurityGitHubRepoEvidenceRoles_PolicyViolator,
		"scanned":           SecurityGitHubRepoEvidenceRoles_Scanned,
		"source":            SecurityGitHubRepoEvidenceRoles_Source,
		"suspicious":        SecurityGitHubRepoEvidenceRoles_Suspicious,
		"unknown":           SecurityGitHubRepoEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGitHubRepoEvidenceRoles(input)
	return &out, nil
}
