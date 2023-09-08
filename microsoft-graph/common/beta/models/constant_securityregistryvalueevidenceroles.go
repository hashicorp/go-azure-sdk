package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRegistryValueEvidenceRoles string

const (
	SecurityRegistryValueEvidenceRolesadded             SecurityRegistryValueEvidenceRoles = "Added"
	SecurityRegistryValueEvidenceRolesattacked          SecurityRegistryValueEvidenceRoles = "Attacked"
	SecurityRegistryValueEvidenceRolesattacker          SecurityRegistryValueEvidenceRoles = "Attacker"
	SecurityRegistryValueEvidenceRolescommandAndControl SecurityRegistryValueEvidenceRoles = "CommandAndControl"
	SecurityRegistryValueEvidenceRolescompromised       SecurityRegistryValueEvidenceRoles = "Compromised"
	SecurityRegistryValueEvidenceRolescontextual        SecurityRegistryValueEvidenceRoles = "Contextual"
	SecurityRegistryValueEvidenceRolescreated           SecurityRegistryValueEvidenceRoles = "Created"
	SecurityRegistryValueEvidenceRolesdestination       SecurityRegistryValueEvidenceRoles = "Destination"
	SecurityRegistryValueEvidenceRolesedited            SecurityRegistryValueEvidenceRoles = "Edited"
	SecurityRegistryValueEvidenceRolesloaded            SecurityRegistryValueEvidenceRoles = "Loaded"
	SecurityRegistryValueEvidenceRolespolicyViolator    SecurityRegistryValueEvidenceRoles = "PolicyViolator"
	SecurityRegistryValueEvidenceRolesscanned           SecurityRegistryValueEvidenceRoles = "Scanned"
	SecurityRegistryValueEvidenceRolessource            SecurityRegistryValueEvidenceRoles = "Source"
	SecurityRegistryValueEvidenceRolessuspicious        SecurityRegistryValueEvidenceRoles = "Suspicious"
	SecurityRegistryValueEvidenceRolesunknown           SecurityRegistryValueEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityRegistryValueEvidenceRoles() []string {
	return []string{
		string(SecurityRegistryValueEvidenceRolesadded),
		string(SecurityRegistryValueEvidenceRolesattacked),
		string(SecurityRegistryValueEvidenceRolesattacker),
		string(SecurityRegistryValueEvidenceRolescommandAndControl),
		string(SecurityRegistryValueEvidenceRolescompromised),
		string(SecurityRegistryValueEvidenceRolescontextual),
		string(SecurityRegistryValueEvidenceRolescreated),
		string(SecurityRegistryValueEvidenceRolesdestination),
		string(SecurityRegistryValueEvidenceRolesedited),
		string(SecurityRegistryValueEvidenceRolesloaded),
		string(SecurityRegistryValueEvidenceRolespolicyViolator),
		string(SecurityRegistryValueEvidenceRolesscanned),
		string(SecurityRegistryValueEvidenceRolessource),
		string(SecurityRegistryValueEvidenceRolessuspicious),
		string(SecurityRegistryValueEvidenceRolesunknown),
	}
}

func parseSecurityRegistryValueEvidenceRoles(input string) (*SecurityRegistryValueEvidenceRoles, error) {
	vals := map[string]SecurityRegistryValueEvidenceRoles{
		"added":             SecurityRegistryValueEvidenceRolesadded,
		"attacked":          SecurityRegistryValueEvidenceRolesattacked,
		"attacker":          SecurityRegistryValueEvidenceRolesattacker,
		"commandandcontrol": SecurityRegistryValueEvidenceRolescommandAndControl,
		"compromised":       SecurityRegistryValueEvidenceRolescompromised,
		"contextual":        SecurityRegistryValueEvidenceRolescontextual,
		"created":           SecurityRegistryValueEvidenceRolescreated,
		"destination":       SecurityRegistryValueEvidenceRolesdestination,
		"edited":            SecurityRegistryValueEvidenceRolesedited,
		"loaded":            SecurityRegistryValueEvidenceRolesloaded,
		"policyviolator":    SecurityRegistryValueEvidenceRolespolicyViolator,
		"scanned":           SecurityRegistryValueEvidenceRolesscanned,
		"source":            SecurityRegistryValueEvidenceRolessource,
		"suspicious":        SecurityRegistryValueEvidenceRolessuspicious,
		"unknown":           SecurityRegistryValueEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityRegistryValueEvidenceRoles(input)
	return &out, nil
}
