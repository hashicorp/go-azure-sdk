package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidenceRoles string

const (
	SecurityKubernetesClusterEvidenceRoles_Added             SecurityKubernetesClusterEvidenceRoles = "added"
	SecurityKubernetesClusterEvidenceRoles_Attacked          SecurityKubernetesClusterEvidenceRoles = "attacked"
	SecurityKubernetesClusterEvidenceRoles_Attacker          SecurityKubernetesClusterEvidenceRoles = "attacker"
	SecurityKubernetesClusterEvidenceRoles_CommandAndControl SecurityKubernetesClusterEvidenceRoles = "commandAndControl"
	SecurityKubernetesClusterEvidenceRoles_Compromised       SecurityKubernetesClusterEvidenceRoles = "compromised"
	SecurityKubernetesClusterEvidenceRoles_Contextual        SecurityKubernetesClusterEvidenceRoles = "contextual"
	SecurityKubernetesClusterEvidenceRoles_Created           SecurityKubernetesClusterEvidenceRoles = "created"
	SecurityKubernetesClusterEvidenceRoles_Destination       SecurityKubernetesClusterEvidenceRoles = "destination"
	SecurityKubernetesClusterEvidenceRoles_Edited            SecurityKubernetesClusterEvidenceRoles = "edited"
	SecurityKubernetesClusterEvidenceRoles_Loaded            SecurityKubernetesClusterEvidenceRoles = "loaded"
	SecurityKubernetesClusterEvidenceRoles_PolicyViolator    SecurityKubernetesClusterEvidenceRoles = "policyViolator"
	SecurityKubernetesClusterEvidenceRoles_Scanned           SecurityKubernetesClusterEvidenceRoles = "scanned"
	SecurityKubernetesClusterEvidenceRoles_Source            SecurityKubernetesClusterEvidenceRoles = "source"
	SecurityKubernetesClusterEvidenceRoles_Suspicious        SecurityKubernetesClusterEvidenceRoles = "suspicious"
	SecurityKubernetesClusterEvidenceRoles_Unknown           SecurityKubernetesClusterEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesClusterEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesClusterEvidenceRoles_Added),
		string(SecurityKubernetesClusterEvidenceRoles_Attacked),
		string(SecurityKubernetesClusterEvidenceRoles_Attacker),
		string(SecurityKubernetesClusterEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesClusterEvidenceRoles_Compromised),
		string(SecurityKubernetesClusterEvidenceRoles_Contextual),
		string(SecurityKubernetesClusterEvidenceRoles_Created),
		string(SecurityKubernetesClusterEvidenceRoles_Destination),
		string(SecurityKubernetesClusterEvidenceRoles_Edited),
		string(SecurityKubernetesClusterEvidenceRoles_Loaded),
		string(SecurityKubernetesClusterEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesClusterEvidenceRoles_Scanned),
		string(SecurityKubernetesClusterEvidenceRoles_Source),
		string(SecurityKubernetesClusterEvidenceRoles_Suspicious),
		string(SecurityKubernetesClusterEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesClusterEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesClusterEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesClusterEvidenceRoles(input string) (*SecurityKubernetesClusterEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesClusterEvidenceRoles{
		"added":             SecurityKubernetesClusterEvidenceRoles_Added,
		"attacked":          SecurityKubernetesClusterEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesClusterEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesClusterEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesClusterEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesClusterEvidenceRoles_Contextual,
		"created":           SecurityKubernetesClusterEvidenceRoles_Created,
		"destination":       SecurityKubernetesClusterEvidenceRoles_Destination,
		"edited":            SecurityKubernetesClusterEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesClusterEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesClusterEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesClusterEvidenceRoles_Scanned,
		"source":            SecurityKubernetesClusterEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesClusterEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesClusterEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidenceRoles(input)
	return &out, nil
}
