package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceRoles string

const (
	SecurityAlertEvidenceRoles_Added             SecurityAlertEvidenceRoles = "added"
	SecurityAlertEvidenceRoles_Attacked          SecurityAlertEvidenceRoles = "attacked"
	SecurityAlertEvidenceRoles_Attacker          SecurityAlertEvidenceRoles = "attacker"
	SecurityAlertEvidenceRoles_CommandAndControl SecurityAlertEvidenceRoles = "commandAndControl"
	SecurityAlertEvidenceRoles_Compromised       SecurityAlertEvidenceRoles = "compromised"
	SecurityAlertEvidenceRoles_Contextual        SecurityAlertEvidenceRoles = "contextual"
	SecurityAlertEvidenceRoles_Created           SecurityAlertEvidenceRoles = "created"
	SecurityAlertEvidenceRoles_Destination       SecurityAlertEvidenceRoles = "destination"
	SecurityAlertEvidenceRoles_Edited            SecurityAlertEvidenceRoles = "edited"
	SecurityAlertEvidenceRoles_Loaded            SecurityAlertEvidenceRoles = "loaded"
	SecurityAlertEvidenceRoles_PolicyViolator    SecurityAlertEvidenceRoles = "policyViolator"
	SecurityAlertEvidenceRoles_Scanned           SecurityAlertEvidenceRoles = "scanned"
	SecurityAlertEvidenceRoles_Source            SecurityAlertEvidenceRoles = "source"
	SecurityAlertEvidenceRoles_Suspicious        SecurityAlertEvidenceRoles = "suspicious"
	SecurityAlertEvidenceRoles_Unknown           SecurityAlertEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityAlertEvidenceRoles() []string {
	return []string{
		string(SecurityAlertEvidenceRoles_Added),
		string(SecurityAlertEvidenceRoles_Attacked),
		string(SecurityAlertEvidenceRoles_Attacker),
		string(SecurityAlertEvidenceRoles_CommandAndControl),
		string(SecurityAlertEvidenceRoles_Compromised),
		string(SecurityAlertEvidenceRoles_Contextual),
		string(SecurityAlertEvidenceRoles_Created),
		string(SecurityAlertEvidenceRoles_Destination),
		string(SecurityAlertEvidenceRoles_Edited),
		string(SecurityAlertEvidenceRoles_Loaded),
		string(SecurityAlertEvidenceRoles_PolicyViolator),
		string(SecurityAlertEvidenceRoles_Scanned),
		string(SecurityAlertEvidenceRoles_Source),
		string(SecurityAlertEvidenceRoles_Suspicious),
		string(SecurityAlertEvidenceRoles_Unknown),
	}
}

func (s *SecurityAlertEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAlertEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAlertEvidenceRoles(input string) (*SecurityAlertEvidenceRoles, error) {
	vals := map[string]SecurityAlertEvidenceRoles{
		"added":             SecurityAlertEvidenceRoles_Added,
		"attacked":          SecurityAlertEvidenceRoles_Attacked,
		"attacker":          SecurityAlertEvidenceRoles_Attacker,
		"commandandcontrol": SecurityAlertEvidenceRoles_CommandAndControl,
		"compromised":       SecurityAlertEvidenceRoles_Compromised,
		"contextual":        SecurityAlertEvidenceRoles_Contextual,
		"created":           SecurityAlertEvidenceRoles_Created,
		"destination":       SecurityAlertEvidenceRoles_Destination,
		"edited":            SecurityAlertEvidenceRoles_Edited,
		"loaded":            SecurityAlertEvidenceRoles_Loaded,
		"policyviolator":    SecurityAlertEvidenceRoles_PolicyViolator,
		"scanned":           SecurityAlertEvidenceRoles_Scanned,
		"source":            SecurityAlertEvidenceRoles_Source,
		"suspicious":        SecurityAlertEvidenceRoles_Suspicious,
		"unknown":           SecurityAlertEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceRoles(input)
	return &out, nil
}
