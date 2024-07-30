package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceRoles string

const (
	SecurityContainerEvidenceRoles_Added             SecurityContainerEvidenceRoles = "added"
	SecurityContainerEvidenceRoles_Attacked          SecurityContainerEvidenceRoles = "attacked"
	SecurityContainerEvidenceRoles_Attacker          SecurityContainerEvidenceRoles = "attacker"
	SecurityContainerEvidenceRoles_CommandAndControl SecurityContainerEvidenceRoles = "commandAndControl"
	SecurityContainerEvidenceRoles_Compromised       SecurityContainerEvidenceRoles = "compromised"
	SecurityContainerEvidenceRoles_Contextual        SecurityContainerEvidenceRoles = "contextual"
	SecurityContainerEvidenceRoles_Created           SecurityContainerEvidenceRoles = "created"
	SecurityContainerEvidenceRoles_Destination       SecurityContainerEvidenceRoles = "destination"
	SecurityContainerEvidenceRoles_Edited            SecurityContainerEvidenceRoles = "edited"
	SecurityContainerEvidenceRoles_Loaded            SecurityContainerEvidenceRoles = "loaded"
	SecurityContainerEvidenceRoles_PolicyViolator    SecurityContainerEvidenceRoles = "policyViolator"
	SecurityContainerEvidenceRoles_Scanned           SecurityContainerEvidenceRoles = "scanned"
	SecurityContainerEvidenceRoles_Source            SecurityContainerEvidenceRoles = "source"
	SecurityContainerEvidenceRoles_Suspicious        SecurityContainerEvidenceRoles = "suspicious"
	SecurityContainerEvidenceRoles_Unknown           SecurityContainerEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityContainerEvidenceRoles() []string {
	return []string{
		string(SecurityContainerEvidenceRoles_Added),
		string(SecurityContainerEvidenceRoles_Attacked),
		string(SecurityContainerEvidenceRoles_Attacker),
		string(SecurityContainerEvidenceRoles_CommandAndControl),
		string(SecurityContainerEvidenceRoles_Compromised),
		string(SecurityContainerEvidenceRoles_Contextual),
		string(SecurityContainerEvidenceRoles_Created),
		string(SecurityContainerEvidenceRoles_Destination),
		string(SecurityContainerEvidenceRoles_Edited),
		string(SecurityContainerEvidenceRoles_Loaded),
		string(SecurityContainerEvidenceRoles_PolicyViolator),
		string(SecurityContainerEvidenceRoles_Scanned),
		string(SecurityContainerEvidenceRoles_Source),
		string(SecurityContainerEvidenceRoles_Suspicious),
		string(SecurityContainerEvidenceRoles_Unknown),
	}
}

func (s *SecurityContainerEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerEvidenceRoles(input string) (*SecurityContainerEvidenceRoles, error) {
	vals := map[string]SecurityContainerEvidenceRoles{
		"added":             SecurityContainerEvidenceRoles_Added,
		"attacked":          SecurityContainerEvidenceRoles_Attacked,
		"attacker":          SecurityContainerEvidenceRoles_Attacker,
		"commandandcontrol": SecurityContainerEvidenceRoles_CommandAndControl,
		"compromised":       SecurityContainerEvidenceRoles_Compromised,
		"contextual":        SecurityContainerEvidenceRoles_Contextual,
		"created":           SecurityContainerEvidenceRoles_Created,
		"destination":       SecurityContainerEvidenceRoles_Destination,
		"edited":            SecurityContainerEvidenceRoles_Edited,
		"loaded":            SecurityContainerEvidenceRoles_Loaded,
		"policyviolator":    SecurityContainerEvidenceRoles_PolicyViolator,
		"scanned":           SecurityContainerEvidenceRoles_Scanned,
		"source":            SecurityContainerEvidenceRoles_Source,
		"suspicious":        SecurityContainerEvidenceRoles_Suspicious,
		"unknown":           SecurityContainerEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerEvidenceRoles(input)
	return &out, nil
}
