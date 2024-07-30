package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceRoles string

const (
	SecurityContainerImageEvidenceRoles_Added             SecurityContainerImageEvidenceRoles = "added"
	SecurityContainerImageEvidenceRoles_Attacked          SecurityContainerImageEvidenceRoles = "attacked"
	SecurityContainerImageEvidenceRoles_Attacker          SecurityContainerImageEvidenceRoles = "attacker"
	SecurityContainerImageEvidenceRoles_CommandAndControl SecurityContainerImageEvidenceRoles = "commandAndControl"
	SecurityContainerImageEvidenceRoles_Compromised       SecurityContainerImageEvidenceRoles = "compromised"
	SecurityContainerImageEvidenceRoles_Contextual        SecurityContainerImageEvidenceRoles = "contextual"
	SecurityContainerImageEvidenceRoles_Created           SecurityContainerImageEvidenceRoles = "created"
	SecurityContainerImageEvidenceRoles_Destination       SecurityContainerImageEvidenceRoles = "destination"
	SecurityContainerImageEvidenceRoles_Edited            SecurityContainerImageEvidenceRoles = "edited"
	SecurityContainerImageEvidenceRoles_Loaded            SecurityContainerImageEvidenceRoles = "loaded"
	SecurityContainerImageEvidenceRoles_PolicyViolator    SecurityContainerImageEvidenceRoles = "policyViolator"
	SecurityContainerImageEvidenceRoles_Scanned           SecurityContainerImageEvidenceRoles = "scanned"
	SecurityContainerImageEvidenceRoles_Source            SecurityContainerImageEvidenceRoles = "source"
	SecurityContainerImageEvidenceRoles_Suspicious        SecurityContainerImageEvidenceRoles = "suspicious"
	SecurityContainerImageEvidenceRoles_Unknown           SecurityContainerImageEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityContainerImageEvidenceRoles() []string {
	return []string{
		string(SecurityContainerImageEvidenceRoles_Added),
		string(SecurityContainerImageEvidenceRoles_Attacked),
		string(SecurityContainerImageEvidenceRoles_Attacker),
		string(SecurityContainerImageEvidenceRoles_CommandAndControl),
		string(SecurityContainerImageEvidenceRoles_Compromised),
		string(SecurityContainerImageEvidenceRoles_Contextual),
		string(SecurityContainerImageEvidenceRoles_Created),
		string(SecurityContainerImageEvidenceRoles_Destination),
		string(SecurityContainerImageEvidenceRoles_Edited),
		string(SecurityContainerImageEvidenceRoles_Loaded),
		string(SecurityContainerImageEvidenceRoles_PolicyViolator),
		string(SecurityContainerImageEvidenceRoles_Scanned),
		string(SecurityContainerImageEvidenceRoles_Source),
		string(SecurityContainerImageEvidenceRoles_Suspicious),
		string(SecurityContainerImageEvidenceRoles_Unknown),
	}
}

func (s *SecurityContainerImageEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityContainerImageEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityContainerImageEvidenceRoles(input string) (*SecurityContainerImageEvidenceRoles, error) {
	vals := map[string]SecurityContainerImageEvidenceRoles{
		"added":             SecurityContainerImageEvidenceRoles_Added,
		"attacked":          SecurityContainerImageEvidenceRoles_Attacked,
		"attacker":          SecurityContainerImageEvidenceRoles_Attacker,
		"commandandcontrol": SecurityContainerImageEvidenceRoles_CommandAndControl,
		"compromised":       SecurityContainerImageEvidenceRoles_Compromised,
		"contextual":        SecurityContainerImageEvidenceRoles_Contextual,
		"created":           SecurityContainerImageEvidenceRoles_Created,
		"destination":       SecurityContainerImageEvidenceRoles_Destination,
		"edited":            SecurityContainerImageEvidenceRoles_Edited,
		"loaded":            SecurityContainerImageEvidenceRoles_Loaded,
		"policyviolator":    SecurityContainerImageEvidenceRoles_PolicyViolator,
		"scanned":           SecurityContainerImageEvidenceRoles_Scanned,
		"source":            SecurityContainerImageEvidenceRoles_Source,
		"suspicious":        SecurityContainerImageEvidenceRoles_Suspicious,
		"unknown":           SecurityContainerImageEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceRoles(input)
	return &out, nil
}
