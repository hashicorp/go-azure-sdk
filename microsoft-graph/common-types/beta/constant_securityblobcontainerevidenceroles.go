package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceRoles string

const (
	SecurityBlobContainerEvidenceRoles_Added             SecurityBlobContainerEvidenceRoles = "added"
	SecurityBlobContainerEvidenceRoles_Attacked          SecurityBlobContainerEvidenceRoles = "attacked"
	SecurityBlobContainerEvidenceRoles_Attacker          SecurityBlobContainerEvidenceRoles = "attacker"
	SecurityBlobContainerEvidenceRoles_CommandAndControl SecurityBlobContainerEvidenceRoles = "commandAndControl"
	SecurityBlobContainerEvidenceRoles_Compromised       SecurityBlobContainerEvidenceRoles = "compromised"
	SecurityBlobContainerEvidenceRoles_Contextual        SecurityBlobContainerEvidenceRoles = "contextual"
	SecurityBlobContainerEvidenceRoles_Created           SecurityBlobContainerEvidenceRoles = "created"
	SecurityBlobContainerEvidenceRoles_Destination       SecurityBlobContainerEvidenceRoles = "destination"
	SecurityBlobContainerEvidenceRoles_Edited            SecurityBlobContainerEvidenceRoles = "edited"
	SecurityBlobContainerEvidenceRoles_Loaded            SecurityBlobContainerEvidenceRoles = "loaded"
	SecurityBlobContainerEvidenceRoles_PolicyViolator    SecurityBlobContainerEvidenceRoles = "policyViolator"
	SecurityBlobContainerEvidenceRoles_Scanned           SecurityBlobContainerEvidenceRoles = "scanned"
	SecurityBlobContainerEvidenceRoles_Source            SecurityBlobContainerEvidenceRoles = "source"
	SecurityBlobContainerEvidenceRoles_Suspicious        SecurityBlobContainerEvidenceRoles = "suspicious"
	SecurityBlobContainerEvidenceRoles_Unknown           SecurityBlobContainerEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityBlobContainerEvidenceRoles() []string {
	return []string{
		string(SecurityBlobContainerEvidenceRoles_Added),
		string(SecurityBlobContainerEvidenceRoles_Attacked),
		string(SecurityBlobContainerEvidenceRoles_Attacker),
		string(SecurityBlobContainerEvidenceRoles_CommandAndControl),
		string(SecurityBlobContainerEvidenceRoles_Compromised),
		string(SecurityBlobContainerEvidenceRoles_Contextual),
		string(SecurityBlobContainerEvidenceRoles_Created),
		string(SecurityBlobContainerEvidenceRoles_Destination),
		string(SecurityBlobContainerEvidenceRoles_Edited),
		string(SecurityBlobContainerEvidenceRoles_Loaded),
		string(SecurityBlobContainerEvidenceRoles_PolicyViolator),
		string(SecurityBlobContainerEvidenceRoles_Scanned),
		string(SecurityBlobContainerEvidenceRoles_Source),
		string(SecurityBlobContainerEvidenceRoles_Suspicious),
		string(SecurityBlobContainerEvidenceRoles_Unknown),
	}
}

func (s *SecurityBlobContainerEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityBlobContainerEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityBlobContainerEvidenceRoles(input string) (*SecurityBlobContainerEvidenceRoles, error) {
	vals := map[string]SecurityBlobContainerEvidenceRoles{
		"added":             SecurityBlobContainerEvidenceRoles_Added,
		"attacked":          SecurityBlobContainerEvidenceRoles_Attacked,
		"attacker":          SecurityBlobContainerEvidenceRoles_Attacker,
		"commandandcontrol": SecurityBlobContainerEvidenceRoles_CommandAndControl,
		"compromised":       SecurityBlobContainerEvidenceRoles_Compromised,
		"contextual":        SecurityBlobContainerEvidenceRoles_Contextual,
		"created":           SecurityBlobContainerEvidenceRoles_Created,
		"destination":       SecurityBlobContainerEvidenceRoles_Destination,
		"edited":            SecurityBlobContainerEvidenceRoles_Edited,
		"loaded":            SecurityBlobContainerEvidenceRoles_Loaded,
		"policyviolator":    SecurityBlobContainerEvidenceRoles_PolicyViolator,
		"scanned":           SecurityBlobContainerEvidenceRoles_Scanned,
		"source":            SecurityBlobContainerEvidenceRoles_Source,
		"suspicious":        SecurityBlobContainerEvidenceRoles_Suspicious,
		"unknown":           SecurityBlobContainerEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceRoles(input)
	return &out, nil
}
