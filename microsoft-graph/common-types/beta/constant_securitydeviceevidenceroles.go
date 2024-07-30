package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRoles string

const (
	SecurityDeviceEvidenceRoles_Added             SecurityDeviceEvidenceRoles = "added"
	SecurityDeviceEvidenceRoles_Attacked          SecurityDeviceEvidenceRoles = "attacked"
	SecurityDeviceEvidenceRoles_Attacker          SecurityDeviceEvidenceRoles = "attacker"
	SecurityDeviceEvidenceRoles_CommandAndControl SecurityDeviceEvidenceRoles = "commandAndControl"
	SecurityDeviceEvidenceRoles_Compromised       SecurityDeviceEvidenceRoles = "compromised"
	SecurityDeviceEvidenceRoles_Contextual        SecurityDeviceEvidenceRoles = "contextual"
	SecurityDeviceEvidenceRoles_Created           SecurityDeviceEvidenceRoles = "created"
	SecurityDeviceEvidenceRoles_Destination       SecurityDeviceEvidenceRoles = "destination"
	SecurityDeviceEvidenceRoles_Edited            SecurityDeviceEvidenceRoles = "edited"
	SecurityDeviceEvidenceRoles_Loaded            SecurityDeviceEvidenceRoles = "loaded"
	SecurityDeviceEvidenceRoles_PolicyViolator    SecurityDeviceEvidenceRoles = "policyViolator"
	SecurityDeviceEvidenceRoles_Scanned           SecurityDeviceEvidenceRoles = "scanned"
	SecurityDeviceEvidenceRoles_Source            SecurityDeviceEvidenceRoles = "source"
	SecurityDeviceEvidenceRoles_Suspicious        SecurityDeviceEvidenceRoles = "suspicious"
	SecurityDeviceEvidenceRoles_Unknown           SecurityDeviceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityDeviceEvidenceRoles() []string {
	return []string{
		string(SecurityDeviceEvidenceRoles_Added),
		string(SecurityDeviceEvidenceRoles_Attacked),
		string(SecurityDeviceEvidenceRoles_Attacker),
		string(SecurityDeviceEvidenceRoles_CommandAndControl),
		string(SecurityDeviceEvidenceRoles_Compromised),
		string(SecurityDeviceEvidenceRoles_Contextual),
		string(SecurityDeviceEvidenceRoles_Created),
		string(SecurityDeviceEvidenceRoles_Destination),
		string(SecurityDeviceEvidenceRoles_Edited),
		string(SecurityDeviceEvidenceRoles_Loaded),
		string(SecurityDeviceEvidenceRoles_PolicyViolator),
		string(SecurityDeviceEvidenceRoles_Scanned),
		string(SecurityDeviceEvidenceRoles_Source),
		string(SecurityDeviceEvidenceRoles_Suspicious),
		string(SecurityDeviceEvidenceRoles_Unknown),
	}
}

func (s *SecurityDeviceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityDeviceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityDeviceEvidenceRoles(input string) (*SecurityDeviceEvidenceRoles, error) {
	vals := map[string]SecurityDeviceEvidenceRoles{
		"added":             SecurityDeviceEvidenceRoles_Added,
		"attacked":          SecurityDeviceEvidenceRoles_Attacked,
		"attacker":          SecurityDeviceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityDeviceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityDeviceEvidenceRoles_Compromised,
		"contextual":        SecurityDeviceEvidenceRoles_Contextual,
		"created":           SecurityDeviceEvidenceRoles_Created,
		"destination":       SecurityDeviceEvidenceRoles_Destination,
		"edited":            SecurityDeviceEvidenceRoles_Edited,
		"loaded":            SecurityDeviceEvidenceRoles_Loaded,
		"policyviolator":    SecurityDeviceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityDeviceEvidenceRoles_Scanned,
		"source":            SecurityDeviceEvidenceRoles_Source,
		"suspicious":        SecurityDeviceEvidenceRoles_Suspicious,
		"unknown":           SecurityDeviceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRoles(input)
	return &out, nil
}
