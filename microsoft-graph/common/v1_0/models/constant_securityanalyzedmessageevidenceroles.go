package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAnalyzedMessageEvidenceRoles string

const (
	SecurityAnalyzedMessageEvidenceRolesadded             SecurityAnalyzedMessageEvidenceRoles = "Added"
	SecurityAnalyzedMessageEvidenceRolesattacked          SecurityAnalyzedMessageEvidenceRoles = "Attacked"
	SecurityAnalyzedMessageEvidenceRolesattacker          SecurityAnalyzedMessageEvidenceRoles = "Attacker"
	SecurityAnalyzedMessageEvidenceRolescommandAndControl SecurityAnalyzedMessageEvidenceRoles = "CommandAndControl"
	SecurityAnalyzedMessageEvidenceRolescompromised       SecurityAnalyzedMessageEvidenceRoles = "Compromised"
	SecurityAnalyzedMessageEvidenceRolescontextual        SecurityAnalyzedMessageEvidenceRoles = "Contextual"
	SecurityAnalyzedMessageEvidenceRolescreated           SecurityAnalyzedMessageEvidenceRoles = "Created"
	SecurityAnalyzedMessageEvidenceRolesdestination       SecurityAnalyzedMessageEvidenceRoles = "Destination"
	SecurityAnalyzedMessageEvidenceRolesedited            SecurityAnalyzedMessageEvidenceRoles = "Edited"
	SecurityAnalyzedMessageEvidenceRolesloaded            SecurityAnalyzedMessageEvidenceRoles = "Loaded"
	SecurityAnalyzedMessageEvidenceRolespolicyViolator    SecurityAnalyzedMessageEvidenceRoles = "PolicyViolator"
	SecurityAnalyzedMessageEvidenceRolesscanned           SecurityAnalyzedMessageEvidenceRoles = "Scanned"
	SecurityAnalyzedMessageEvidenceRolessource            SecurityAnalyzedMessageEvidenceRoles = "Source"
	SecurityAnalyzedMessageEvidenceRolessuspicious        SecurityAnalyzedMessageEvidenceRoles = "Suspicious"
	SecurityAnalyzedMessageEvidenceRolesunknown           SecurityAnalyzedMessageEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityAnalyzedMessageEvidenceRoles() []string {
	return []string{
		string(SecurityAnalyzedMessageEvidenceRolesadded),
		string(SecurityAnalyzedMessageEvidenceRolesattacked),
		string(SecurityAnalyzedMessageEvidenceRolesattacker),
		string(SecurityAnalyzedMessageEvidenceRolescommandAndControl),
		string(SecurityAnalyzedMessageEvidenceRolescompromised),
		string(SecurityAnalyzedMessageEvidenceRolescontextual),
		string(SecurityAnalyzedMessageEvidenceRolescreated),
		string(SecurityAnalyzedMessageEvidenceRolesdestination),
		string(SecurityAnalyzedMessageEvidenceRolesedited),
		string(SecurityAnalyzedMessageEvidenceRolesloaded),
		string(SecurityAnalyzedMessageEvidenceRolespolicyViolator),
		string(SecurityAnalyzedMessageEvidenceRolesscanned),
		string(SecurityAnalyzedMessageEvidenceRolessource),
		string(SecurityAnalyzedMessageEvidenceRolessuspicious),
		string(SecurityAnalyzedMessageEvidenceRolesunknown),
	}
}

func parseSecurityAnalyzedMessageEvidenceRoles(input string) (*SecurityAnalyzedMessageEvidenceRoles, error) {
	vals := map[string]SecurityAnalyzedMessageEvidenceRoles{
		"added":             SecurityAnalyzedMessageEvidenceRolesadded,
		"attacked":          SecurityAnalyzedMessageEvidenceRolesattacked,
		"attacker":          SecurityAnalyzedMessageEvidenceRolesattacker,
		"commandandcontrol": SecurityAnalyzedMessageEvidenceRolescommandAndControl,
		"compromised":       SecurityAnalyzedMessageEvidenceRolescompromised,
		"contextual":        SecurityAnalyzedMessageEvidenceRolescontextual,
		"created":           SecurityAnalyzedMessageEvidenceRolescreated,
		"destination":       SecurityAnalyzedMessageEvidenceRolesdestination,
		"edited":            SecurityAnalyzedMessageEvidenceRolesedited,
		"loaded":            SecurityAnalyzedMessageEvidenceRolesloaded,
		"policyviolator":    SecurityAnalyzedMessageEvidenceRolespolicyViolator,
		"scanned":           SecurityAnalyzedMessageEvidenceRolesscanned,
		"source":            SecurityAnalyzedMessageEvidenceRolessource,
		"suspicious":        SecurityAnalyzedMessageEvidenceRolessuspicious,
		"unknown":           SecurityAnalyzedMessageEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAnalyzedMessageEvidenceRoles(input)
	return &out, nil
}
