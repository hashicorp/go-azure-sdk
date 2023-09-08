package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityKubernetesPodEvidenceRoles string

const (
	SecurityKubernetesPodEvidenceRolesadded             SecurityKubernetesPodEvidenceRoles = "Added"
	SecurityKubernetesPodEvidenceRolesattacked          SecurityKubernetesPodEvidenceRoles = "Attacked"
	SecurityKubernetesPodEvidenceRolesattacker          SecurityKubernetesPodEvidenceRoles = "Attacker"
	SecurityKubernetesPodEvidenceRolescommandAndControl SecurityKubernetesPodEvidenceRoles = "CommandAndControl"
	SecurityKubernetesPodEvidenceRolescompromised       SecurityKubernetesPodEvidenceRoles = "Compromised"
	SecurityKubernetesPodEvidenceRolescontextual        SecurityKubernetesPodEvidenceRoles = "Contextual"
	SecurityKubernetesPodEvidenceRolescreated           SecurityKubernetesPodEvidenceRoles = "Created"
	SecurityKubernetesPodEvidenceRolesdestination       SecurityKubernetesPodEvidenceRoles = "Destination"
	SecurityKubernetesPodEvidenceRolesedited            SecurityKubernetesPodEvidenceRoles = "Edited"
	SecurityKubernetesPodEvidenceRolesloaded            SecurityKubernetesPodEvidenceRoles = "Loaded"
	SecurityKubernetesPodEvidenceRolespolicyViolator    SecurityKubernetesPodEvidenceRoles = "PolicyViolator"
	SecurityKubernetesPodEvidenceRolesscanned           SecurityKubernetesPodEvidenceRoles = "Scanned"
	SecurityKubernetesPodEvidenceRolessource            SecurityKubernetesPodEvidenceRoles = "Source"
	SecurityKubernetesPodEvidenceRolessuspicious        SecurityKubernetesPodEvidenceRoles = "Suspicious"
	SecurityKubernetesPodEvidenceRolesunknown           SecurityKubernetesPodEvidenceRoles = "Unknown"
)

func PossibleValuesForSecurityKubernetesPodEvidenceRoles() []string {
	return []string{
		string(SecurityKubernetesPodEvidenceRolesadded),
		string(SecurityKubernetesPodEvidenceRolesattacked),
		string(SecurityKubernetesPodEvidenceRolesattacker),
		string(SecurityKubernetesPodEvidenceRolescommandAndControl),
		string(SecurityKubernetesPodEvidenceRolescompromised),
		string(SecurityKubernetesPodEvidenceRolescontextual),
		string(SecurityKubernetesPodEvidenceRolescreated),
		string(SecurityKubernetesPodEvidenceRolesdestination),
		string(SecurityKubernetesPodEvidenceRolesedited),
		string(SecurityKubernetesPodEvidenceRolesloaded),
		string(SecurityKubernetesPodEvidenceRolespolicyViolator),
		string(SecurityKubernetesPodEvidenceRolesscanned),
		string(SecurityKubernetesPodEvidenceRolessource),
		string(SecurityKubernetesPodEvidenceRolessuspicious),
		string(SecurityKubernetesPodEvidenceRolesunknown),
	}
}

func parseSecurityKubernetesPodEvidenceRoles(input string) (*SecurityKubernetesPodEvidenceRoles, error) {
	vals := map[string]SecurityKubernetesPodEvidenceRoles{
		"added":             SecurityKubernetesPodEvidenceRolesadded,
		"attacked":          SecurityKubernetesPodEvidenceRolesattacked,
		"attacker":          SecurityKubernetesPodEvidenceRolesattacker,
		"commandandcontrol": SecurityKubernetesPodEvidenceRolescommandAndControl,
		"compromised":       SecurityKubernetesPodEvidenceRolescompromised,
		"contextual":        SecurityKubernetesPodEvidenceRolescontextual,
		"created":           SecurityKubernetesPodEvidenceRolescreated,
		"destination":       SecurityKubernetesPodEvidenceRolesdestination,
		"edited":            SecurityKubernetesPodEvidenceRolesedited,
		"loaded":            SecurityKubernetesPodEvidenceRolesloaded,
		"policyviolator":    SecurityKubernetesPodEvidenceRolespolicyViolator,
		"scanned":           SecurityKubernetesPodEvidenceRolesscanned,
		"source":            SecurityKubernetesPodEvidenceRolessource,
		"suspicious":        SecurityKubernetesPodEvidenceRolessuspicious,
		"unknown":           SecurityKubernetesPodEvidenceRolesunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityKubernetesPodEvidenceRoles(input)
	return &out, nil
}
