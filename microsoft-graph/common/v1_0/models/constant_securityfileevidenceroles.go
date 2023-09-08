package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFileEvidenceRoles string

const (
	SecurityFileEvidenceRolesadded             SecurityFileEvidenceRoles = "Added"
	SecurityFileEvidenceRolesattacked          SecurityFileEvidenceRoles = "Attacked"
	SecurityFileEvidenceRolesattacker          SecurityFileEvidenceRoles = "Attacker"
	SecurityFileEvidenceRolescommandAndControl SecurityFileEvidenceRoles = "CommandAndControl"
	SecurityFileEvidenceRolescompromised       SecurityFileEvidenceRoles = "Compromised"
	SecurityFileEvidenceRolescontextual        SecurityFileEvidenceRoles = "Contextual"
	SecurityFileEvidenceRolescreated           SecurityFileEvidenceRoles = "Created"
	SecurityFileEvidenceRolesdestination       SecurityFileEvidenceRoles = "Destination"
	SecurityFileEvidenceRolesedited            SecurityFileEvidenceRoles = "Edited"
	SecurityFileEvidenceRolesloaded            SecurityFileEvidenceRoles = "Loaded"
	SecurityFileEvidenceRolespolicyViolator    SecurityFileEvidenceRoles = "PolicyViolator"
	SecurityFileEvidenceRolesscanned           SecurityFileEvidenceRoles = "Scanned"
	SecurityFileEvidenceRolessource            SecurityFileEvidenceRoles = "Source"
	SecurityFileEvidenceRolessuspicious        SecurityFileEvidenceRoles = "Suspicious"
	SecurityFileEvidenceRolesunknown           SecurityFileEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityFileEvidenceRoles() []string {
	return []string{
		string(SecurityFileEvidenceRolesadded),
		string(SecurityFileEvidenceRolesattacked),
		string(SecurityFileEvidenceRolesattacker),
		string(SecurityFileEvidenceRolescommandAndControl),
		string(SecurityFileEvidenceRolescompromised),
		string(SecurityFileEvidenceRolescontextual),
		string(SecurityFileEvidenceRolescreated),
		string(SecurityFileEvidenceRolesdestination),
		string(SecurityFileEvidenceRolesedited),
		string(SecurityFileEvidenceRolesloaded),
		string(SecurityFileEvidenceRolespolicyViolator),
		string(SecurityFileEvidenceRolesscanned),
		string(SecurityFileEvidenceRolessource),
		string(SecurityFileEvidenceRolessuspicious),
		string(SecurityFileEvidenceRolesunknown),
	}
}

func parseSecurityFileEvidenceRoles(input string) (*SecurityFileEvidenceRoles, error) {
	vals := map[string]SecurityFileEvidenceRoles{
		"added":             SecurityFileEvidenceRolesadded,
		"attacked":          SecurityFileEvidenceRolesattacked,
		"attacker":          SecurityFileEvidenceRolesattacker,
		"commandandcontrol": SecurityFileEvidenceRolescommandAndControl,
		"compromised":       SecurityFileEvidenceRolescompromised,
		"contextual":        SecurityFileEvidenceRolescontextual,
		"created":           SecurityFileEvidenceRolescreated,
		"destination":       SecurityFileEvidenceRolesdestination,
		"edited":            SecurityFileEvidenceRolesedited,
		"loaded":            SecurityFileEvidenceRolesloaded,
		"policyviolator":    SecurityFileEvidenceRolespolicyViolator,
		"scanned":           SecurityFileEvidenceRolesscanned,
		"source":            SecurityFileEvidenceRolessource,
		"suspicious":        SecurityFileEvidenceRolessuspicious,
		"unknown":           SecurityFileEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityFileEvidenceRoles(input)
	return &out, nil
}
