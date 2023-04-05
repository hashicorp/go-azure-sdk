package diagnosticsettings

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
