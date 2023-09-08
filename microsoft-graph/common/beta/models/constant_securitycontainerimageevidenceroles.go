package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContainerImageEvidenceRoles string

const (
	SecurityContainerImageEvidenceRolesadded             SecurityContainerImageEvidenceRoles = "Added"
	SecurityContainerImageEvidenceRolesattacked          SecurityContainerImageEvidenceRoles = "Attacked"
	SecurityContainerImageEvidenceRolesattacker          SecurityContainerImageEvidenceRoles = "Attacker"
	SecurityContainerImageEvidenceRolescommandAndControl SecurityContainerImageEvidenceRoles = "CommandAndControl"
	SecurityContainerImageEvidenceRolescompromised       SecurityContainerImageEvidenceRoles = "Compromised"
	SecurityContainerImageEvidenceRolescontextual        SecurityContainerImageEvidenceRoles = "Contextual"
	SecurityContainerImageEvidenceRolescreated           SecurityContainerImageEvidenceRoles = "Created"
	SecurityContainerImageEvidenceRolesdestination       SecurityContainerImageEvidenceRoles = "Destination"
	SecurityContainerImageEvidenceRolesedited            SecurityContainerImageEvidenceRoles = "Edited"
	SecurityContainerImageEvidenceRolesloaded            SecurityContainerImageEvidenceRoles = "Loaded"
	SecurityContainerImageEvidenceRolespolicyViolator    SecurityContainerImageEvidenceRoles = "PolicyViolator"
	SecurityContainerImageEvidenceRolesscanned           SecurityContainerImageEvidenceRoles = "Scanned"
	SecurityContainerImageEvidenceRolessource            SecurityContainerImageEvidenceRoles = "Source"
	SecurityContainerImageEvidenceRolessuspicious        SecurityContainerImageEvidenceRoles = "Suspicious"
	SecurityContainerImageEvidenceRolesunknown           SecurityContainerImageEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityContainerImageEvidenceRoles() []string {
	return []string{
		string(SecurityContainerImageEvidenceRolesadded),
		string(SecurityContainerImageEvidenceRolesattacked),
		string(SecurityContainerImageEvidenceRolesattacker),
		string(SecurityContainerImageEvidenceRolescommandAndControl),
		string(SecurityContainerImageEvidenceRolescompromised),
		string(SecurityContainerImageEvidenceRolescontextual),
		string(SecurityContainerImageEvidenceRolescreated),
		string(SecurityContainerImageEvidenceRolesdestination),
		string(SecurityContainerImageEvidenceRolesedited),
		string(SecurityContainerImageEvidenceRolesloaded),
		string(SecurityContainerImageEvidenceRolespolicyViolator),
		string(SecurityContainerImageEvidenceRolesscanned),
		string(SecurityContainerImageEvidenceRolessource),
		string(SecurityContainerImageEvidenceRolessuspicious),
		string(SecurityContainerImageEvidenceRolesunknown),
	}
}

func parseSecurityContainerImageEvidenceRoles(input string) (*SecurityContainerImageEvidenceRoles, error) {
	vals := map[string]SecurityContainerImageEvidenceRoles{
		"added":             SecurityContainerImageEvidenceRolesadded,
		"attacked":          SecurityContainerImageEvidenceRolesattacked,
		"attacker":          SecurityContainerImageEvidenceRolesattacker,
		"commandandcontrol": SecurityContainerImageEvidenceRolescommandAndControl,
		"compromised":       SecurityContainerImageEvidenceRolescompromised,
		"contextual":        SecurityContainerImageEvidenceRolescontextual,
		"created":           SecurityContainerImageEvidenceRolescreated,
		"destination":       SecurityContainerImageEvidenceRolesdestination,
		"edited":            SecurityContainerImageEvidenceRolesedited,
		"loaded":            SecurityContainerImageEvidenceRolesloaded,
		"policyviolator":    SecurityContainerImageEvidenceRolespolicyViolator,
		"scanned":           SecurityContainerImageEvidenceRolesscanned,
		"source":            SecurityContainerImageEvidenceRolessource,
		"suspicious":        SecurityContainerImageEvidenceRolessuspicious,
		"unknown":           SecurityContainerImageEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContainerImageEvidenceRoles(input)
	return &out, nil
}
