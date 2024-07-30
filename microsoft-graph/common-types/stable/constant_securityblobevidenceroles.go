package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobEvidenceRoles string

const (
	SecurityBlobEvidenceRoles_Added             SecurityBlobEvidenceRoles = "added"
	SecurityBlobEvidenceRoles_Attacked          SecurityBlobEvidenceRoles = "attacked"
	SecurityBlobEvidenceRoles_Attacker          SecurityBlobEvidenceRoles = "attacker"
	SecurityBlobEvidenceRoles_CommandAndControl SecurityBlobEvidenceRoles = "commandAndControl"
	SecurityBlobEvidenceRoles_Compromised       SecurityBlobEvidenceRoles = "compromised"
	SecurityBlobEvidenceRoles_Contextual        SecurityBlobEvidenceRoles = "contextual"
	SecurityBlobEvidenceRoles_Created           SecurityBlobEvidenceRoles = "created"
	SecurityBlobEvidenceRoles_Destination       SecurityBlobEvidenceRoles = "destination"
	SecurityBlobEvidenceRoles_Edited            SecurityBlobEvidenceRoles = "edited"
	SecurityBlobEvidenceRoles_Loaded            SecurityBlobEvidenceRoles = "loaded"
	SecurityBlobEvidenceRoles_PolicyViolator    SecurityBlobEvidenceRoles = "policyViolator"
	SecurityBlobEvidenceRoles_Scanned           SecurityBlobEvidenceRoles = "scanned"
	SecurityBlobEvidenceRoles_Source            SecurityBlobEvidenceRoles = "source"
	SecurityBlobEvidenceRoles_Suspicious        SecurityBlobEvidenceRoles = "suspicious"
	SecurityBlobEvidenceRoles_Unknown           SecurityBlobEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityBlobEvidenceRoles() []string {
	return []string{
		string(SecurityBlobEvidenceRoles_Added),
		string(SecurityBlobEvidenceRoles_Attacked),
		string(SecurityBlobEvidenceRoles_Attacker),
		string(SecurityBlobEvidenceRoles_CommandAndControl),
		string(SecurityBlobEvidenceRoles_Compromised),
		string(SecurityBlobEvidenceRoles_Contextual),
		string(SecurityBlobEvidenceRoles_Created),
		string(SecurityBlobEvidenceRoles_Destination),
		string(SecurityBlobEvidenceRoles_Edited),
		string(SecurityBlobEvidenceRoles_Loaded),
		string(SecurityBlobEvidenceRoles_PolicyViolator),
		string(SecurityBlobEvidenceRoles_Scanned),
		string(SecurityBlobEvidenceRoles_Source),
		string(SecurityBlobEvidenceRoles_Suspicious),
		string(SecurityBlobEvidenceRoles_Unknown),
	}
}

func (s *SecurityBlobEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobEvidenceRoles(input string) (*SecurityBlobEvidenceRoles, error) {
	vals := map[string]SecurityBlobEvidenceRoles{
		"added":             SecurityBlobEvidenceRoles_Added,
		"attacked":          SecurityBlobEvidenceRoles_Attacked,
		"attacker":          SecurityBlobEvidenceRoles_Attacker,
		"commandandcontrol": SecurityBlobEvidenceRoles_CommandAndControl,
		"compromised":       SecurityBlobEvidenceRoles_Compromised,
		"contextual":        SecurityBlobEvidenceRoles_Contextual,
		"created":           SecurityBlobEvidenceRoles_Created,
		"destination":       SecurityBlobEvidenceRoles_Destination,
		"edited":            SecurityBlobEvidenceRoles_Edited,
		"loaded":            SecurityBlobEvidenceRoles_Loaded,
		"policyviolator":    SecurityBlobEvidenceRoles_PolicyViolator,
		"scanned":           SecurityBlobEvidenceRoles_Scanned,
		"source":            SecurityBlobEvidenceRoles_Source,
		"suspicious":        SecurityBlobEvidenceRoles_Suspicious,
		"unknown":           SecurityBlobEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobEvidenceRoles(input)
	return &out, nil
}
