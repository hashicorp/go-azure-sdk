package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceRoles string

const (
	SecurityMailboxEvidenceRoles_Added             SecurityMailboxEvidenceRoles = "added"
	SecurityMailboxEvidenceRoles_Attacked          SecurityMailboxEvidenceRoles = "attacked"
	SecurityMailboxEvidenceRoles_Attacker          SecurityMailboxEvidenceRoles = "attacker"
	SecurityMailboxEvidenceRoles_CommandAndControl SecurityMailboxEvidenceRoles = "commandAndControl"
	SecurityMailboxEvidenceRoles_Compromised       SecurityMailboxEvidenceRoles = "compromised"
	SecurityMailboxEvidenceRoles_Contextual        SecurityMailboxEvidenceRoles = "contextual"
	SecurityMailboxEvidenceRoles_Created           SecurityMailboxEvidenceRoles = "created"
	SecurityMailboxEvidenceRoles_Destination       SecurityMailboxEvidenceRoles = "destination"
	SecurityMailboxEvidenceRoles_Edited            SecurityMailboxEvidenceRoles = "edited"
	SecurityMailboxEvidenceRoles_Loaded            SecurityMailboxEvidenceRoles = "loaded"
	SecurityMailboxEvidenceRoles_PolicyViolator    SecurityMailboxEvidenceRoles = "policyViolator"
	SecurityMailboxEvidenceRoles_Scanned           SecurityMailboxEvidenceRoles = "scanned"
	SecurityMailboxEvidenceRoles_Source            SecurityMailboxEvidenceRoles = "source"
	SecurityMailboxEvidenceRoles_Suspicious        SecurityMailboxEvidenceRoles = "suspicious"
	SecurityMailboxEvidenceRoles_Unknown           SecurityMailboxEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityMailboxEvidenceRoles() []string {
	return []string{
		string(SecurityMailboxEvidenceRoles_Added),
		string(SecurityMailboxEvidenceRoles_Attacked),
		string(SecurityMailboxEvidenceRoles_Attacker),
		string(SecurityMailboxEvidenceRoles_CommandAndControl),
		string(SecurityMailboxEvidenceRoles_Compromised),
		string(SecurityMailboxEvidenceRoles_Contextual),
		string(SecurityMailboxEvidenceRoles_Created),
		string(SecurityMailboxEvidenceRoles_Destination),
		string(SecurityMailboxEvidenceRoles_Edited),
		string(SecurityMailboxEvidenceRoles_Loaded),
		string(SecurityMailboxEvidenceRoles_PolicyViolator),
		string(SecurityMailboxEvidenceRoles_Scanned),
		string(SecurityMailboxEvidenceRoles_Source),
		string(SecurityMailboxEvidenceRoles_Suspicious),
		string(SecurityMailboxEvidenceRoles_Unknown),
	}
}

func (s *SecurityMailboxEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityMailboxEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityMailboxEvidenceRoles(input string) (*SecurityMailboxEvidenceRoles, error) {
	vals := map[string]SecurityMailboxEvidenceRoles{
		"added":             SecurityMailboxEvidenceRoles_Added,
		"attacked":          SecurityMailboxEvidenceRoles_Attacked,
		"attacker":          SecurityMailboxEvidenceRoles_Attacker,
		"commandandcontrol": SecurityMailboxEvidenceRoles_CommandAndControl,
		"compromised":       SecurityMailboxEvidenceRoles_Compromised,
		"contextual":        SecurityMailboxEvidenceRoles_Contextual,
		"created":           SecurityMailboxEvidenceRoles_Created,
		"destination":       SecurityMailboxEvidenceRoles_Destination,
		"edited":            SecurityMailboxEvidenceRoles_Edited,
		"loaded":            SecurityMailboxEvidenceRoles_Loaded,
		"policyviolator":    SecurityMailboxEvidenceRoles_PolicyViolator,
		"scanned":           SecurityMailboxEvidenceRoles_Scanned,
		"source":            SecurityMailboxEvidenceRoles_Source,
		"suspicious":        SecurityMailboxEvidenceRoles_Suspicious,
		"unknown":           SecurityMailboxEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxEvidenceRoles(input)
	return &out, nil
}
