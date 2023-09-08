package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType string

const (
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeauditMode   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeblock       Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "Block"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypedisable     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeuserDefined Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypewarn        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType(input string) (*Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScriptDownloadedPayloadExecutionType(input)
	return &out, nil
}
