package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUrlEvidenceRoles string

const (
	SecurityUrlEvidenceRoles_Added             SecurityUrlEvidenceRoles = "added"
	SecurityUrlEvidenceRoles_Attacked          SecurityUrlEvidenceRoles = "attacked"
	SecurityUrlEvidenceRoles_Attacker          SecurityUrlEvidenceRoles = "attacker"
	SecurityUrlEvidenceRoles_CommandAndControl SecurityUrlEvidenceRoles = "commandAndControl"
	SecurityUrlEvidenceRoles_Compromised       SecurityUrlEvidenceRoles = "compromised"
	SecurityUrlEvidenceRoles_Contextual        SecurityUrlEvidenceRoles = "contextual"
	SecurityUrlEvidenceRoles_Created           SecurityUrlEvidenceRoles = "created"
	SecurityUrlEvidenceRoles_Destination       SecurityUrlEvidenceRoles = "destination"
	SecurityUrlEvidenceRoles_Edited            SecurityUrlEvidenceRoles = "edited"
	SecurityUrlEvidenceRoles_Loaded            SecurityUrlEvidenceRoles = "loaded"
	SecurityUrlEvidenceRoles_PolicyViolator    SecurityUrlEvidenceRoles = "policyViolator"
	SecurityUrlEvidenceRoles_Scanned           SecurityUrlEvidenceRoles = "scanned"
	SecurityUrlEvidenceRoles_Source            SecurityUrlEvidenceRoles = "source"
	SecurityUrlEvidenceRoles_Suspicious        SecurityUrlEvidenceRoles = "suspicious"
	SecurityUrlEvidenceRoles_Unknown           SecurityUrlEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityUrlEvidenceRoles() []string {
	return []string{
		string(SecurityUrlEvidenceRoles_Added),
		string(SecurityUrlEvidenceRoles_Attacked),
		string(SecurityUrlEvidenceRoles_Attacker),
		string(SecurityUrlEvidenceRoles_CommandAndControl),
		string(SecurityUrlEvidenceRoles_Compromised),
		string(SecurityUrlEvidenceRoles_Contextual),
		string(SecurityUrlEvidenceRoles_Created),
		string(SecurityUrlEvidenceRoles_Destination),
		string(SecurityUrlEvidenceRoles_Edited),
		string(SecurityUrlEvidenceRoles_Loaded),
		string(SecurityUrlEvidenceRoles_PolicyViolator),
		string(SecurityUrlEvidenceRoles_Scanned),
		string(SecurityUrlEvidenceRoles_Source),
		string(SecurityUrlEvidenceRoles_Suspicious),
		string(SecurityUrlEvidenceRoles_Unknown),
	}
}

func (s *SecurityUrlEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUrlEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUrlEvidenceRoles(input string) (*SecurityUrlEvidenceRoles, error) {
	vals := map[string]SecurityUrlEvidenceRoles{
		"added":             SecurityUrlEvidenceRoles_Added,
		"attacked":          SecurityUrlEvidenceRoles_Attacked,
		"attacker":          SecurityUrlEvidenceRoles_Attacker,
		"commandandcontrol": SecurityUrlEvidenceRoles_CommandAndControl,
		"compromised":       SecurityUrlEvidenceRoles_Compromised,
		"contextual":        SecurityUrlEvidenceRoles_Contextual,
		"created":           SecurityUrlEvidenceRoles_Created,
		"destination":       SecurityUrlEvidenceRoles_Destination,
		"edited":            SecurityUrlEvidenceRoles_Edited,
		"loaded":            SecurityUrlEvidenceRoles_Loaded,
		"policyviolator":    SecurityUrlEvidenceRoles_PolicyViolator,
		"scanned":           SecurityUrlEvidenceRoles_Scanned,
		"source":            SecurityUrlEvidenceRoles_Source,
		"suspicious":        SecurityUrlEvidenceRoles_Suspicious,
		"unknown":           SecurityUrlEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUrlEvidenceRoles(input)
	return &out, nil
}
