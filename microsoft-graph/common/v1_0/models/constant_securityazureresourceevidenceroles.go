package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAzureResourceEvidenceRoles string

const (
	SecurityAzureResourceEvidenceRolesadded             SecurityAzureResourceEvidenceRoles = "Added"
	SecurityAzureResourceEvidenceRolesattacked          SecurityAzureResourceEvidenceRoles = "Attacked"
	SecurityAzureResourceEvidenceRolesattacker          SecurityAzureResourceEvidenceRoles = "Attacker"
	SecurityAzureResourceEvidenceRolescommandAndControl SecurityAzureResourceEvidenceRoles = "CommandAndControl"
	SecurityAzureResourceEvidenceRolescompromised       SecurityAzureResourceEvidenceRoles = "Compromised"
	SecurityAzureResourceEvidenceRolescontextual        SecurityAzureResourceEvidenceRoles = "Contextual"
	SecurityAzureResourceEvidenceRolescreated           SecurityAzureResourceEvidenceRoles = "Created"
	SecurityAzureResourceEvidenceRolesdestination       SecurityAzureResourceEvidenceRoles = "Destination"
	SecurityAzureResourceEvidenceRolesedited            SecurityAzureResourceEvidenceRoles = "Edited"
	SecurityAzureResourceEvidenceRolesloaded            SecurityAzureResourceEvidenceRoles = "Loaded"
	SecurityAzureResourceEvidenceRolespolicyViolator    SecurityAzureResourceEvidenceRoles = "PolicyViolator"
	SecurityAzureResourceEvidenceRolesscanned           SecurityAzureResourceEvidenceRoles = "Scanned"
	SecurityAzureResourceEvidenceRolessource            SecurityAzureResourceEvidenceRoles = "Source"
	SecurityAzureResourceEvidenceRolessuspicious        SecurityAzureResourceEvidenceRoles = "Suspicious"
	SecurityAzureResourceEvidenceRolesunknown           SecurityAzureResourceEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityAzureResourceEvidenceRoles() []string {
	return []string{
		string(SecurityAzureResourceEvidenceRolesadded),
		string(SecurityAzureResourceEvidenceRolesattacked),
		string(SecurityAzureResourceEvidenceRolesattacker),
		string(SecurityAzureResourceEvidenceRolescommandAndControl),
		string(SecurityAzureResourceEvidenceRolescompromised),
		string(SecurityAzureResourceEvidenceRolescontextual),
		string(SecurityAzureResourceEvidenceRolescreated),
		string(SecurityAzureResourceEvidenceRolesdestination),
		string(SecurityAzureResourceEvidenceRolesedited),
		string(SecurityAzureResourceEvidenceRolesloaded),
		string(SecurityAzureResourceEvidenceRolespolicyViolator),
		string(SecurityAzureResourceEvidenceRolesscanned),
		string(SecurityAzureResourceEvidenceRolessource),
		string(SecurityAzureResourceEvidenceRolessuspicious),
		string(SecurityAzureResourceEvidenceRolesunknown),
	}
}

func parseSecurityAzureResourceEvidenceRoles(input string) (*SecurityAzureResourceEvidenceRoles, error) {
	vals := map[string]SecurityAzureResourceEvidenceRoles{
		"added":             SecurityAzureResourceEvidenceRolesadded,
		"attacked":          SecurityAzureResourceEvidenceRolesattacked,
		"attacker":          SecurityAzureResourceEvidenceRolesattacker,
		"commandandcontrol": SecurityAzureResourceEvidenceRolescommandAndControl,
		"compromised":       SecurityAzureResourceEvidenceRolescompromised,
		"contextual":        SecurityAzureResourceEvidenceRolescontextual,
		"created":           SecurityAzureResourceEvidenceRolescreated,
		"destination":       SecurityAzureResourceEvidenceRolesdestination,
		"edited":            SecurityAzureResourceEvidenceRolesedited,
		"loaded":            SecurityAzureResourceEvidenceRolesloaded,
		"policyviolator":    SecurityAzureResourceEvidenceRolespolicyViolator,
		"scanned":           SecurityAzureResourceEvidenceRolesscanned,
		"source":            SecurityAzureResourceEvidenceRolessource,
		"suspicious":        SecurityAzureResourceEvidenceRolessuspicious,
		"unknown":           SecurityAzureResourceEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAzureResourceEvidenceRoles(input)
	return &out, nil
}
