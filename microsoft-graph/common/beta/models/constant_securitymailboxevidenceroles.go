package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityMailboxEvidenceRoles string

const (
	SecurityMailboxEvidenceRolesadded             SecurityMailboxEvidenceRoles = "Added"
	SecurityMailboxEvidenceRolesattacked          SecurityMailboxEvidenceRoles = "Attacked"
	SecurityMailboxEvidenceRolesattacker          SecurityMailboxEvidenceRoles = "Attacker"
	SecurityMailboxEvidenceRolescommandAndControl SecurityMailboxEvidenceRoles = "CommandAndControl"
	SecurityMailboxEvidenceRolescompromised       SecurityMailboxEvidenceRoles = "Compromised"
	SecurityMailboxEvidenceRolescontextual        SecurityMailboxEvidenceRoles = "Contextual"
	SecurityMailboxEvidenceRolescreated           SecurityMailboxEvidenceRoles = "Created"
	SecurityMailboxEvidenceRolesdestination       SecurityMailboxEvidenceRoles = "Destination"
	SecurityMailboxEvidenceRolesedited            SecurityMailboxEvidenceRoles = "Edited"
	SecurityMailboxEvidenceRolesloaded            SecurityMailboxEvidenceRoles = "Loaded"
	SecurityMailboxEvidenceRolespolicyViolator    SecurityMailboxEvidenceRoles = "PolicyViolator"
	SecurityMailboxEvidenceRolesscanned           SecurityMailboxEvidenceRoles = "Scanned"
	SecurityMailboxEvidenceRolessource            SecurityMailboxEvidenceRoles = "Source"
	SecurityMailboxEvidenceRolessuspicious        SecurityMailboxEvidenceRoles = "Suspicious"
	SecurityMailboxEvidenceRolesunknown           SecurityMailboxEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityMailboxEvidenceRoles() []string {
	return []string{
		string(SecurityMailboxEvidenceRolesadded),
		string(SecurityMailboxEvidenceRolesattacked),
		string(SecurityMailboxEvidenceRolesattacker),
		string(SecurityMailboxEvidenceRolescommandAndControl),
		string(SecurityMailboxEvidenceRolescompromised),
		string(SecurityMailboxEvidenceRolescontextual),
		string(SecurityMailboxEvidenceRolescreated),
		string(SecurityMailboxEvidenceRolesdestination),
		string(SecurityMailboxEvidenceRolesedited),
		string(SecurityMailboxEvidenceRolesloaded),
		string(SecurityMailboxEvidenceRolespolicyViolator),
		string(SecurityMailboxEvidenceRolesscanned),
		string(SecurityMailboxEvidenceRolessource),
		string(SecurityMailboxEvidenceRolessuspicious),
		string(SecurityMailboxEvidenceRolesunknown),
	}
}

func parseSecurityMailboxEvidenceRoles(input string) (*SecurityMailboxEvidenceRoles, error) {
	vals := map[string]SecurityMailboxEvidenceRoles{
		"added":             SecurityMailboxEvidenceRolesadded,
		"attacked":          SecurityMailboxEvidenceRolesattacked,
		"attacker":          SecurityMailboxEvidenceRolesattacker,
		"commandandcontrol": SecurityMailboxEvidenceRolescommandAndControl,
		"compromised":       SecurityMailboxEvidenceRolescompromised,
		"contextual":        SecurityMailboxEvidenceRolescontextual,
		"created":           SecurityMailboxEvidenceRolescreated,
		"destination":       SecurityMailboxEvidenceRolesdestination,
		"edited":            SecurityMailboxEvidenceRolesedited,
		"loaded":            SecurityMailboxEvidenceRolesloaded,
		"policyviolator":    SecurityMailboxEvidenceRolespolicyViolator,
		"scanned":           SecurityMailboxEvidenceRolesscanned,
		"source":            SecurityMailboxEvidenceRolessource,
		"suspicious":        SecurityMailboxEvidenceRolessuspicious,
		"unknown":           SecurityMailboxEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityMailboxEvidenceRoles(input)
	return &out, nil
}
