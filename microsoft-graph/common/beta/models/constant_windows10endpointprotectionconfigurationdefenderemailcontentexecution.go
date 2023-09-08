package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderEmailContentExecution string

const (
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionauditMode     Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionenable        Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "Enable"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionnotConfigured Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionuserDefined   Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionwarn          Windows10EndpointProtectionConfigurationDefenderEmailContentExecution = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderEmailContentExecution() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionenable),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionwarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecution(input string) (*Windows10EndpointProtectionConfigurationDefenderEmailContentExecution, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderEmailContentExecution{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionwarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderEmailContentExecution(input)
	return &out, nil
}
