package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityProcessEvidenceRoles string

const (
	SecurityProcessEvidenceRolesadded             SecurityProcessEvidenceRoles = "Added"
	SecurityProcessEvidenceRolesattacked          SecurityProcessEvidenceRoles = "Attacked"
	SecurityProcessEvidenceRolesattacker          SecurityProcessEvidenceRoles = "Attacker"
	SecurityProcessEvidenceRolescommandAndControl SecurityProcessEvidenceRoles = "CommandAndControl"
	SecurityProcessEvidenceRolescompromised       SecurityProcessEvidenceRoles = "Compromised"
	SecurityProcessEvidenceRolescontextual        SecurityProcessEvidenceRoles = "Contextual"
	SecurityProcessEvidenceRolescreated           SecurityProcessEvidenceRoles = "Created"
	SecurityProcessEvidenceRolesdestination       SecurityProcessEvidenceRoles = "Destination"
	SecurityProcessEvidenceRolesedited            SecurityProcessEvidenceRoles = "Edited"
	SecurityProcessEvidenceRolesloaded            SecurityProcessEvidenceRoles = "Loaded"
	SecurityProcessEvidenceRolespolicyViolator    SecurityProcessEvidenceRoles = "PolicyViolator"
	SecurityProcessEvidenceRolesscanned           SecurityProcessEvidenceRoles = "Scanned"
	SecurityProcessEvidenceRolessource            SecurityProcessEvidenceRoles = "Source"
	SecurityProcessEvidenceRolessuspicious        SecurityProcessEvidenceRoles = "Suspicious"
	SecurityProcessEvidenceRolesunknown           SecurityProcessEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityProcessEvidenceRoles() []string {
	return []string{
		string(SecurityProcessEvidenceRolesadded),
		string(SecurityProcessEvidenceRolesattacked),
		string(SecurityProcessEvidenceRolesattacker),
		string(SecurityProcessEvidenceRolescommandAndControl),
		string(SecurityProcessEvidenceRolescompromised),
		string(SecurityProcessEvidenceRolescontextual),
		string(SecurityProcessEvidenceRolescreated),
		string(SecurityProcessEvidenceRolesdestination),
		string(SecurityProcessEvidenceRolesedited),
		string(SecurityProcessEvidenceRolesloaded),
		string(SecurityProcessEvidenceRolespolicyViolator),
		string(SecurityProcessEvidenceRolesscanned),
		string(SecurityProcessEvidenceRolessource),
		string(SecurityProcessEvidenceRolessuspicious),
		string(SecurityProcessEvidenceRolesunknown),
	}
}

func parseSecurityProcessEvidenceRoles(input string) (*SecurityProcessEvidenceRoles, error) {
	vals := map[string]SecurityProcessEvidenceRoles{
		"added":             SecurityProcessEvidenceRolesadded,
		"attacked":          SecurityProcessEvidenceRolesattacked,
		"attacker":          SecurityProcessEvidenceRolesattacker,
		"commandandcontrol": SecurityProcessEvidenceRolescommandAndControl,
		"compromised":       SecurityProcessEvidenceRolescompromised,
		"contextual":        SecurityProcessEvidenceRolescontextual,
		"created":           SecurityProcessEvidenceRolescreated,
		"destination":       SecurityProcessEvidenceRolesdestination,
		"edited":            SecurityProcessEvidenceRolesedited,
		"loaded":            SecurityProcessEvidenceRolesloaded,
		"policyviolator":    SecurityProcessEvidenceRolespolicyViolator,
		"scanned":           SecurityProcessEvidenceRolesscanned,
		"source":            SecurityProcessEvidenceRolessource,
		"suspicious":        SecurityProcessEvidenceRolessuspicious,
		"unknown":           SecurityProcessEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityProcessEvidenceRoles(input)
	return &out, nil
}
