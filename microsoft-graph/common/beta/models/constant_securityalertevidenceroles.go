package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertEvidenceRoles string

const (
	SecurityAlertEvidenceRolesadded             SecurityAlertEvidenceRoles = "Added"
	SecurityAlertEvidenceRolesattacked          SecurityAlertEvidenceRoles = "Attacked"
	SecurityAlertEvidenceRolesattacker          SecurityAlertEvidenceRoles = "Attacker"
	SecurityAlertEvidenceRolescommandAndControl SecurityAlertEvidenceRoles = "CommandAndControl"
	SecurityAlertEvidenceRolescompromised       SecurityAlertEvidenceRoles = "Compromised"
	SecurityAlertEvidenceRolescontextual        SecurityAlertEvidenceRoles = "Contextual"
	SecurityAlertEvidenceRolescreated           SecurityAlertEvidenceRoles = "Created"
	SecurityAlertEvidenceRolesdestination       SecurityAlertEvidenceRoles = "Destination"
	SecurityAlertEvidenceRolesedited            SecurityAlertEvidenceRoles = "Edited"
	SecurityAlertEvidenceRolesloaded            SecurityAlertEvidenceRoles = "Loaded"
	SecurityAlertEvidenceRolespolicyViolator    SecurityAlertEvidenceRoles = "PolicyViolator"
	SecurityAlertEvidenceRolesscanned           SecurityAlertEvidenceRoles = "Scanned"
	SecurityAlertEvidenceRolessource            SecurityAlertEvidenceRoles = "Source"
	SecurityAlertEvidenceRolessuspicious        SecurityAlertEvidenceRoles = "Suspicious"
	SecurityAlertEvidenceRolesunknown           SecurityAlertEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityAlertEvidenceRoles() []string {
	return []string{
		string(SecurityAlertEvidenceRolesadded),
		string(SecurityAlertEvidenceRolesattacked),
		string(SecurityAlertEvidenceRolesattacker),
		string(SecurityAlertEvidenceRolescommandAndControl),
		string(SecurityAlertEvidenceRolescompromised),
		string(SecurityAlertEvidenceRolescontextual),
		string(SecurityAlertEvidenceRolescreated),
		string(SecurityAlertEvidenceRolesdestination),
		string(SecurityAlertEvidenceRolesedited),
		string(SecurityAlertEvidenceRolesloaded),
		string(SecurityAlertEvidenceRolespolicyViolator),
		string(SecurityAlertEvidenceRolesscanned),
		string(SecurityAlertEvidenceRolessource),
		string(SecurityAlertEvidenceRolessuspicious),
		string(SecurityAlertEvidenceRolesunknown),
	}
}

func parseSecurityAlertEvidenceRoles(input string) (*SecurityAlertEvidenceRoles, error) {
	vals := map[string]SecurityAlertEvidenceRoles{
		"added":             SecurityAlertEvidenceRolesadded,
		"attacked":          SecurityAlertEvidenceRolesattacked,
		"attacker":          SecurityAlertEvidenceRolesattacker,
		"commandandcontrol": SecurityAlertEvidenceRolescommandAndControl,
		"compromised":       SecurityAlertEvidenceRolescompromised,
		"contextual":        SecurityAlertEvidenceRolescontextual,
		"created":           SecurityAlertEvidenceRolescreated,
		"destination":       SecurityAlertEvidenceRolesdestination,
		"edited":            SecurityAlertEvidenceRolesedited,
		"loaded":            SecurityAlertEvidenceRolesloaded,
		"policyviolator":    SecurityAlertEvidenceRolespolicyViolator,
		"scanned":           SecurityAlertEvidenceRolesscanned,
		"source":            SecurityAlertEvidenceRolessource,
		"suspicious":        SecurityAlertEvidenceRolessuspicious,
		"unknown":           SecurityAlertEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertEvidenceRoles(input)
	return &out, nil
}
