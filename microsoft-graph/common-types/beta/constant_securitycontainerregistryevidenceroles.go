package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidenceRoles string

const (
	SecurityContainerRegistryEvidenceRoles_Added             SecurityContainerRegistryEvidenceRoles = "added"
	SecurityContainerRegistryEvidenceRoles_Attacked          SecurityContainerRegistryEvidenceRoles = "attacked"
	SecurityContainerRegistryEvidenceRoles_Attacker          SecurityContainerRegistryEvidenceRoles = "attacker"
	SecurityContainerRegistryEvidenceRoles_CommandAndControl SecurityContainerRegistryEvidenceRoles = "commandAndControl"
	SecurityContainerRegistryEvidenceRoles_Compromised       SecurityContainerRegistryEvidenceRoles = "compromised"
	SecurityContainerRegistryEvidenceRoles_Contextual        SecurityContainerRegistryEvidenceRoles = "contextual"
	SecurityContainerRegistryEvidenceRoles_Created           SecurityContainerRegistryEvidenceRoles = "created"
	SecurityContainerRegistryEvidenceRoles_Destination       SecurityContainerRegistryEvidenceRoles = "destination"
	SecurityContainerRegistryEvidenceRoles_Edited            SecurityContainerRegistryEvidenceRoles = "edited"
	SecurityContainerRegistryEvidenceRoles_Loaded            SecurityContainerRegistryEvidenceRoles = "loaded"
	SecurityContainerRegistryEvidenceRoles_PolicyViolator    SecurityContainerRegistryEvidenceRoles = "policyViolator"
	SecurityContainerRegistryEvidenceRoles_Scanned           SecurityContainerRegistryEvidenceRoles = "scanned"
	SecurityContainerRegistryEvidenceRoles_Source            SecurityContainerRegistryEvidenceRoles = "source"
	SecurityContainerRegistryEvidenceRoles_Suspicious        SecurityContainerRegistryEvidenceRoles = "suspicious"
	SecurityContainerRegistryEvidenceRoles_Unknown           SecurityContainerRegistryEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityContainerRegistryEvidenceRoles() []string {
	return []string{
		string(SecurityContainerRegistryEvidenceRoles_Added),
		string(SecurityContainerRegistryEvidenceRoles_Attacked),
		string(SecurityContainerRegistryEvidenceRoles_Attacker),
		string(SecurityContainerRegistryEvidenceRoles_CommandAndControl),
		string(SecurityContainerRegistryEvidenceRoles_Compromised),
		string(SecurityContainerRegistryEvidenceRoles_Contextual),
		string(SecurityContainerRegistryEvidenceRoles_Created),
		string(SecurityContainerRegistryEvidenceRoles_Destination),
		string(SecurityContainerRegistryEvidenceRoles_Edited),
		string(SecurityContainerRegistryEvidenceRoles_Loaded),
		string(SecurityContainerRegistryEvidenceRoles_PolicyViolator),
		string(SecurityContainerRegistryEvidenceRoles_Scanned),
		string(SecurityContainerRegistryEvidenceRoles_Source),
		string(SecurityContainerRegistryEvidenceRoles_Suspicious),
		string(SecurityContainerRegistryEvidenceRoles_Unknown),
	}
}

func (s *SecurityContainerRegistryEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerRegistryEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerRegistryEvidenceRoles(input string) (*SecurityContainerRegistryEvidenceRoles, error) {
	vals := map[string]SecurityContainerRegistryEvidenceRoles{
		"added":             SecurityContainerRegistryEvidenceRoles_Added,
		"attacked":          SecurityContainerRegistryEvidenceRoles_Attacked,
		"attacker":          SecurityContainerRegistryEvidenceRoles_Attacker,
		"commandandcontrol": SecurityContainerRegistryEvidenceRoles_CommandAndControl,
		"compromised":       SecurityContainerRegistryEvidenceRoles_Compromised,
		"contextual":        SecurityContainerRegistryEvidenceRoles_Contextual,
		"created":           SecurityContainerRegistryEvidenceRoles_Created,
		"destination":       SecurityContainerRegistryEvidenceRoles_Destination,
		"edited":            SecurityContainerRegistryEvidenceRoles_Edited,
		"loaded":            SecurityContainerRegistryEvidenceRoles_Loaded,
		"policyviolator":    SecurityContainerRegistryEvidenceRoles_PolicyViolator,
		"scanned":           SecurityContainerRegistryEvidenceRoles_Scanned,
		"source":            SecurityContainerRegistryEvidenceRoles_Source,
		"suspicious":        SecurityContainerRegistryEvidenceRoles_Suspicious,
		"unknown":           SecurityContainerRegistryEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerRegistryEvidenceRoles(input)
	return &out, nil
}
