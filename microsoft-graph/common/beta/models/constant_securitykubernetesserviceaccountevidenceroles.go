package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesServiceAccountEvidenceRoles string

const (
	SecurityKubernetesServiceAccountEvidenceRolesadded             SecurityKubernetesServiceAccountEvidenceRoles = "Added"
	SecurityKubernetesServiceAccountEvidenceRolesattacked          SecurityKubernetesServiceAccountEvidenceRoles = "Attacked"
	SecurityKubernetesServiceAccountEvidenceRolesattacker          SecurityKubernetesServiceAccountEvidenceRoles = "Attacker"
	SecurityKubernetesServiceAccountEvidenceRolescommandAndControl SecurityKubernetesServiceAccountEvidenceRoles = "CommandAndControl"
	SecurityKubernetesServiceAccountEvidenceRolescompromised       SecurityKubernetesServiceAccountEvidenceRoles = "Compromised"
	SecurityKubernetesServiceAccountEvidenceRolescontextual        SecurityKubernetesServiceAccountEvidenceRoles = "Contextual"
	SecurityKubernetesServiceAccountEvidenceRolescreated           SecurityKubernetesServiceAccountEvidenceRoles = "Created"
	SecurityKubernetesServiceAccountEvidenceRolesdestination       SecurityKubernetesServiceAccountEvidenceRoles = "Destination"
	SecurityKubernetesServiceAccountEvidenceRolesedited            SecurityKubernetesServiceAccountEvidenceRoles = "Edited"
	SecurityKubernetesServiceAccountEvidenceRolesloaded            SecurityKubernetesServiceAccountEvidenceRoles = "Loaded"
	SecurityKubernetesServiceAccountEvidenceRolespolicyViolator    SecurityKubernetesServiceAccountEvidenceRoles = "PolicyViolator"
	SecurityKubernetesServiceAccountEvidenceRolesscanned           SecurityKubernetesServiceAccountEvidenceRoles = "Scanned"
	SecurityKubernetesServiceAccountEvidenceRolessource            SecurityKubernetesServiceAccountEvidenceRoles = "Source"
	SecurityKubernetesServiceAccountEvidenceRolessuspicious        SecurityKubernetesServiceAccountEvidenceRoles = "Suspicious"
	SecurityKubernetesServiceAccountEvidenceRolesunknown           SecurityKubernetesServiceAccountEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityKubernetesServiceAccountEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesServiceAccountEvidenceRolesadded),
		string(SecurityKubernetesServiceAccountEvidenceRolesattacked),
		string(SecurityKubernetesServiceAccountEvidenceRolesattacker),
		string(SecurityKubernetesServiceAccountEvidenceRolescommandAndControl),
		string(SecurityKubernetesServiceAccountEvidenceRolescompromised),
		string(SecurityKubernetesServiceAccountEvidenceRolescontextual),
		string(SecurityKubernetesServiceAccountEvidenceRolescreated),
		string(SecurityKubernetesServiceAccountEvidenceRolesdestination),
		string(SecurityKubernetesServiceAccountEvidenceRolesedited),
		string(SecurityKubernetesServiceAccountEvidenceRolesloaded),
		string(SecurityKubernetesServiceAccountEvidenceRolespolicyViolator),
		string(SecurityKubernetesServiceAccountEvidenceRolesscanned),
		string(SecurityKubernetesServiceAccountEvidenceRolessource),
		string(SecurityKubernetesServiceAccountEvidenceRolessuspicious),
		string(SecurityKubernetesServiceAccountEvidenceRolesunknown),
	}
}

func parseSecurityKubernetesServiceAccountEvidenceRoles(input string) (*SecurityKubernetesServiceAccountEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesServiceAccountEvidenceRoles{
		"added":             SecurityKubernetesServiceAccountEvidenceRolesadded,
		"attacked":          SecurityKubernetesServiceAccountEvidenceRolesattacked,
		"attacker":          SecurityKubernetesServiceAccountEvidenceRolesattacker,
		"commandandcontrol": SecurityKubernetesServiceAccountEvidenceRolescommandAndControl,
		"compromised":       SecurityKubernetesServiceAccountEvidenceRolescompromised,
		"contextual":        SecurityKubernetesServiceAccountEvidenceRolescontextual,
		"created":           SecurityKubernetesServiceAccountEvidenceRolescreated,
		"destination":       SecurityKubernetesServiceAccountEvidenceRolesdestination,
		"edited":            SecurityKubernetesServiceAccountEvidenceRolesedited,
		"loaded":            SecurityKubernetesServiceAccountEvidenceRolesloaded,
		"policyviolator":    SecurityKubernetesServiceAccountEvidenceRolespolicyViolator,
		"scanned":           SecurityKubernetesServiceAccountEvidenceRolesscanned,
		"source":            SecurityKubernetesServiceAccountEvidenceRolessource,
		"suspicious":        SecurityKubernetesServiceAccountEvidenceRolessuspicious,
		"unknown":           SecurityKubernetesServiceAccountEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesServiceAccountEvidenceRoles(input)
	return &out, nil
}
