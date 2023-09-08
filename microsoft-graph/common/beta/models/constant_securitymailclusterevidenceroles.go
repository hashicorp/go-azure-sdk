package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailClusterEvidenceRoles string

const (
	SecurityMailClusterEvidenceRolesadded             SecurityMailClusterEvidenceRoles = "Added"
	SecurityMailClusterEvidenceRolesattacked          SecurityMailClusterEvidenceRoles = "Attacked"
	SecurityMailClusterEvidenceRolesattacker          SecurityMailClusterEvidenceRoles = "Attacker"
	SecurityMailClusterEvidenceRolescommandAndControl SecurityMailClusterEvidenceRoles = "CommandAndControl"
	SecurityMailClusterEvidenceRolescompromised       SecurityMailClusterEvidenceRoles = "Compromised"
	SecurityMailClusterEvidenceRolescontextual        SecurityMailClusterEvidenceRoles = "Contextual"
	SecurityMailClusterEvidenceRolescreated           SecurityMailClusterEvidenceRoles = "Created"
	SecurityMailClusterEvidenceRolesdestination       SecurityMailClusterEvidenceRoles = "Destination"
	SecurityMailClusterEvidenceRolesedited            SecurityMailClusterEvidenceRoles = "Edited"
	SecurityMailClusterEvidenceRolesloaded            SecurityMailClusterEvidenceRoles = "Loaded"
	SecurityMailClusterEvidenceRolespolicyViolator    SecurityMailClusterEvidenceRoles = "PolicyViolator"
	SecurityMailClusterEvidenceRolesscanned           SecurityMailClusterEvidenceRoles = "Scanned"
	SecurityMailClusterEvidenceRolessource            SecurityMailClusterEvidenceRoles = "Source"
	SecurityMailClusterEvidenceRolessuspicious        SecurityMailClusterEvidenceRoles = "Suspicious"
	SecurityMailClusterEvidenceRolesunknown           SecurityMailClusterEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityMailClusterEvidenceRoles() []string {
	return []string{
		string(SecurityMailClusterEvidenceRolesadded),
		string(SecurityMailClusterEvidenceRolesattacked),
		string(SecurityMailClusterEvidenceRolesattacker),
		string(SecurityMailClusterEvidenceRolescommandAndControl),
		string(SecurityMailClusterEvidenceRolescompromised),
		string(SecurityMailClusterEvidenceRolescontextual),
		string(SecurityMailClusterEvidenceRolescreated),
		string(SecurityMailClusterEvidenceRolesdestination),
		string(SecurityMailClusterEvidenceRolesedited),
		string(SecurityMailClusterEvidenceRolesloaded),
		string(SecurityMailClusterEvidenceRolespolicyViolator),
		string(SecurityMailClusterEvidenceRolesscanned),
		string(SecurityMailClusterEvidenceRolessource),
		string(SecurityMailClusterEvidenceRolessuspicious),
		string(SecurityMailClusterEvidenceRolesunknown),
	}
}

func parseSecurityMailClusterEvidenceRoles(input string) (*SecurityMailClusterEvidenceRoles, error) {
	vals := map[string]SecurityMailClusterEvidenceRoles{
		"added":             SecurityMailClusterEvidenceRolesadded,
		"attacked":          SecurityMailClusterEvidenceRolesattacked,
		"attacker":          SecurityMailClusterEvidenceRolesattacker,
		"commandandcontrol": SecurityMailClusterEvidenceRolescommandAndControl,
		"compromised":       SecurityMailClusterEvidenceRolescompromised,
		"contextual":        SecurityMailClusterEvidenceRolescontextual,
		"created":           SecurityMailClusterEvidenceRolescreated,
		"destination":       SecurityMailClusterEvidenceRolesdestination,
		"edited":            SecurityMailClusterEvidenceRolesedited,
		"loaded":            SecurityMailClusterEvidenceRolesloaded,
		"policyviolator":    SecurityMailClusterEvidenceRolespolicyViolator,
		"scanned":           SecurityMailClusterEvidenceRolesscanned,
		"source":            SecurityMailClusterEvidenceRolessource,
		"suspicious":        SecurityMailClusterEvidenceRolessuspicious,
		"unknown":           SecurityMailClusterEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailClusterEvidenceRoles(input)
	return &out, nil
}
