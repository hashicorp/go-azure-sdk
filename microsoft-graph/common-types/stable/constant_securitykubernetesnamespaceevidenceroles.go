package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceRoles string

const (
	SecurityKubernetesNamespaceEvidenceRoles_Added             SecurityKubernetesNamespaceEvidenceRoles = "added"
	SecurityKubernetesNamespaceEvidenceRoles_Attacked          SecurityKubernetesNamespaceEvidenceRoles = "attacked"
	SecurityKubernetesNamespaceEvidenceRoles_Attacker          SecurityKubernetesNamespaceEvidenceRoles = "attacker"
	SecurityKubernetesNamespaceEvidenceRoles_CommandAndControl SecurityKubernetesNamespaceEvidenceRoles = "commandAndControl"
	SecurityKubernetesNamespaceEvidenceRoles_Compromised       SecurityKubernetesNamespaceEvidenceRoles = "compromised"
	SecurityKubernetesNamespaceEvidenceRoles_Contextual        SecurityKubernetesNamespaceEvidenceRoles = "contextual"
	SecurityKubernetesNamespaceEvidenceRoles_Created           SecurityKubernetesNamespaceEvidenceRoles = "created"
	SecurityKubernetesNamespaceEvidenceRoles_Destination       SecurityKubernetesNamespaceEvidenceRoles = "destination"
	SecurityKubernetesNamespaceEvidenceRoles_Edited            SecurityKubernetesNamespaceEvidenceRoles = "edited"
	SecurityKubernetesNamespaceEvidenceRoles_Loaded            SecurityKubernetesNamespaceEvidenceRoles = "loaded"
	SecurityKubernetesNamespaceEvidenceRoles_PolicyViolator    SecurityKubernetesNamespaceEvidenceRoles = "policyViolator"
	SecurityKubernetesNamespaceEvidenceRoles_Scanned           SecurityKubernetesNamespaceEvidenceRoles = "scanned"
	SecurityKubernetesNamespaceEvidenceRoles_Source            SecurityKubernetesNamespaceEvidenceRoles = "source"
	SecurityKubernetesNamespaceEvidenceRoles_Suspicious        SecurityKubernetesNamespaceEvidenceRoles = "suspicious"
	SecurityKubernetesNamespaceEvidenceRoles_Unknown           SecurityKubernetesNamespaceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceRoles_Added),
		string(SecurityKubernetesNamespaceEvidenceRoles_Attacked),
		string(SecurityKubernetesNamespaceEvidenceRoles_Attacker),
		string(SecurityKubernetesNamespaceEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesNamespaceEvidenceRoles_Compromised),
		string(SecurityKubernetesNamespaceEvidenceRoles_Contextual),
		string(SecurityKubernetesNamespaceEvidenceRoles_Created),
		string(SecurityKubernetesNamespaceEvidenceRoles_Destination),
		string(SecurityKubernetesNamespaceEvidenceRoles_Edited),
		string(SecurityKubernetesNamespaceEvidenceRoles_Loaded),
		string(SecurityKubernetesNamespaceEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesNamespaceEvidenceRoles_Scanned),
		string(SecurityKubernetesNamespaceEvidenceRoles_Source),
		string(SecurityKubernetesNamespaceEvidenceRoles_Suspicious),
		string(SecurityKubernetesNamespaceEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesNamespaceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesNamespaceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesNamespaceEvidenceRoles(input string) (*SecurityKubernetesNamespaceEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceRoles{
		"added":             SecurityKubernetesNamespaceEvidenceRoles_Added,
		"attacked":          SecurityKubernetesNamespaceEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesNamespaceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesNamespaceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesNamespaceEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesNamespaceEvidenceRoles_Contextual,
		"created":           SecurityKubernetesNamespaceEvidenceRoles_Created,
		"destination":       SecurityKubernetesNamespaceEvidenceRoles_Destination,
		"edited":            SecurityKubernetesNamespaceEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesNamespaceEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesNamespaceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesNamespaceEvidenceRoles_Scanned,
		"source":            SecurityKubernetesNamespaceEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesNamespaceEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesNamespaceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceRoles(input)
	return &out, nil
}
