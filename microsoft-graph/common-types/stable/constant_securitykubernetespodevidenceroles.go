package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidenceRoles string

const (
	SecurityKubernetesPodEvidenceRoles_Added             SecurityKubernetesPodEvidenceRoles = "added"
	SecurityKubernetesPodEvidenceRoles_Attacked          SecurityKubernetesPodEvidenceRoles = "attacked"
	SecurityKubernetesPodEvidenceRoles_Attacker          SecurityKubernetesPodEvidenceRoles = "attacker"
	SecurityKubernetesPodEvidenceRoles_CommandAndControl SecurityKubernetesPodEvidenceRoles = "commandAndControl"
	SecurityKubernetesPodEvidenceRoles_Compromised       SecurityKubernetesPodEvidenceRoles = "compromised"
	SecurityKubernetesPodEvidenceRoles_Contextual        SecurityKubernetesPodEvidenceRoles = "contextual"
	SecurityKubernetesPodEvidenceRoles_Created           SecurityKubernetesPodEvidenceRoles = "created"
	SecurityKubernetesPodEvidenceRoles_Destination       SecurityKubernetesPodEvidenceRoles = "destination"
	SecurityKubernetesPodEvidenceRoles_Edited            SecurityKubernetesPodEvidenceRoles = "edited"
	SecurityKubernetesPodEvidenceRoles_Loaded            SecurityKubernetesPodEvidenceRoles = "loaded"
	SecurityKubernetesPodEvidenceRoles_PolicyViolator    SecurityKubernetesPodEvidenceRoles = "policyViolator"
	SecurityKubernetesPodEvidenceRoles_Scanned           SecurityKubernetesPodEvidenceRoles = "scanned"
	SecurityKubernetesPodEvidenceRoles_Source            SecurityKubernetesPodEvidenceRoles = "source"
	SecurityKubernetesPodEvidenceRoles_Suspicious        SecurityKubernetesPodEvidenceRoles = "suspicious"
	SecurityKubernetesPodEvidenceRoles_Unknown           SecurityKubernetesPodEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityKubernetesPodEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesPodEvidenceRoles_Added),
		string(SecurityKubernetesPodEvidenceRoles_Attacked),
		string(SecurityKubernetesPodEvidenceRoles_Attacker),
		string(SecurityKubernetesPodEvidenceRoles_CommandAndControl),
		string(SecurityKubernetesPodEvidenceRoles_Compromised),
		string(SecurityKubernetesPodEvidenceRoles_Contextual),
		string(SecurityKubernetesPodEvidenceRoles_Created),
		string(SecurityKubernetesPodEvidenceRoles_Destination),
		string(SecurityKubernetesPodEvidenceRoles_Edited),
		string(SecurityKubernetesPodEvidenceRoles_Loaded),
		string(SecurityKubernetesPodEvidenceRoles_PolicyViolator),
		string(SecurityKubernetesPodEvidenceRoles_Scanned),
		string(SecurityKubernetesPodEvidenceRoles_Source),
		string(SecurityKubernetesPodEvidenceRoles_Suspicious),
		string(SecurityKubernetesPodEvidenceRoles_Unknown),
	}
}

func (s *SecurityKubernetesPodEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityKubernetesPodEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityKubernetesPodEvidenceRoles(input string) (*SecurityKubernetesPodEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesPodEvidenceRoles{
		"added":             SecurityKubernetesPodEvidenceRoles_Added,
		"attacked":          SecurityKubernetesPodEvidenceRoles_Attacked,
		"attacker":          SecurityKubernetesPodEvidenceRoles_Attacker,
		"commandandcontrol": SecurityKubernetesPodEvidenceRoles_CommandAndControl,
		"compromised":       SecurityKubernetesPodEvidenceRoles_Compromised,
		"contextual":        SecurityKubernetesPodEvidenceRoles_Contextual,
		"created":           SecurityKubernetesPodEvidenceRoles_Created,
		"destination":       SecurityKubernetesPodEvidenceRoles_Destination,
		"edited":            SecurityKubernetesPodEvidenceRoles_Edited,
		"loaded":            SecurityKubernetesPodEvidenceRoles_Loaded,
		"policyviolator":    SecurityKubernetesPodEvidenceRoles_PolicyViolator,
		"scanned":           SecurityKubernetesPodEvidenceRoles_Scanned,
		"source":            SecurityKubernetesPodEvidenceRoles_Source,
		"suspicious":        SecurityKubernetesPodEvidenceRoles_Suspicious,
		"unknown":           SecurityKubernetesPodEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPodEvidenceRoles(input)
	return &out, nil
}
