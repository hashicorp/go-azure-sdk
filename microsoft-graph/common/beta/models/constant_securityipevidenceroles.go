package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIpEvidenceRoles string

const (
	SecurityIpEvidenceRolesadded             SecurityIpEvidenceRoles = "Added"
	SecurityIpEvidenceRolesattacked          SecurityIpEvidenceRoles = "Attacked"
	SecurityIpEvidenceRolesattacker          SecurityIpEvidenceRoles = "Attacker"
	SecurityIpEvidenceRolescommandAndControl SecurityIpEvidenceRoles = "CommandAndControl"
	SecurityIpEvidenceRolescompromised       SecurityIpEvidenceRoles = "Compromised"
	SecurityIpEvidenceRolescontextual        SecurityIpEvidenceRoles = "Contextual"
	SecurityIpEvidenceRolescreated           SecurityIpEvidenceRoles = "Created"
	SecurityIpEvidenceRolesdestination       SecurityIpEvidenceRoles = "Destination"
	SecurityIpEvidenceRolesedited            SecurityIpEvidenceRoles = "Edited"
	SecurityIpEvidenceRolesloaded            SecurityIpEvidenceRoles = "Loaded"
	SecurityIpEvidenceRolespolicyViolator    SecurityIpEvidenceRoles = "PolicyViolator"
	SecurityIpEvidenceRolesscanned           SecurityIpEvidenceRoles = "Scanned"
	SecurityIpEvidenceRolessource            SecurityIpEvidenceRoles = "Source"
	SecurityIpEvidenceRolessuspicious        SecurityIpEvidenceRoles = "Suspicious"
	SecurityIpEvidenceRolesunknown           SecurityIpEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityIpEvidenceRoles() []string {
	return []string{
		string(SecurityIpEvidenceRolesadded),
		string(SecurityIpEvidenceRolesattacked),
		string(SecurityIpEvidenceRolesattacker),
		string(SecurityIpEvidenceRolescommandAndControl),
		string(SecurityIpEvidenceRolescompromised),
		string(SecurityIpEvidenceRolescontextual),
		string(SecurityIpEvidenceRolescreated),
		string(SecurityIpEvidenceRolesdestination),
		string(SecurityIpEvidenceRolesedited),
		string(SecurityIpEvidenceRolesloaded),
		string(SecurityIpEvidenceRolespolicyViolator),
		string(SecurityIpEvidenceRolesscanned),
		string(SecurityIpEvidenceRolessource),
		string(SecurityIpEvidenceRolessuspicious),
		string(SecurityIpEvidenceRolesunknown),
	}
}

func parseSecurityIpEvidenceRoles(input string) (*SecurityIpEvidenceRoles, error) {
	vals := map[string]SecurityIpEvidenceRoles{
		"added":             SecurityIpEvidenceRolesadded,
		"attacked":          SecurityIpEvidenceRolesattacked,
		"attacker":          SecurityIpEvidenceRolesattacker,
		"commandandcontrol": SecurityIpEvidenceRolescommandAndControl,
		"compromised":       SecurityIpEvidenceRolescompromised,
		"contextual":        SecurityIpEvidenceRolescontextual,
		"created":           SecurityIpEvidenceRolescreated,
		"destination":       SecurityIpEvidenceRolesdestination,
		"edited":            SecurityIpEvidenceRolesedited,
		"loaded":            SecurityIpEvidenceRolesloaded,
		"policyviolator":    SecurityIpEvidenceRolespolicyViolator,
		"scanned":           SecurityIpEvidenceRolesscanned,
		"source":            SecurityIpEvidenceRolessource,
		"suspicious":        SecurityIpEvidenceRolessuspicious,
		"unknown":           SecurityIpEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityIpEvidenceRoles(input)
	return &out, nil
}
