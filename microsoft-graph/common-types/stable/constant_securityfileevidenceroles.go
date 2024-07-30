package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceRoles string

const (
	SecurityFileEvidenceRoles_Added             SecurityFileEvidenceRoles = "added"
	SecurityFileEvidenceRoles_Attacked          SecurityFileEvidenceRoles = "attacked"
	SecurityFileEvidenceRoles_Attacker          SecurityFileEvidenceRoles = "attacker"
	SecurityFileEvidenceRoles_CommandAndControl SecurityFileEvidenceRoles = "commandAndControl"
	SecurityFileEvidenceRoles_Compromised       SecurityFileEvidenceRoles = "compromised"
	SecurityFileEvidenceRoles_Contextual        SecurityFileEvidenceRoles = "contextual"
	SecurityFileEvidenceRoles_Created           SecurityFileEvidenceRoles = "created"
	SecurityFileEvidenceRoles_Destination       SecurityFileEvidenceRoles = "destination"
	SecurityFileEvidenceRoles_Edited            SecurityFileEvidenceRoles = "edited"
	SecurityFileEvidenceRoles_Loaded            SecurityFileEvidenceRoles = "loaded"
	SecurityFileEvidenceRoles_PolicyViolator    SecurityFileEvidenceRoles = "policyViolator"
	SecurityFileEvidenceRoles_Scanned           SecurityFileEvidenceRoles = "scanned"
	SecurityFileEvidenceRoles_Source            SecurityFileEvidenceRoles = "source"
	SecurityFileEvidenceRoles_Suspicious        SecurityFileEvidenceRoles = "suspicious"
	SecurityFileEvidenceRoles_Unknown           SecurityFileEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityFileEvidenceRoles() []string {
	return []string{
		string(SecurityFileEvidenceRoles_Added),
		string(SecurityFileEvidenceRoles_Attacked),
		string(SecurityFileEvidenceRoles_Attacker),
		string(SecurityFileEvidenceRoles_CommandAndControl),
		string(SecurityFileEvidenceRoles_Compromised),
		string(SecurityFileEvidenceRoles_Contextual),
		string(SecurityFileEvidenceRoles_Created),
		string(SecurityFileEvidenceRoles_Destination),
		string(SecurityFileEvidenceRoles_Edited),
		string(SecurityFileEvidenceRoles_Loaded),
		string(SecurityFileEvidenceRoles_PolicyViolator),
		string(SecurityFileEvidenceRoles_Scanned),
		string(SecurityFileEvidenceRoles_Source),
		string(SecurityFileEvidenceRoles_Suspicious),
		string(SecurityFileEvidenceRoles_Unknown),
	}
}

func (s *SecurityFileEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityFileEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityFileEvidenceRoles(input string) (*SecurityFileEvidenceRoles, error) {
	vals := map[string]SecurityFileEvidenceRoles{
		"added":             SecurityFileEvidenceRoles_Added,
		"attacked":          SecurityFileEvidenceRoles_Attacked,
		"attacker":          SecurityFileEvidenceRoles_Attacker,
		"commandandcontrol": SecurityFileEvidenceRoles_CommandAndControl,
		"compromised":       SecurityFileEvidenceRoles_Compromised,
		"contextual":        SecurityFileEvidenceRoles_Contextual,
		"created":           SecurityFileEvidenceRoles_Created,
		"destination":       SecurityFileEvidenceRoles_Destination,
		"edited":            SecurityFileEvidenceRoles_Edited,
		"loaded":            SecurityFileEvidenceRoles_Loaded,
		"policyviolator":    SecurityFileEvidenceRoles_PolicyViolator,
		"scanned":           SecurityFileEvidenceRoles_Scanned,
		"source":            SecurityFileEvidenceRoles_Source,
		"suspicious":        SecurityFileEvidenceRoles_Suspicious,
		"unknown":           SecurityFileEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceRoles(input)
	return &out, nil
}
