package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityGoogleCloudResourceEvidenceRoles string

const (
	SecurityGoogleCloudResourceEvidenceRoles_Added             SecurityGoogleCloudResourceEvidenceRoles = "added"
	SecurityGoogleCloudResourceEvidenceRoles_Attacked          SecurityGoogleCloudResourceEvidenceRoles = "attacked"
	SecurityGoogleCloudResourceEvidenceRoles_Attacker          SecurityGoogleCloudResourceEvidenceRoles = "attacker"
	SecurityGoogleCloudResourceEvidenceRoles_CommandAndControl SecurityGoogleCloudResourceEvidenceRoles = "commandAndControl"
	SecurityGoogleCloudResourceEvidenceRoles_Compromised       SecurityGoogleCloudResourceEvidenceRoles = "compromised"
	SecurityGoogleCloudResourceEvidenceRoles_Contextual        SecurityGoogleCloudResourceEvidenceRoles = "contextual"
	SecurityGoogleCloudResourceEvidenceRoles_Created           SecurityGoogleCloudResourceEvidenceRoles = "created"
	SecurityGoogleCloudResourceEvidenceRoles_Destination       SecurityGoogleCloudResourceEvidenceRoles = "destination"
	SecurityGoogleCloudResourceEvidenceRoles_Edited            SecurityGoogleCloudResourceEvidenceRoles = "edited"
	SecurityGoogleCloudResourceEvidenceRoles_Loaded            SecurityGoogleCloudResourceEvidenceRoles = "loaded"
	SecurityGoogleCloudResourceEvidenceRoles_PolicyViolator    SecurityGoogleCloudResourceEvidenceRoles = "policyViolator"
	SecurityGoogleCloudResourceEvidenceRoles_Scanned           SecurityGoogleCloudResourceEvidenceRoles = "scanned"
	SecurityGoogleCloudResourceEvidenceRoles_Source            SecurityGoogleCloudResourceEvidenceRoles = "source"
	SecurityGoogleCloudResourceEvidenceRoles_Suspicious        SecurityGoogleCloudResourceEvidenceRoles = "suspicious"
	SecurityGoogleCloudResourceEvidenceRoles_Unknown           SecurityGoogleCloudResourceEvidenceRoles = "unknown"
)

func PossibleValuesForSecurityGoogleCloudResourceEvidenceRoles() []string {
	return []string{
		string(SecurityGoogleCloudResourceEvidenceRoles_Added),
		string(SecurityGoogleCloudResourceEvidenceRoles_Attacked),
		string(SecurityGoogleCloudResourceEvidenceRoles_Attacker),
		string(SecurityGoogleCloudResourceEvidenceRoles_CommandAndControl),
		string(SecurityGoogleCloudResourceEvidenceRoles_Compromised),
		string(SecurityGoogleCloudResourceEvidenceRoles_Contextual),
		string(SecurityGoogleCloudResourceEvidenceRoles_Created),
		string(SecurityGoogleCloudResourceEvidenceRoles_Destination),
		string(SecurityGoogleCloudResourceEvidenceRoles_Edited),
		string(SecurityGoogleCloudResourceEvidenceRoles_Loaded),
		string(SecurityGoogleCloudResourceEvidenceRoles_PolicyViolator),
		string(SecurityGoogleCloudResourceEvidenceRoles_Scanned),
		string(SecurityGoogleCloudResourceEvidenceRoles_Source),
		string(SecurityGoogleCloudResourceEvidenceRoles_Suspicious),
		string(SecurityGoogleCloudResourceEvidenceRoles_Unknown),
	}
}

func (s *SecurityGoogleCloudResourceEvidenceRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityGoogleCloudResourceEvidenceRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityGoogleCloudResourceEvidenceRoles(input string) (*SecurityGoogleCloudResourceEvidenceRoles, error) {
	vals := map[string]SecurityGoogleCloudResourceEvidenceRoles{
		"added":             SecurityGoogleCloudResourceEvidenceRoles_Added,
		"attacked":          SecurityGoogleCloudResourceEvidenceRoles_Attacked,
		"attacker":          SecurityGoogleCloudResourceEvidenceRoles_Attacker,
		"commandandcontrol": SecurityGoogleCloudResourceEvidenceRoles_CommandAndControl,
		"compromised":       SecurityGoogleCloudResourceEvidenceRoles_Compromised,
		"contextual":        SecurityGoogleCloudResourceEvidenceRoles_Contextual,
		"created":           SecurityGoogleCloudResourceEvidenceRoles_Created,
		"destination":       SecurityGoogleCloudResourceEvidenceRoles_Destination,
		"edited":            SecurityGoogleCloudResourceEvidenceRoles_Edited,
		"loaded":            SecurityGoogleCloudResourceEvidenceRoles_Loaded,
		"policyviolator":    SecurityGoogleCloudResourceEvidenceRoles_PolicyViolator,
		"scanned":           SecurityGoogleCloudResourceEvidenceRoles_Scanned,
		"source":            SecurityGoogleCloudResourceEvidenceRoles_Source,
		"suspicious":        SecurityGoogleCloudResourceEvidenceRoles_Suspicious,
		"unknown":           SecurityGoogleCloudResourceEvidenceRoles_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityGoogleCloudResourceEvidenceRoles(input)
	return &out, nil
}
