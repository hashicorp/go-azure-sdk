package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesControllerEvidenceRoles string

const (
	SecurityKubernetesControllerEvidenceRoles_Added             SecurityKubernetesControllerEvidenceRoles = "added"
	SecurityKubernetesControllerEvidenceRoles_Attacked          SecurityKubernetesControllerEvidenceRoles = "attacked"
	SecurityKubernetesControllerEvidenceRoles_Attacker          SecurityKubernetesControllerEvidenceRoles = "attacker"
	SecurityKubernetesControllerEvidenceRoles_CommandAndControl SecurityKubernetesControllerEvidenceRoles = "commandAndControl"
	SecurityKubernetesControllerEvidenceRoles_Compromised       SecurityKubernetesControllerEvidenceRoles = "compromised"
	SecurityKubernetesControllerEvidenceRoles_Contextual        SecurityKubernetesControllerEvidenceRoles = "contextual"
	SecurityKubernetesControllerEvidenceRoles_Created           SecurityKubernetesControllerEvidenceRoles = "created"
	SecurityKubernetesControllerEvidenceRoles_Destination       SecurityKubernetesControllerEvidenceRoles = "destination"
	SecurityKubernetesControllerEvidenceRoles_Edited            SecurityKubernetesControllerEvidenceRoles = "edited"
	SecurityKubernetesControllerEvidenceRoles_Loaded            SecurityKubernetesControllerEvidenceRoles = "loaded"
	SecurityKubernetesControllerEvidenceRoles_PolicyViolator    SecurityKubernetesControllerEvidenceRoles = "policyViolator"
	SecurityKubernetesControllerEvidenceRoles_Scanned           SecurityKubernetesControllerEvidenceRoles = "scanned"
	SecurityKubernetesControllerEvidenceRoles_Source            SecurityKubernetesControllerEvidenceRoles = "source"
	SecurityKubernetesControllerEvidenceRoles_Suspicious        SecurityKubernetesControllerEvidenceRoles = "suspicious"
	SecurityKubernetesControllerEvidenceRoles_Unknown           SecurityKubernetesControllerEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesControllerEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesControllerEvidenceRoles_Added),
		string(SecurityKubernetesControllerEvidenceRoles_Attacked),
		string(SecurityKubernetesControllerEvidenceRoles_Attacker),
		string(SecurityKubernetesControllerEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesControllerEvidenceRoles_Compromised),
		string(SecurityKubernetesControllerEvidenceRoles_Contextual),
		string(SecurityKubernetesControllerEvidenceRoles_Created),
		string(SecurityKubernetesControllerEvidenceRoles_Destination),
		string(SecurityKubernetesControllerEvidenceRoles_Edited),
		string(SecurityKubernetesControllerEvidenceRoles_Loaded),
		string(SecurityKubernetesControllerEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesControllerEvidenceRoles_Scanned),
		string(SecurityKubernetesControllerEvidenceRoles_Source),
		string(SecurityKubernetesControllerEvidenceRoles_Suspicious),
		string(SecurityKubernetesControllerEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesControllerEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesControllerEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesControllerEvidenceRoles(input string) (*SecurityKubernetesControllerEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesControllerEvidenceRoles{
		"added":             SecurityKubernetesControllerEvidenceRoles_Added,
		"attacked":          SecurityKubernetesControllerEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesControllerEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesControllerEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesControllerEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesControllerEvidenceRoles_Contextual,
		"created":           SecurityKubernetesControllerEvidenceRoles_Created,
		"destination":       SecurityKubernetesControllerEvidenceRoles_Destination,
		"edited":            SecurityKubernetesControllerEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesControllerEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesControllerEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesControllerEvidenceRoles_Scanned,
		"source":            SecurityKubernetesControllerEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesControllerEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesControllerEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesControllerEvidenceRoles(input)
	return &out, nil
}
