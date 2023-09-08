package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceEvidenceRoles string

const (
	SecurityKubernetesServiceEvidenceRolesadded             SecurityKubernetesServiceEvidenceRoles = "Added"
	SecurityKubernetesServiceEvidenceRolesattacked          SecurityKubernetesServiceEvidenceRoles = "Attacked"
	SecurityKubernetesServiceEvidenceRolesattacker          SecurityKubernetesServiceEvidenceRoles = "Attacker"
	SecurityKubernetesServiceEvidenceRolescommandAndControl SecurityKubernetesServiceEvidenceRoles = "CommandAndControl"
	SecurityKubernetesServiceEvidenceRolescompromised       SecurityKubernetesServiceEvidenceRoles = "Compromised"
	SecurityKubernetesServiceEvidenceRolescontextual        SecurityKubernetesServiceEvidenceRoles = "Contextual"
	SecurityKubernetesServiceEvidenceRolescreated           SecurityKubernetesServiceEvidenceRoles = "Created"
	SecurityKubernetesServiceEvidenceRolesdestination       SecurityKubernetesServiceEvidenceRoles = "Destination"
	SecurityKubernetesServiceEvidenceRolesedited            SecurityKubernetesServiceEvidenceRoles = "Edited"
	SecurityKubernetesServiceEvidenceRolesloaded            SecurityKubernetesServiceEvidenceRoles = "Loaded"
	SecurityKubernetesServiceEvidenceRolespolicyViolator    SecurityKubernetesServiceEvidenceRoles = "PolicyViolator"
	SecurityKubernetesServiceEvidenceRolesscanned           SecurityKubernetesServiceEvidenceRoles = "Scanned"
	SecurityKubernetesServiceEvidenceRolessource            SecurityKubernetesServiceEvidenceRoles = "Source"
	SecurityKubernetesServiceEvidenceRolessuspicious        SecurityKubernetesServiceEvidenceRoles = "Suspicious"
	SecurityKubernetesServiceEvidenceRolesunknown           SecurityKubernetesServiceEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityKubernetesServiceEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesServiceEvidenceRolesadded),
		string(SecurityKubernetesServiceEvidenceRolesattacked),
		string(SecurityKubernetesServiceEvidenceRolesattacker),
		string(SecurityKubernetesServiceEvidenceRolescommandAndControl),
		string(SecurityKubernetesServiceEvidenceRolescompromised),
		string(SecurityKubernetesServiceEvidenceRolescontextual),
		string(SecurityKubernetesServiceEvidenceRolescreated),
		string(SecurityKubernetesServiceEvidenceRolesdestination),
		string(SecurityKubernetesServiceEvidenceRolesedited),
		string(SecurityKubernetesServiceEvidenceRolesloaded),
		string(SecurityKubernetesServiceEvidenceRolespolicyViolator),
		string(SecurityKubernetesServiceEvidenceRolesscanned),
		string(SecurityKubernetesServiceEvidenceRolessource),
		string(SecurityKubernetesServiceEvidenceRolessuspicious),
		string(SecurityKubernetesServiceEvidenceRolesunknown),
	}
}

func parseSecurityKubernetesServiceEvidenceRoles(input string) (*SecurityKubernetesServiceEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesServiceEvidenceRoles{
		"added":             SecurityKubernetesServiceEvidenceRolesadded,
		"attacked":          SecurityKubernetesServiceEvidenceRolesattacked,
		"attacker":          SecurityKubernetesServiceEvidenceRolesattacker,
		"commandandcontrol": SecurityKubernetesServiceEvidenceRolescommandAndControl,
		"compromised":       SecurityKubernetesServiceEvidenceRolescompromised,
		"contextual":        SecurityKubernetesServiceEvidenceRolescontextual,
		"created":           SecurityKubernetesServiceEvidenceRolescreated,
		"destination":       SecurityKubernetesServiceEvidenceRolesdestination,
		"edited":            SecurityKubernetesServiceEvidenceRolesedited,
		"loaded":            SecurityKubernetesServiceEvidenceRolesloaded,
		"policyviolator":    SecurityKubernetesServiceEvidenceRolespolicyViolator,
		"scanned":           SecurityKubernetesServiceEvidenceRolesscanned,
		"source":            SecurityKubernetesServiceEvidenceRolessource,
		"suspicious":        SecurityKubernetesServiceEvidenceRolessuspicious,
		"unknown":           SecurityKubernetesServiceEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceEvidenceRoles(input)
	return &out, nil
}
