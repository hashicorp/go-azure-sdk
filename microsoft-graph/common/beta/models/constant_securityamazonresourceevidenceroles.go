package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAmazonResourceEvidenceRoles string

const (
	SecurityAmazonResourceEvidenceRolesadded             SecurityAmazonResourceEvidenceRoles = "Added"
	SecurityAmazonResourceEvidenceRolesattacked          SecurityAmazonResourceEvidenceRoles = "Attacked"
	SecurityAmazonResourceEvidenceRolesattacker          SecurityAmazonResourceEvidenceRoles = "Attacker"
	SecurityAmazonResourceEvidenceRolescommandAndControl SecurityAmazonResourceEvidenceRoles = "CommandAndControl"
	SecurityAmazonResourceEvidenceRolescompromised       SecurityAmazonResourceEvidenceRoles = "Compromised"
	SecurityAmazonResourceEvidenceRolescontextual        SecurityAmazonResourceEvidenceRoles = "Contextual"
	SecurityAmazonResourceEvidenceRolescreated           SecurityAmazonResourceEvidenceRoles = "Created"
	SecurityAmazonResourceEvidenceRolesdestination       SecurityAmazonResourceEvidenceRoles = "Destination"
	SecurityAmazonResourceEvidenceRolesedited            SecurityAmazonResourceEvidenceRoles = "Edited"
	SecurityAmazonResourceEvidenceRolesloaded            SecurityAmazonResourceEvidenceRoles = "Loaded"
	SecurityAmazonResourceEvidenceRolespolicyViolator    SecurityAmazonResourceEvidenceRoles = "PolicyViolator"
	SecurityAmazonResourceEvidenceRolesscanned           SecurityAmazonResourceEvidenceRoles = "Scanned"
	SecurityAmazonResourceEvidenceRolessource            SecurityAmazonResourceEvidenceRoles = "Source"
	SecurityAmazonResourceEvidenceRolessuspicious        SecurityAmazonResourceEvidenceRoles = "Suspicious"
	SecurityAmazonResourceEvidenceRolesunknown           SecurityAmazonResourceEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityAmazonResourceEvidenceRoles() []string {
	return []string{
		string(SecurityAmazonResourceEvidenceRolesadded),
		string(SecurityAmazonResourceEvidenceRolesattacked),
		string(SecurityAmazonResourceEvidenceRolesattacker),
		string(SecurityAmazonResourceEvidenceRolescommandAndControl),
		string(SecurityAmazonResourceEvidenceRolescompromised),
		string(SecurityAmazonResourceEvidenceRolescontextual),
		string(SecurityAmazonResourceEvidenceRolescreated),
		string(SecurityAmazonResourceEvidenceRolesdestination),
		string(SecurityAmazonResourceEvidenceRolesedited),
		string(SecurityAmazonResourceEvidenceRolesloaded),
		string(SecurityAmazonResourceEvidenceRolespolicyViolator),
		string(SecurityAmazonResourceEvidenceRolesscanned),
		string(SecurityAmazonResourceEvidenceRolessource),
		string(SecurityAmazonResourceEvidenceRolessuspicious),
		string(SecurityAmazonResourceEvidenceRolesunknown),
	}
}

func parseSecurityAmazonResourceEvidenceRoles(input string) (*SecurityAmazonResourceEvidenceRoles, error) {
	vals := map[string]SecurityAmazonResourceEvidenceRoles{
		"added":             SecurityAmazonResourceEvidenceRolesadded,
		"attacked":          SecurityAmazonResourceEvidenceRolesattacked,
		"attacker":          SecurityAmazonResourceEvidenceRolesattacker,
		"commandandcontrol": SecurityAmazonResourceEvidenceRolescommandAndControl,
		"compromised":       SecurityAmazonResourceEvidenceRolescompromised,
		"contextual":        SecurityAmazonResourceEvidenceRolescontextual,
		"created":           SecurityAmazonResourceEvidenceRolescreated,
		"destination":       SecurityAmazonResourceEvidenceRolesdestination,
		"edited":            SecurityAmazonResourceEvidenceRolesedited,
		"loaded":            SecurityAmazonResourceEvidenceRolesloaded,
		"policyviolator":    SecurityAmazonResourceEvidenceRolespolicyViolator,
		"scanned":           SecurityAmazonResourceEvidenceRolesscanned,
		"source":            SecurityAmazonResourceEvidenceRolessource,
		"suspicious":        SecurityAmazonResourceEvidenceRolessuspicious,
		"unknown":           SecurityAmazonResourceEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAmazonResourceEvidenceRoles(input)
	return &out, nil
}
