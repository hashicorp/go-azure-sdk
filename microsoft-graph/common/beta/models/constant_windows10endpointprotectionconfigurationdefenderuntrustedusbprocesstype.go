package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeauditMode   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeblock       Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "Block"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypedisable     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "Disable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeuserDefined Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypewarn        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeblock),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypedisable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypewarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType{
		"auditmode":   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeauditMode,
		"block":       Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeblock,
		"disable":     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypedisable,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypeuserDefined,
		"warn":        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessTypewarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessType(input)
	return &out, nil
}
