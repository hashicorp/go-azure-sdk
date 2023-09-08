package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecuritySecurityGroupEvidenceRoles string

const (
	SecuritySecurityGroupEvidenceRolesadded             SecuritySecurityGroupEvidenceRoles = "Added"
	SecuritySecurityGroupEvidenceRolesattacked          SecuritySecurityGroupEvidenceRoles = "Attacked"
	SecuritySecurityGroupEvidenceRolesattacker          SecuritySecurityGroupEvidenceRoles = "Attacker"
	SecuritySecurityGroupEvidenceRolescommandAndControl SecuritySecurityGroupEvidenceRoles = "CommandAndControl"
	SecuritySecurityGroupEvidenceRolescompromised       SecuritySecurityGroupEvidenceRoles = "Compromised"
	SecuritySecurityGroupEvidenceRolescontextual        SecuritySecurityGroupEvidenceRoles = "Contextual"
	SecuritySecurityGroupEvidenceRolescreated           SecuritySecurityGroupEvidenceRoles = "Created"
	SecuritySecurityGroupEvidenceRolesdestination       SecuritySecurityGroupEvidenceRoles = "Destination"
	SecuritySecurityGroupEvidenceRolesedited            SecuritySecurityGroupEvidenceRoles = "Edited"
	SecuritySecurityGroupEvidenceRolesloaded            SecuritySecurityGroupEvidenceRoles = "Loaded"
	SecuritySecurityGroupEvidenceRolespolicyViolator    SecuritySecurityGroupEvidenceRoles = "PolicyViolator"
	SecuritySecurityGroupEvidenceRolesscanned           SecuritySecurityGroupEvidenceRoles = "Scanned"
	SecuritySecurityGroupEvidenceRolessource            SecuritySecurityGroupEvidenceRoles = "Source"
	SecuritySecurityGroupEvidenceRolessuspicious        SecuritySecurityGroupEvidenceRoles = "Suspicious"
	SecuritySecurityGroupEvidenceRolesunknown           SecuritySecurityGroupEvidenceRoles = "Unknown"
)

func PossibleValuesForSecuritySecurityGroupEvidenceRoles() []string {
	return []string{
		string(SecuritySecurityGroupEvidenceRolesadded),
		string(SecuritySecurityGroupEvidenceRolesattacked),
		string(SecuritySecurityGroupEvidenceRolesattacker),
		string(SecuritySecurityGroupEvidenceRolescommandAndControl),
		string(SecuritySecurityGroupEvidenceRolescompromised),
		string(SecuritySecurityGroupEvidenceRolescontextual),
		string(SecuritySecurityGroupEvidenceRolescreated),
		string(SecuritySecurityGroupEvidenceRolesdestination),
		string(SecuritySecurityGroupEvidenceRolesedited),
		string(SecuritySecurityGroupEvidenceRolesloaded),
		string(SecuritySecurityGroupEvidenceRolespolicyViolator),
		string(SecuritySecurityGroupEvidenceRolesscanned),
		string(SecuritySecurityGroupEvidenceRolessource),
		string(SecuritySecurityGroupEvidenceRolessuspicious),
		string(SecuritySecurityGroupEvidenceRolesunknown),
	}
}

func parseSecuritySecurityGroupEvidenceRoles(input string) (*SecuritySecurityGroupEvidenceRoles, error) {
	vals := map[string]SecuritySecurityGroupEvidenceRoles{
		"added":             SecuritySecurityGroupEvidenceRolesadded,
		"attacked":          SecuritySecurityGroupEvidenceRolesattacked,
		"attacker":          SecuritySecurityGroupEvidenceRolesattacker,
		"commandandcontrol": SecuritySecurityGroupEvidenceRolescommandAndControl,
		"compromised":       SecuritySecurityGroupEvidenceRolescompromised,
		"contextual":        SecuritySecurityGroupEvidenceRolescontextual,
		"created":           SecuritySecurityGroupEvidenceRolescreated,
		"destination":       SecuritySecurityGroupEvidenceRolesdestination,
		"edited":            SecuritySecurityGroupEvidenceRolesedited,
		"loaded":            SecuritySecurityGroupEvidenceRolesloaded,
		"policyviolator":    SecuritySecurityGroupEvidenceRolespolicyViolator,
		"scanned":           SecuritySecurityGroupEvidenceRolesscanned,
		"source":            SecuritySecurityGroupEvidenceRolessource,
		"suspicious":        SecuritySecurityGroupEvidenceRolessuspicious,
		"unknown":           SecuritySecurityGroupEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecuritySecurityGroupEvidenceRoles(input)
	return &out, nil
}
