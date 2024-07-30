package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidenceRoles string

const (
	SecurityUserEvidenceRoles_Added             SecurityUserEvidenceRoles = "added"
	SecurityUserEvidenceRoles_Attacked          SecurityUserEvidenceRoles = "attacked"
	SecurityUserEvidenceRoles_Attacker          SecurityUserEvidenceRoles = "attacker"
	SecurityUserEvidenceRoles_CommandAndControl SecurityUserEvidenceRoles = "commandAndControl"
	SecurityUserEvidenceRoles_Compromised       SecurityUserEvidenceRoles = "compromised"
	SecurityUserEvidenceRoles_Contextual        SecurityUserEvidenceRoles = "contextual"
	SecurityUserEvidenceRoles_Created           SecurityUserEvidenceRoles = "created"
	SecurityUserEvidenceRoles_Destination       SecurityUserEvidenceRoles = "destination"
	SecurityUserEvidenceRoles_Edited            SecurityUserEvidenceRoles = "edited"
	SecurityUserEvidenceRoles_Loaded            SecurityUserEvidenceRoles = "loaded"
	SecurityUserEvidenceRoles_PolicyViolator    SecurityUserEvidenceRoles = "policyViolator"
	SecurityUserEvidenceRoles_Scanned           SecurityUserEvidenceRoles = "scanned"
	SecurityUserEvidenceRoles_Source            SecurityUserEvidenceRoles = "source"
	SecurityUserEvidenceRoles_Suspicious        SecurityUserEvidenceRoles = "suspicious"
	SecurityUserEvidenceRoles_Unknown           SecurityUserEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityUserEvidenceRoles() []string {
	return []string{
		string(SecurityUserEvidenceRoles_Added),
		string(SecurityUserEvidenceRoles_Attacked),
		string(SecurityUserEvidenceRoles_Attacker),
		string(SecurityUserEvidenceRoles_CommandAndControl),
		string(SecurityUserEvidenceRoles_Compromised),
		string(SecurityUserEvidenceRoles_Contextual),
		string(SecurityUserEvidenceRoles_Created),
		string(SecurityUserEvidenceRoles_Destination),
		string(SecurityUserEvidenceRoles_Edited),
		string(SecurityUserEvidenceRoles_Loaded),
		string(SecurityUserEvidenceRoles_PolicyViolator),
		string(SecurityUserEvidenceRoles_Scanned),
		string(SecurityUserEvidenceRoles_Source),
		string(SecurityUserEvidenceRoles_Suspicious),
		string(SecurityUserEvidenceRoles_Unknown),
	}
}

func (s *SecurityUserEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityUserEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityUserEvidenceRoles(input string) (*SecurityUserEvidenceRoles, error) {
	vals := map[string]SecurityUserEvidenceRoles{
		"added":             SecurityUserEvidenceRoles_Added,
		"attacked":          SecurityUserEvidenceRoles_Attacked,
		"attacker":          SecurityUserEvidenceRoles_Attacker,
		"commandandcontrol": SecurityUserEvidenceRoles_CommandAndControl,
		"compromised":       SecurityUserEvidenceRoles_Compromised,
		"contextual":        SecurityUserEvidenceRoles_Contextual,
		"created":           SecurityUserEvidenceRoles_Created,
		"destination":       SecurityUserEvidenceRoles_Destination,
		"edited":            SecurityUserEvidenceRoles_Edited,
		"loaded":            SecurityUserEvidenceRoles_Loaded,
		"policyviolator":    SecurityUserEvidenceRoles_PolicyViolator,
		"scanned":           SecurityUserEvidenceRoles_Scanned,
		"source":            SecurityUserEvidenceRoles_Source,
		"suspicious":        SecurityUserEvidenceRoles_Suspicious,
		"unknown":           SecurityUserEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserEvidenceRoles(input)
	return &out, nil
}
