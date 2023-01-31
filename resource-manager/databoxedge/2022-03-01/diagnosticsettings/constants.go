package diagnosticsettings

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessLevel string

const (
	AccessLevelFullAccess AccessLevel = "FullAccess"
	AccessLevelNone       AccessLevel = "None"
	AccessLevelReadOnly   AccessLevel = "ReadOnly"
	AccessLevelReadWrite  AccessLevel = "ReadWrite"
)

func PossibleValuesForAccessLevel() []string {
	return []string{
		string(AccessLevelFullAccess),
		string(AccessLevelNone),
		string(AccessLevelReadOnly),
		string(AccessLevelReadWrite),
	}
}

func parseAccessLevel(input string) (*AccessLevel, error) {
	vals := map[string]AccessLevel{
		"fullaccess": AccessLevelFullAccess,
		"none":       AccessLevelNone,
		"readonly":   AccessLevelReadOnly,
		"readwrite":  AccessLevelReadWrite,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessLevel(input)
	return &out, nil
}

type ProactiveDiagnosticsConsent string

const (
	ProactiveDiagnosticsConsentDisabled ProactiveDiagnosticsConsent = "Disabled"
	ProactiveDiagnosticsConsentEnabled  ProactiveDiagnosticsConsent = "Enabled"
)

func PossibleValuesForProactiveDiagnosticsConsent() []string {
	return []string{
		string(ProactiveDiagnosticsConsentDisabled),
		string(ProactiveDiagnosticsConsentEnabled),
	}
}

func parseProactiveDiagnosticsConsent(input string) (*ProactiveDiagnosticsConsent, error) {
	vals := map[string]ProactiveDiagnosticsConsent{
		"disabled": ProactiveDiagnosticsConsentDisabled,
		"enabled":  ProactiveDiagnosticsConsentEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProactiveDiagnosticsConsent(input)
	return &out, nil
}

type RemoteApplicationType string

const (
	RemoteApplicationTypeAllApplications RemoteApplicationType = "AllApplications"
	RemoteApplicationTypeLocalUI         RemoteApplicationType = "LocalUI"
	RemoteApplicationTypePowershell      RemoteApplicationType = "Powershell"
	RemoteApplicationTypeWAC             RemoteApplicationType = "WAC"
)

func PossibleValuesForRemoteApplicationType() []string {
	return []string{
		string(RemoteApplicationTypeAllApplications),
		string(RemoteApplicationTypeLocalUI),
		string(RemoteApplicationTypePowershell),
		string(RemoteApplicationTypeWAC),
	}
}

func parseRemoteApplicationType(input string) (*RemoteApplicationType, error) {
	vals := map[string]RemoteApplicationType{
		"allapplications": RemoteApplicationTypeAllApplications,
		"localui":         RemoteApplicationTypeLocalUI,
		"powershell":      RemoteApplicationTypePowershell,
		"wac":             RemoteApplicationTypeWAC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RemoteApplicationType(input)
	return &out, nil
}
