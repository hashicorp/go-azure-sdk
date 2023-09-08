package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesClusterEvidenceRoles string

const (
	SecurityKubernetesClusterEvidenceRolesadded             SecurityKubernetesClusterEvidenceRoles = "Added"
	SecurityKubernetesClusterEvidenceRolesattacked          SecurityKubernetesClusterEvidenceRoles = "Attacked"
	SecurityKubernetesClusterEvidenceRolesattacker          SecurityKubernetesClusterEvidenceRoles = "Attacker"
	SecurityKubernetesClusterEvidenceRolescommandAndControl SecurityKubernetesClusterEvidenceRoles = "CommandAndControl"
	SecurityKubernetesClusterEvidenceRolescompromised       SecurityKubernetesClusterEvidenceRoles = "Compromised"
	SecurityKubernetesClusterEvidenceRolescontextual        SecurityKubernetesClusterEvidenceRoles = "Contextual"
	SecurityKubernetesClusterEvidenceRolescreated           SecurityKubernetesClusterEvidenceRoles = "Created"
	SecurityKubernetesClusterEvidenceRolesdestination       SecurityKubernetesClusterEvidenceRoles = "Destination"
	SecurityKubernetesClusterEvidenceRolesedited            SecurityKubernetesClusterEvidenceRoles = "Edited"
	SecurityKubernetesClusterEvidenceRolesloaded            SecurityKubernetesClusterEvidenceRoles = "Loaded"
	SecurityKubernetesClusterEvidenceRolespolicyViolator    SecurityKubernetesClusterEvidenceRoles = "PolicyViolator"
	SecurityKubernetesClusterEvidenceRolesscanned           SecurityKubernetesClusterEvidenceRoles = "Scanned"
	SecurityKubernetesClusterEvidenceRolessource            SecurityKubernetesClusterEvidenceRoles = "Source"
	SecurityKubernetesClusterEvidenceRolessuspicious        SecurityKubernetesClusterEvidenceRoles = "Suspicious"
	SecurityKubernetesClusterEvidenceRolesunknown           SecurityKubernetesClusterEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityKubernetesClusterEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesClusterEvidenceRolesadded),
		string(SecurityKubernetesClusterEvidenceRolesattacked),
		string(SecurityKubernetesClusterEvidenceRolesattacker),
		string(SecurityKubernetesClusterEvidenceRolescommandAndControl),
		string(SecurityKubernetesClusterEvidenceRolescompromised),
		string(SecurityKubernetesClusterEvidenceRolescontextual),
		string(SecurityKubernetesClusterEvidenceRolescreated),
		string(SecurityKubernetesClusterEvidenceRolesdestination),
		string(SecurityKubernetesClusterEvidenceRolesedited),
		string(SecurityKubernetesClusterEvidenceRolesloaded),
		string(SecurityKubernetesClusterEvidenceRolespolicyViolator),
		string(SecurityKubernetesClusterEvidenceRolesscanned),
		string(SecurityKubernetesClusterEvidenceRolessource),
		string(SecurityKubernetesClusterEvidenceRolessuspicious),
		string(SecurityKubernetesClusterEvidenceRolesunknown),
	}
}

func parseSecurityKubernetesClusterEvidenceRoles(input string) (*SecurityKubernetesClusterEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesClusterEvidenceRoles{
		"added":             SecurityKubernetesClusterEvidenceRolesadded,
		"attacked":          SecurityKubernetesClusterEvidenceRolesattacked,
		"attacker":          SecurityKubernetesClusterEvidenceRolesattacker,
		"commandandcontrol": SecurityKubernetesClusterEvidenceRolescommandAndControl,
		"compromised":       SecurityKubernetesClusterEvidenceRolescompromised,
		"contextual":        SecurityKubernetesClusterEvidenceRolescontextual,
		"created":           SecurityKubernetesClusterEvidenceRolescreated,
		"destination":       SecurityKubernetesClusterEvidenceRolesdestination,
		"edited":            SecurityKubernetesClusterEvidenceRolesedited,
		"loaded":            SecurityKubernetesClusterEvidenceRolesloaded,
		"policyviolator":    SecurityKubernetesClusterEvidenceRolespolicyViolator,
		"scanned":           SecurityKubernetesClusterEvidenceRolesscanned,
		"source":            SecurityKubernetesClusterEvidenceRolessource,
		"suspicious":        SecurityKubernetesClusterEvidenceRolessuspicious,
		"unknown":           SecurityKubernetesClusterEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesClusterEvidenceRoles(input)
	return &out, nil
}
