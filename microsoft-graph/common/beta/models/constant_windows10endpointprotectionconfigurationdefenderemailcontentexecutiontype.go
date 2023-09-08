package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType string

const (
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeauditMode   Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeblock       Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "Block"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypedisable     Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeuserDefined Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypewarn        Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderEmailContentExecutionType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderEmailContentExecutionType(input string) (*Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderEmailContentExecutionType(input)
	return &out, nil
}
