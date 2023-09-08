package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerRegistryEvidenceRoles string

const (
	SecurityContainerRegistryEvidenceRolesadded             SecurityContainerRegistryEvidenceRoles = "Added"
	SecurityContainerRegistryEvidenceRolesattacked          SecurityContainerRegistryEvidenceRoles = "Attacked"
	SecurityContainerRegistryEvidenceRolesattacker          SecurityContainerRegistryEvidenceRoles = "Attacker"
	SecurityContainerRegistryEvidenceRolescommandAndControl SecurityContainerRegistryEvidenceRoles = "CommandAndControl"
	SecurityContainerRegistryEvidenceRolescompromised       SecurityContainerRegistryEvidenceRoles = "Compromised"
	SecurityContainerRegistryEvidenceRolescontextual        SecurityContainerRegistryEvidenceRoles = "Contextual"
	SecurityContainerRegistryEvidenceRolescreated           SecurityContainerRegistryEvidenceRoles = "Created"
	SecurityContainerRegistryEvidenceRolesdestination       SecurityContainerRegistryEvidenceRoles = "Destination"
	SecurityContainerRegistryEvidenceRolesedited            SecurityContainerRegistryEvidenceRoles = "Edited"
	SecurityContainerRegistryEvidenceRolesloaded            SecurityContainerRegistryEvidenceRoles = "Loaded"
	SecurityContainerRegistryEvidenceRolespolicyViolator    SecurityContainerRegistryEvidenceRoles = "PolicyViolator"
	SecurityContainerRegistryEvidenceRolesscanned           SecurityContainerRegistryEvidenceRoles = "Scanned"
	SecurityContainerRegistryEvidenceRolessource            SecurityContainerRegistryEvidenceRoles = "Source"
	SecurityContainerRegistryEvidenceRolessuspicious        SecurityContainerRegistryEvidenceRoles = "Suspicious"
	SecurityContainerRegistryEvidenceRolesunknown           SecurityContainerRegistryEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityContainerRegistryEvidenceRoles() []string {
	return []string{
		string(SecurityContainerRegistryEvidenceRolesadded),
		string(SecurityContainerRegistryEvidenceRolesattacked),
		string(SecurityContainerRegistryEvidenceRolesattacker),
		string(SecurityContainerRegistryEvidenceRolescommandAndControl),
		string(SecurityContainerRegistryEvidenceRolescompromised),
		string(SecurityContainerRegistryEvidenceRolescontextual),
		string(SecurityContainerRegistryEvidenceRolescreated),
		string(SecurityContainerRegistryEvidenceRolesdestination),
		string(SecurityContainerRegistryEvidenceRolesedited),
		string(SecurityContainerRegistryEvidenceRolesloaded),
		string(SecurityContainerRegistryEvidenceRolespolicyViolator),
		string(SecurityContainerRegistryEvidenceRolesscanned),
		string(SecurityContainerRegistryEvidenceRolessource),
		string(SecurityContainerRegistryEvidenceRolessuspicious),
		string(SecurityContainerRegistryEvidenceRolesunknown),
	}
}

func parseSecurityContainerRegistryEvidenceRoles(input string) (*SecurityContainerRegistryEvidenceRoles, error) {
	vals := map[string]SecurityContainerRegistryEvidenceRoles{
		"added":             SecurityContainerRegistryEvidenceRolesadded,
		"attacked":          SecurityContainerRegistryEvidenceRolesattacked,
		"attacker":          SecurityContainerRegistryEvidenceRolesattacker,
		"commandandcontrol": SecurityContainerRegistryEvidenceRolescommandAndControl,
		"compromised":       SecurityContainerRegistryEvidenceRolescompromised,
		"contextual":        SecurityContainerRegistryEvidenceRolescontextual,
		"created":           SecurityContainerRegistryEvidenceRolescreated,
		"destination":       SecurityContainerRegistryEvidenceRolesdestination,
		"edited":            SecurityContainerRegistryEvidenceRolesedited,
		"loaded":            SecurityContainerRegistryEvidenceRolesloaded,
		"policyviolator":    SecurityContainerRegistryEvidenceRolespolicyViolator,
		"scanned":           SecurityContainerRegistryEvidenceRolesscanned,
		"source":            SecurityContainerRegistryEvidenceRolessource,
		"suspicious":        SecurityContainerRegistryEvidenceRolessuspicious,
		"unknown":           SecurityContainerRegistryEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerRegistryEvidenceRoles(input)
	return &out, nil
}
