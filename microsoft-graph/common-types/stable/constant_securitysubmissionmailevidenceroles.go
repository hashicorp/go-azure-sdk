package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySubmissionMailEvidenceRoles string

const (
	SecuritySubmissionMailEvidenceRoles_Added             SecuritySubmissionMailEvidenceRoles = "added"
	SecuritySubmissionMailEvidenceRoles_Attacked          SecuritySubmissionMailEvidenceRoles = "attacked"
	SecuritySubmissionMailEvidenceRoles_Attacker          SecuritySubmissionMailEvidenceRoles = "attacker"
	SecuritySubmissionMailEvidenceRoles_CommandAndControl SecuritySubmissionMailEvidenceRoles = "commandAndControl"
	SecuritySubmissionMailEvidenceRoles_Compromised       SecuritySubmissionMailEvidenceRoles = "compromised"
	SecuritySubmissionMailEvidenceRoles_Contextual        SecuritySubmissionMailEvidenceRoles = "contextual"
	SecuritySubmissionMailEvidenceRoles_Created           SecuritySubmissionMailEvidenceRoles = "created"
	SecuritySubmissionMailEvidenceRoles_Destination       SecuritySubmissionMailEvidenceRoles = "destination"
	SecuritySubmissionMailEvidenceRoles_Edited            SecuritySubmissionMailEvidenceRoles = "edited"
	SecuritySubmissionMailEvidenceRoles_Loaded            SecuritySubmissionMailEvidenceRoles = "loaded"
	SecuritySubmissionMailEvidenceRoles_PolicyViolator    SecuritySubmissionMailEvidenceRoles = "policyViolator"
	SecuritySubmissionMailEvidenceRoles_Scanned           SecuritySubmissionMailEvidenceRoles = "scanned"
	SecuritySubmissionMailEvidenceRoles_Source            SecuritySubmissionMailEvidenceRoles = "source"
	SecuritySubmissionMailEvidenceRoles_Suspicious        SecuritySubmissionMailEvidenceRoles = "suspicious"
	SecuritySubmissionMailEvidenceRoles_Unknown           SecuritySubmissionMailEvidenceRoles = "unknown"
)

func PossibleValuesForSecuritySubmissionMailEvidenceRoles() []string {
	return []string{
		string(SecuritySubmissionMailEvidenceRoles_Added),
		string(SecuritySubmissionMailEvidenceRoles_Attacked),
		string(SecuritySubmissionMailEvidenceRoles_Attacker),
		string(SecuritySubmissionMailEvidenceRoles_CommandAndControl),
		string(SecuritySubmissionMailEvidenceRoles_Compromised),
		string(SecuritySubmissionMailEvidenceRoles_Contextual),
		string(SecuritySubmissionMailEvidenceRoles_Created),
		string(SecuritySubmissionMailEvidenceRoles_Destination),
		string(SecuritySubmissionMailEvidenceRoles_Edited),
		string(SecuritySubmissionMailEvidenceRoles_Loaded),
		string(SecuritySubmissionMailEvidenceRoles_PolicyViolator),
		string(SecuritySubmissionMailEvidenceRoles_Scanned),
		string(SecuritySubmissionMailEvidenceRoles_Source),
		string(SecuritySubmissionMailEvidenceRoles_Suspicious),
		string(SecuritySubmissionMailEvidenceRoles_Unknown),
	}
}

func (s *SecuritySubmissionMailEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecuritySubmissionMailEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecuritySubmissionMailEvidenceRoles(input string) (*SecuritySubmissionMailEvidenceRoles, error) {
	vals := map[string]SecuritySubmissionMailEvidenceRoles{
		"added":             SecuritySubmissionMailEvidenceRoles_Added,
		"attacked":          SecuritySubmissionMailEvidenceRoles_Attacked,
		"attacker":          SecuritySubmissionMailEvidenceRoles_Attacker,
		"commandandcontrol": SecuritySubmissionMailEvidenceRoles_CommandAndControl,
		"compromised":       SecuritySubmissionMailEvidenceRoles_Compromised,
		"contextual":        SecuritySubmissionMailEvidenceRoles_Contextual,
		"created":           SecuritySubmissionMailEvidenceRoles_Created,
		"destination":       SecuritySubmissionMailEvidenceRoles_Destination,
		"edited":            SecuritySubmissionMailEvidenceRoles_Edited,
		"loaded":            SecuritySubmissionMailEvidenceRoles_Loaded,
		"policyviolator":    SecuritySubmissionMailEvidenceRoles_PolicyViolator,
		"scanned":           SecuritySubmissionMailEvidenceRoles_Scanned,
		"source":            SecuritySubmissionMailEvidenceRoles_Source,
		"suspicious":        SecuritySubmissionMailEvidenceRoles_Suspicious,
		"unknown":           SecuritySubmissionMailEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySubmissionMailEvidenceRoles(input)
	return &out, nil
}
