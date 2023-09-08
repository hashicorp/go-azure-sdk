package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBlobContainerEvidenceRoles string

const (
	SecurityBlobContainerEvidenceRolesadded             SecurityBlobContainerEvidenceRoles = "Added"
	SecurityBlobContainerEvidenceRolesattacked          SecurityBlobContainerEvidenceRoles = "Attacked"
	SecurityBlobContainerEvidenceRolesattacker          SecurityBlobContainerEvidenceRoles = "Attacker"
	SecurityBlobContainerEvidenceRolescommandAndControl SecurityBlobContainerEvidenceRoles = "CommandAndControl"
	SecurityBlobContainerEvidenceRolescompromised       SecurityBlobContainerEvidenceRoles = "Compromised"
	SecurityBlobContainerEvidenceRolescontextual        SecurityBlobContainerEvidenceRoles = "Contextual"
	SecurityBlobContainerEvidenceRolescreated           SecurityBlobContainerEvidenceRoles = "Created"
	SecurityBlobContainerEvidenceRolesdestination       SecurityBlobContainerEvidenceRoles = "Destination"
	SecurityBlobContainerEvidenceRolesedited            SecurityBlobContainerEvidenceRoles = "Edited"
	SecurityBlobContainerEvidenceRolesloaded            SecurityBlobContainerEvidenceRoles = "Loaded"
	SecurityBlobContainerEvidenceRolespolicyViolator    SecurityBlobContainerEvidenceRoles = "PolicyViolator"
	SecurityBlobContainerEvidenceRolesscanned           SecurityBlobContainerEvidenceRoles = "Scanned"
	SecurityBlobContainerEvidenceRolessource            SecurityBlobContainerEvidenceRoles = "Source"
	SecurityBlobContainerEvidenceRolessuspicious        SecurityBlobContainerEvidenceRoles = "Suspicious"
	SecurityBlobContainerEvidenceRolesunknown           SecurityBlobContainerEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityBlobContainerEvidenceRoles() []string {
	return []string{
		string(SecurityBlobContainerEvidenceRolesadded),
		string(SecurityBlobContainerEvidenceRolesattacked),
		string(SecurityBlobContainerEvidenceRolesattacker),
		string(SecurityBlobContainerEvidenceRolescommandAndControl),
		string(SecurityBlobContainerEvidenceRolescompromised),
		string(SecurityBlobContainerEvidenceRolescontextual),
		string(SecurityBlobContainerEvidenceRolescreated),
		string(SecurityBlobContainerEvidenceRolesdestination),
		string(SecurityBlobContainerEvidenceRolesedited),
		string(SecurityBlobContainerEvidenceRolesloaded),
		string(SecurityBlobContainerEvidenceRolespolicyViolator),
		string(SecurityBlobContainerEvidenceRolesscanned),
		string(SecurityBlobContainerEvidenceRolessource),
		string(SecurityBlobContainerEvidenceRolessuspicious),
		string(SecurityBlobContainerEvidenceRolesunknown),
	}
}

func parseSecurityBlobContainerEvidenceRoles(input string) (*SecurityBlobContainerEvidenceRoles, error) {
	vals := map[string]SecurityBlobContainerEvidenceRoles{
		"added":             SecurityBlobContainerEvidenceRolesadded,
		"attacked":          SecurityBlobContainerEvidenceRolesattacked,
		"attacker":          SecurityBlobContainerEvidenceRolesattacker,
		"commandandcontrol": SecurityBlobContainerEvidenceRolescommandAndControl,
		"compromised":       SecurityBlobContainerEvidenceRolescompromised,
		"contextual":        SecurityBlobContainerEvidenceRolescontextual,
		"created":           SecurityBlobContainerEvidenceRolescreated,
		"destination":       SecurityBlobContainerEvidenceRolesdestination,
		"edited":            SecurityBlobContainerEvidenceRolesedited,
		"loaded":            SecurityBlobContainerEvidenceRolesloaded,
		"policyviolator":    SecurityBlobContainerEvidenceRolespolicyViolator,
		"scanned":           SecurityBlobContainerEvidenceRolesscanned,
		"source":            SecurityBlobContainerEvidenceRolessource,
		"suspicious":        SecurityBlobContainerEvidenceRolessuspicious,
		"unknown":           SecurityBlobContainerEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBlobContainerEvidenceRoles(input)
	return &out, nil
}
