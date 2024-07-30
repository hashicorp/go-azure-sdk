package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceRoles string

const (
	SecurityProcessEvidenceRoles_Added             SecurityProcessEvidenceRoles = "added"
	SecurityProcessEvidenceRoles_Attacked          SecurityProcessEvidenceRoles = "attacked"
	SecurityProcessEvidenceRoles_Attacker          SecurityProcessEvidenceRoles = "attacker"
	SecurityProcessEvidenceRoles_CommandAndControl SecurityProcessEvidenceRoles = "commandAndControl"
	SecurityProcessEvidenceRoles_Compromised       SecurityProcessEvidenceRoles = "compromised"
	SecurityProcessEvidenceRoles_Contextual        SecurityProcessEvidenceRoles = "contextual"
	SecurityProcessEvidenceRoles_Created           SecurityProcessEvidenceRoles = "created"
	SecurityProcessEvidenceRoles_Destination       SecurityProcessEvidenceRoles = "destination"
	SecurityProcessEvidenceRoles_Edited            SecurityProcessEvidenceRoles = "edited"
	SecurityProcessEvidenceRoles_Loaded            SecurityProcessEvidenceRoles = "loaded"
	SecurityProcessEvidenceRoles_PolicyViolator    SecurityProcessEvidenceRoles = "policyViolator"
	SecurityProcessEvidenceRoles_Scanned           SecurityProcessEvidenceRoles = "scanned"
	SecurityProcessEvidenceRoles_Source            SecurityProcessEvidenceRoles = "source"
	SecurityProcessEvidenceRoles_Suspicious        SecurityProcessEvidenceRoles = "suspicious"
	SecurityProcessEvidenceRoles_Unknown           SecurityProcessEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityProcessEvidenceRoles() []string {
	return []string{
		string(SecurityProcessEvidenceRoles_Added),
		string(SecurityProcessEvidenceRoles_Attacked),
		string(SecurityProcessEvidenceRoles_Attacker),
		string(SecurityProcessEvidenceRoles_CommandAndControl),
		string(SecurityProcessEvidenceRoles_Compromised),
		string(SecurityProcessEvidenceRoles_Contextual),
		string(SecurityProcessEvidenceRoles_Created),
		string(SecurityProcessEvidenceRoles_Destination),
		string(SecurityProcessEvidenceRoles_Edited),
		string(SecurityProcessEvidenceRoles_Loaded),
		string(SecurityProcessEvidenceRoles_PolicyViolator),
		string(SecurityProcessEvidenceRoles_Scanned),
		string(SecurityProcessEvidenceRoles_Source),
		string(SecurityProcessEvidenceRoles_Suspicious),
		string(SecurityProcessEvidenceRoles_Unknown),
	}
}

func (s *SecurityProcessEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityProcessEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityProcessEvidenceRoles(input string) (*SecurityProcessEvidenceRoles, error) {
	vals := map[string]SecurityProcessEvidenceRoles{
		"added":             SecurityProcessEvidenceRoles_Added,
		"attacked":          SecurityProcessEvidenceRoles_Attacked,
		"attacker":          SecurityProcessEvidenceRoles_Attacker,
		"commandandcontrol": SecurityProcessEvidenceRoles_CommandAndControl,
		"compromised":       SecurityProcessEvidenceRoles_Compromised,
		"contextual":        SecurityProcessEvidenceRoles_Contextual,
		"created":           SecurityProcessEvidenceRoles_Created,
		"destination":       SecurityProcessEvidenceRoles_Destination,
		"edited":            SecurityProcessEvidenceRoles_Edited,
		"loaded":            SecurityProcessEvidenceRoles_Loaded,
		"policyviolator":    SecurityProcessEvidenceRoles_PolicyViolator,
		"scanned":           SecurityProcessEvidenceRoles_Scanned,
		"source":            SecurityProcessEvidenceRoles_Source,
		"suspicious":        SecurityProcessEvidenceRoles_Suspicious,
		"unknown":           SecurityProcessEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceRoles(input)
	return &out, nil
}
