package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIoTDeviceEvidenceRoles string

const (
	SecurityIoTDeviceEvidenceRoles_Added             SecurityIoTDeviceEvidenceRoles = "added"
	SecurityIoTDeviceEvidenceRoles_Attacked          SecurityIoTDeviceEvidenceRoles = "attacked"
	SecurityIoTDeviceEvidenceRoles_Attacker          SecurityIoTDeviceEvidenceRoles = "attacker"
	SecurityIoTDeviceEvidenceRoles_CommandAndControl SecurityIoTDeviceEvidenceRoles = "commandAndControl"
	SecurityIoTDeviceEvidenceRoles_Compromised       SecurityIoTDeviceEvidenceRoles = "compromised"
	SecurityIoTDeviceEvidenceRoles_Contextual        SecurityIoTDeviceEvidenceRoles = "contextual"
	SecurityIoTDeviceEvidenceRoles_Created           SecurityIoTDeviceEvidenceRoles = "created"
	SecurityIoTDeviceEvidenceRoles_Destination       SecurityIoTDeviceEvidenceRoles = "destination"
	SecurityIoTDeviceEvidenceRoles_Edited            SecurityIoTDeviceEvidenceRoles = "edited"
	SecurityIoTDeviceEvidenceRoles_Loaded            SecurityIoTDeviceEvidenceRoles = "loaded"
	SecurityIoTDeviceEvidenceRoles_PolicyViolator    SecurityIoTDeviceEvidenceRoles = "policyViolator"
	SecurityIoTDeviceEvidenceRoles_Scanned           SecurityIoTDeviceEvidenceRoles = "scanned"
	SecurityIoTDeviceEvidenceRoles_Source            SecurityIoTDeviceEvidenceRoles = "source"
	SecurityIoTDeviceEvidenceRoles_Suspicious        SecurityIoTDeviceEvidenceRoles = "suspicious"
	SecurityIoTDeviceEvidenceRoles_Unknown           SecurityIoTDeviceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityIoTDeviceEvidenceRoles() []string {
	return []string{
		string(SecurityIoTDeviceEvidenceRoles_Added),
		string(SecurityIoTDeviceEvidenceRoles_Attacked),
		string(SecurityIoTDeviceEvidenceRoles_Attacker),
		string(SecurityIoTDeviceEvidenceRoles_CommandAndControl),
		string(SecurityIoTDeviceEvidenceRoles_Compromised),
		string(SecurityIoTDeviceEvidenceRoles_Contextual),
		string(SecurityIoTDeviceEvidenceRoles_Created),
		string(SecurityIoTDeviceEvidenceRoles_Destination),
		string(SecurityIoTDeviceEvidenceRoles_Edited),
		string(SecurityIoTDeviceEvidenceRoles_Loaded),
		string(SecurityIoTDeviceEvidenceRoles_PolicyViolator),
		string(SecurityIoTDeviceEvidenceRoles_Scanned),
		string(SecurityIoTDeviceEvidenceRoles_Source),
		string(SecurityIoTDeviceEvidenceRoles_Suspicious),
		string(SecurityIoTDeviceEvidenceRoles_Unknown),
	}
}

func (s *SecurityIoTDeviceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityIoTDeviceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityIoTDeviceEvidenceRoles(input string) (*SecurityIoTDeviceEvidenceRoles, error) {
	vals := map[string]SecurityIoTDeviceEvidenceRoles{
		"added":             SecurityIoTDeviceEvidenceRoles_Added,
		"attacked":          SecurityIoTDeviceEvidenceRoles_Attacked,
		"attacker":          SecurityIoTDeviceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityIoTDeviceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityIoTDeviceEvidenceRoles_Compromised,
		"contextual":        SecurityIoTDeviceEvidenceRoles_Contextual,
		"created":           SecurityIoTDeviceEvidenceRoles_Created,
		"destination":       SecurityIoTDeviceEvidenceRoles_Destination,
		"edited":            SecurityIoTDeviceEvidenceRoles_Edited,
		"loaded":            SecurityIoTDeviceEvidenceRoles_Loaded,
		"policyviolator":    SecurityIoTDeviceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityIoTDeviceEvidenceRoles_Scanned,
		"source":            SecurityIoTDeviceEvidenceRoles_Source,
		"suspicious":        SecurityIoTDeviceEvidenceRoles_Suspicious,
		"unknown":           SecurityIoTDeviceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIoTDeviceEvidenceRoles(input)
	return &out, nil
}
