package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessauditMode     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessenable        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "Enable"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessnotConfigured Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessuserDefined   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcesswarn          Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessenable),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcesswarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcessuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcesswarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderUntrustedUSBProcess(input)
	return &out, nil
}
