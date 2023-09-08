package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityUserEvidenceRoles string

const (
	SecurityUserEvidenceRolesadded             SecurityUserEvidenceRoles = "Added"
	SecurityUserEvidenceRolesattacked          SecurityUserEvidenceRoles = "Attacked"
	SecurityUserEvidenceRolesattacker          SecurityUserEvidenceRoles = "Attacker"
	SecurityUserEvidenceRolescommandAndControl SecurityUserEvidenceRoles = "CommandAndControl"
	SecurityUserEvidenceRolescompromised       SecurityUserEvidenceRoles = "Compromised"
	SecurityUserEvidenceRolescontextual        SecurityUserEvidenceRoles = "Contextual"
	SecurityUserEvidenceRolescreated           SecurityUserEvidenceRoles = "Created"
	SecurityUserEvidenceRolesdestination       SecurityUserEvidenceRoles = "Destination"
	SecurityUserEvidenceRolesedited            SecurityUserEvidenceRoles = "Edited"
	SecurityUserEvidenceRolesloaded            SecurityUserEvidenceRoles = "Loaded"
	SecurityUserEvidenceRolespolicyViolator    SecurityUserEvidenceRoles = "PolicyViolator"
	SecurityUserEvidenceRolesscanned           SecurityUserEvidenceRoles = "Scanned"
	SecurityUserEvidenceRolessource            SecurityUserEvidenceRoles = "Source"
	SecurityUserEvidenceRolessuspicious        SecurityUserEvidenceRoles = "Suspicious"
	SecurityUserEvidenceRolesunknown           SecurityUserEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityUserEvidenceRoles() []string {
	return []string{
		string(SecurityUserEvidenceRolesadded),
		string(SecurityUserEvidenceRolesattacked),
		string(SecurityUserEvidenceRolesattacker),
		string(SecurityUserEvidenceRolescommandAndControl),
		string(SecurityUserEvidenceRolescompromised),
		string(SecurityUserEvidenceRolescontextual),
		string(SecurityUserEvidenceRolescreated),
		string(SecurityUserEvidenceRolesdestination),
		string(SecurityUserEvidenceRolesedited),
		string(SecurityUserEvidenceRolesloaded),
		string(SecurityUserEvidenceRolespolicyViolator),
		string(SecurityUserEvidenceRolesscanned),
		string(SecurityUserEvidenceRolessource),
		string(SecurityUserEvidenceRolessuspicious),
		string(SecurityUserEvidenceRolesunknown),
	}
}

func parseSecurityUserEvidenceRoles(input string) (*SecurityUserEvidenceRoles, error) {
	vals := map[string]SecurityUserEvidenceRoles{
		"added":             SecurityUserEvidenceRolesadded,
		"attacked":          SecurityUserEvidenceRolesattacked,
		"attacker":          SecurityUserEvidenceRolesattacker,
		"commandandcontrol": SecurityUserEvidenceRolescommandAndControl,
		"compromised":       SecurityUserEvidenceRolescompromised,
		"contextual":        SecurityUserEvidenceRolescontextual,
		"created":           SecurityUserEvidenceRolescreated,
		"destination":       SecurityUserEvidenceRolesdestination,
		"edited":            SecurityUserEvidenceRolesedited,
		"loaded":            SecurityUserEvidenceRolesloaded,
		"policyviolator":    SecurityUserEvidenceRolespolicyViolator,
		"scanned":           SecurityUserEvidenceRolesscanned,
		"source":            SecurityUserEvidenceRolessource,
		"suspicious":        SecurityUserEvidenceRolessuspicious,
		"unknown":           SecurityUserEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityUserEvidenceRoles(input)
	return &out, nil
}
