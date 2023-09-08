package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidenceRoles string

const (
	SecurityDeviceEvidenceRolesadded             SecurityDeviceEvidenceRoles = "Added"
	SecurityDeviceEvidenceRolesattacked          SecurityDeviceEvidenceRoles = "Attacked"
	SecurityDeviceEvidenceRolesattacker          SecurityDeviceEvidenceRoles = "Attacker"
	SecurityDeviceEvidenceRolescommandAndControl SecurityDeviceEvidenceRoles = "CommandAndControl"
	SecurityDeviceEvidenceRolescompromised       SecurityDeviceEvidenceRoles = "Compromised"
	SecurityDeviceEvidenceRolescontextual        SecurityDeviceEvidenceRoles = "Contextual"
	SecurityDeviceEvidenceRolescreated           SecurityDeviceEvidenceRoles = "Created"
	SecurityDeviceEvidenceRolesdestination       SecurityDeviceEvidenceRoles = "Destination"
	SecurityDeviceEvidenceRolesedited            SecurityDeviceEvidenceRoles = "Edited"
	SecurityDeviceEvidenceRolesloaded            SecurityDeviceEvidenceRoles = "Loaded"
	SecurityDeviceEvidenceRolespolicyViolator    SecurityDeviceEvidenceRoles = "PolicyViolator"
	SecurityDeviceEvidenceRolesscanned           SecurityDeviceEvidenceRoles = "Scanned"
	SecurityDeviceEvidenceRolessource            SecurityDeviceEvidenceRoles = "Source"
	SecurityDeviceEvidenceRolessuspicious        SecurityDeviceEvidenceRoles = "Suspicious"
	SecurityDeviceEvidenceRolesunknown           SecurityDeviceEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityDeviceEvidenceRoles() []string {
	return []string{
		string(SecurityDeviceEvidenceRolesadded),
		string(SecurityDeviceEvidenceRolesattacked),
		string(SecurityDeviceEvidenceRolesattacker),
		string(SecurityDeviceEvidenceRolescommandAndControl),
		string(SecurityDeviceEvidenceRolescompromised),
		string(SecurityDeviceEvidenceRolescontextual),
		string(SecurityDeviceEvidenceRolescreated),
		string(SecurityDeviceEvidenceRolesdestination),
		string(SecurityDeviceEvidenceRolesedited),
		string(SecurityDeviceEvidenceRolesloaded),
		string(SecurityDeviceEvidenceRolespolicyViolator),
		string(SecurityDeviceEvidenceRolesscanned),
		string(SecurityDeviceEvidenceRolessource),
		string(SecurityDeviceEvidenceRolessuspicious),
		string(SecurityDeviceEvidenceRolesunknown),
	}
}

func parseSecurityDeviceEvidenceRoles(input string) (*SecurityDeviceEvidenceRoles, error) {
	vals := map[string]SecurityDeviceEvidenceRoles{
		"added":             SecurityDeviceEvidenceRolesadded,
		"attacked":          SecurityDeviceEvidenceRolesattacked,
		"attacker":          SecurityDeviceEvidenceRolesattacker,
		"commandandcontrol": SecurityDeviceEvidenceRolescommandAndControl,
		"compromised":       SecurityDeviceEvidenceRolescompromised,
		"contextual":        SecurityDeviceEvidenceRolescontextual,
		"created":           SecurityDeviceEvidenceRolescreated,
		"destination":       SecurityDeviceEvidenceRolesdestination,
		"edited":            SecurityDeviceEvidenceRolesedited,
		"loaded":            SecurityDeviceEvidenceRolesloaded,
		"policyviolator":    SecurityDeviceEvidenceRolespolicyViolator,
		"scanned":           SecurityDeviceEvidenceRolesscanned,
		"source":            SecurityDeviceEvidenceRolessource,
		"suspicious":        SecurityDeviceEvidenceRolessuspicious,
		"unknown":           SecurityDeviceEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityDeviceEvidenceRoles(input)
	return &out, nil
}
