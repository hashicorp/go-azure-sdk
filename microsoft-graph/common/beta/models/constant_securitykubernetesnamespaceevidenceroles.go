package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesNamespaceEvidenceRoles string

const (
	SecurityKubernetesNamespaceEvidenceRolesadded             SecurityKubernetesNamespaceEvidenceRoles = "Added"
	SecurityKubernetesNamespaceEvidenceRolesattacked          SecurityKubernetesNamespaceEvidenceRoles = "Attacked"
	SecurityKubernetesNamespaceEvidenceRolesattacker          SecurityKubernetesNamespaceEvidenceRoles = "Attacker"
	SecurityKubernetesNamespaceEvidenceRolescommandAndControl SecurityKubernetesNamespaceEvidenceRoles = "CommandAndControl"
	SecurityKubernetesNamespaceEvidenceRolescompromised       SecurityKubernetesNamespaceEvidenceRoles = "Compromised"
	SecurityKubernetesNamespaceEvidenceRolescontextual        SecurityKubernetesNamespaceEvidenceRoles = "Contextual"
	SecurityKubernetesNamespaceEvidenceRolescreated           SecurityKubernetesNamespaceEvidenceRoles = "Created"
	SecurityKubernetesNamespaceEvidenceRolesdestination       SecurityKubernetesNamespaceEvidenceRoles = "Destination"
	SecurityKubernetesNamespaceEvidenceRolesedited            SecurityKubernetesNamespaceEvidenceRoles = "Edited"
	SecurityKubernetesNamespaceEvidenceRolesloaded            SecurityKubernetesNamespaceEvidenceRoles = "Loaded"
	SecurityKubernetesNamespaceEvidenceRolespolicyViolator    SecurityKubernetesNamespaceEvidenceRoles = "PolicyViolator"
	SecurityKubernetesNamespaceEvidenceRolesscanned           SecurityKubernetesNamespaceEvidenceRoles = "Scanned"
	SecurityKubernetesNamespaceEvidenceRolessource            SecurityKubernetesNamespaceEvidenceRoles = "Source"
	SecurityKubernetesNamespaceEvidenceRolessuspicious        SecurityKubernetesNamespaceEvidenceRoles = "Suspicious"
	SecurityKubernetesNamespaceEvidenceRolesunknown           SecurityKubernetesNamespaceEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityKubernetesNamespaceEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesNamespaceEvidenceRolesadded),
		string(SecurityKubernetesNamespaceEvidenceRolesattacked),
		string(SecurityKubernetesNamespaceEvidenceRolesattacker),
		string(SecurityKubernetesNamespaceEvidenceRolescommandAndControl),
		string(SecurityKubernetesNamespaceEvidenceRolescompromised),
		string(SecurityKubernetesNamespaceEvidenceRolescontextual),
		string(SecurityKubernetesNamespaceEvidenceRolescreated),
		string(SecurityKubernetesNamespaceEvidenceRolesdestination),
		string(SecurityKubernetesNamespaceEvidenceRolesedited),
		string(SecurityKubernetesNamespaceEvidenceRolesloaded),
		string(SecurityKubernetesNamespaceEvidenceRolespolicyViolator),
		string(SecurityKubernetesNamespaceEvidenceRolesscanned),
		string(SecurityKubernetesNamespaceEvidenceRolessource),
		string(SecurityKubernetesNamespaceEvidenceRolessuspicious),
		string(SecurityKubernetesNamespaceEvidenceRolesunknown),
	}
}

func parseSecurityKubernetesNamespaceEvidenceRoles(input string) (*SecurityKubernetesNamespaceEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesNamespaceEvidenceRoles{
		"added":             SecurityKubernetesNamespaceEvidenceRolesadded,
		"attacked":          SecurityKubernetesNamespaceEvidenceRolesattacked,
		"attacker":          SecurityKubernetesNamespaceEvidenceRolesattacker,
		"commandandcontrol": SecurityKubernetesNamespaceEvidenceRolescommandAndControl,
		"compromised":       SecurityKubernetesNamespaceEvidenceRolescompromised,
		"contextual":        SecurityKubernetesNamespaceEvidenceRolescontextual,
		"created":           SecurityKubernetesNamespaceEvidenceRolescreated,
		"destination":       SecurityKubernetesNamespaceEvidenceRolesdestination,
		"edited":            SecurityKubernetesNamespaceEvidenceRolesedited,
		"loaded":            SecurityKubernetesNamespaceEvidenceRolesloaded,
		"policyviolator":    SecurityKubernetesNamespaceEvidenceRolespolicyViolator,
		"scanned":           SecurityKubernetesNamespaceEvidenceRolesscanned,
		"source":            SecurityKubernetesNamespaceEvidenceRolessource,
		"suspicious":        SecurityKubernetesNamespaceEvidenceRolessuspicious,
		"unknown":           SecurityKubernetesNamespaceEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesNamespaceEvidenceRoles(input)
	return &out, nil
}
