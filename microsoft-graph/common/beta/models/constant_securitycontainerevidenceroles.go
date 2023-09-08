package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerEvidenceRoles string

const (
	SecurityContainerEvidenceRolesadded             SecurityContainerEvidenceRoles = "Added"
	SecurityContainerEvidenceRolesattacked          SecurityContainerEvidenceRoles = "Attacked"
	SecurityContainerEvidenceRolesattacker          SecurityContainerEvidenceRoles = "Attacker"
	SecurityContainerEvidenceRolescommandAndControl SecurityContainerEvidenceRoles = "CommandAndControl"
	SecurityContainerEvidenceRolescompromised       SecurityContainerEvidenceRoles = "Compromised"
	SecurityContainerEvidenceRolescontextual        SecurityContainerEvidenceRoles = "Contextual"
	SecurityContainerEvidenceRolescreated           SecurityContainerEvidenceRoles = "Created"
	SecurityContainerEvidenceRolesdestination       SecurityContainerEvidenceRoles = "Destination"
	SecurityContainerEvidenceRolesedited            SecurityContainerEvidenceRoles = "Edited"
	SecurityContainerEvidenceRolesloaded            SecurityContainerEvidenceRoles = "Loaded"
	SecurityContainerEvidenceRolespolicyViolator    SecurityContainerEvidenceRoles = "PolicyViolator"
	SecurityContainerEvidenceRolesscanned           SecurityContainerEvidenceRoles = "Scanned"
	SecurityContainerEvidenceRolessource            SecurityContainerEvidenceRoles = "Source"
	SecurityContainerEvidenceRolessuspicious        SecurityContainerEvidenceRoles = "Suspicious"
	SecurityContainerEvidenceRolesunknown           SecurityContainerEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityContainerEvidenceRoles() []string {
	return []string{
		string(SecurityContainerEvidenceRolesadded),
		string(SecurityContainerEvidenceRolesattacked),
		string(SecurityContainerEvidenceRolesattacker),
		string(SecurityContainerEvidenceRolescommandAndControl),
		string(SecurityContainerEvidenceRolescompromised),
		string(SecurityContainerEvidenceRolescontextual),
		string(SecurityContainerEvidenceRolescreated),
		string(SecurityContainerEvidenceRolesdestination),
		string(SecurityContainerEvidenceRolesedited),
		string(SecurityContainerEvidenceRolesloaded),
		string(SecurityContainerEvidenceRolespolicyViolator),
		string(SecurityContainerEvidenceRolesscanned),
		string(SecurityContainerEvidenceRolessource),
		string(SecurityContainerEvidenceRolessuspicious),
		string(SecurityContainerEvidenceRolesunknown),
	}
}

func parseSecurityContainerEvidenceRoles(input string) (*SecurityContainerEvidenceRoles, error) {
	vals := map[string]SecurityContainerEvidenceRoles{
		"added":             SecurityContainerEvidenceRolesadded,
		"attacked":          SecurityContainerEvidenceRolesattacked,
		"attacker":          SecurityContainerEvidenceRolesattacker,
		"commandandcontrol": SecurityContainerEvidenceRolescommandAndControl,
		"compromised":       SecurityContainerEvidenceRolescompromised,
		"contextual":        SecurityContainerEvidenceRolescontextual,
		"created":           SecurityContainerEvidenceRolescreated,
		"destination":       SecurityContainerEvidenceRolesdestination,
		"edited":            SecurityContainerEvidenceRolesedited,
		"loaded":            SecurityContainerEvidenceRolesloaded,
		"policyviolator":    SecurityContainerEvidenceRolespolicyViolator,
		"scanned":           SecurityContainerEvidenceRolesscanned,
		"source":            SecurityContainerEvidenceRolessource,
		"suspicious":        SecurityContainerEvidenceRolessuspicious,
		"unknown":           SecurityContainerEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerEvidenceRoles(input)
	return &out, nil
}
