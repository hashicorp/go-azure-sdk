package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryKeyEvidenceRoles string

const (
	SecurityRegistryKeyEvidenceRolesadded             SecurityRegistryKeyEvidenceRoles = "Added"
	SecurityRegistryKeyEvidenceRolesattacked          SecurityRegistryKeyEvidenceRoles = "Attacked"
	SecurityRegistryKeyEvidenceRolesattacker          SecurityRegistryKeyEvidenceRoles = "Attacker"
	SecurityRegistryKeyEvidenceRolescommandAndControl SecurityRegistryKeyEvidenceRoles = "CommandAndControl"
	SecurityRegistryKeyEvidenceRolescompromised       SecurityRegistryKeyEvidenceRoles = "Compromised"
	SecurityRegistryKeyEvidenceRolescontextual        SecurityRegistryKeyEvidenceRoles = "Contextual"
	SecurityRegistryKeyEvidenceRolescreated           SecurityRegistryKeyEvidenceRoles = "Created"
	SecurityRegistryKeyEvidenceRolesdestination       SecurityRegistryKeyEvidenceRoles = "Destination"
	SecurityRegistryKeyEvidenceRolesedited            SecurityRegistryKeyEvidenceRoles = "Edited"
	SecurityRegistryKeyEvidenceRolesloaded            SecurityRegistryKeyEvidenceRoles = "Loaded"
	SecurityRegistryKeyEvidenceRolespolicyViolator    SecurityRegistryKeyEvidenceRoles = "PolicyViolator"
	SecurityRegistryKeyEvidenceRolesscanned           SecurityRegistryKeyEvidenceRoles = "Scanned"
	SecurityRegistryKeyEvidenceRolessource            SecurityRegistryKeyEvidenceRoles = "Source"
	SecurityRegistryKeyEvidenceRolessuspicious        SecurityRegistryKeyEvidenceRoles = "Suspicious"
	SecurityRegistryKeyEvidenceRolesunknown           SecurityRegistryKeyEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityRegistryKeyEvidenceRoles() []string {
	return []string{
		string(SecurityRegistryKeyEvidenceRolesadded),
		string(SecurityRegistryKeyEvidenceRolesattacked),
		string(SecurityRegistryKeyEvidenceRolesattacker),
		string(SecurityRegistryKeyEvidenceRolescommandAndControl),
		string(SecurityRegistryKeyEvidenceRolescompromised),
		string(SecurityRegistryKeyEvidenceRolescontextual),
		string(SecurityRegistryKeyEvidenceRolescreated),
		string(SecurityRegistryKeyEvidenceRolesdestination),
		string(SecurityRegistryKeyEvidenceRolesedited),
		string(SecurityRegistryKeyEvidenceRolesloaded),
		string(SecurityRegistryKeyEvidenceRolespolicyViolator),
		string(SecurityRegistryKeyEvidenceRolesscanned),
		string(SecurityRegistryKeyEvidenceRolessource),
		string(SecurityRegistryKeyEvidenceRolessuspicious),
		string(SecurityRegistryKeyEvidenceRolesunknown),
	}
}

func parseSecurityRegistryKeyEvidenceRoles(input string) (*SecurityRegistryKeyEvidenceRoles, error) {
	vals := map[string]SecurityRegistryKeyEvidenceRoles{
		"added":             SecurityRegistryKeyEvidenceRolesadded,
		"attacked":          SecurityRegistryKeyEvidenceRolesattacked,
		"attacker":          SecurityRegistryKeyEvidenceRolesattacker,
		"commandandcontrol": SecurityRegistryKeyEvidenceRolescommandAndControl,
		"compromised":       SecurityRegistryKeyEvidenceRolescompromised,
		"contextual":        SecurityRegistryKeyEvidenceRolescontextual,
		"created":           SecurityRegistryKeyEvidenceRolescreated,
		"destination":       SecurityRegistryKeyEvidenceRolesdestination,
		"edited":            SecurityRegistryKeyEvidenceRolesedited,
		"loaded":            SecurityRegistryKeyEvidenceRolesloaded,
		"policyviolator":    SecurityRegistryKeyEvidenceRolespolicyViolator,
		"scanned":           SecurityRegistryKeyEvidenceRolesscanned,
		"source":            SecurityRegistryKeyEvidenceRolessource,
		"suspicious":        SecurityRegistryKeyEvidenceRolessuspicious,
		"unknown":           SecurityRegistryKeyEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryKeyEvidenceRoles(input)
	return &out, nil
}
