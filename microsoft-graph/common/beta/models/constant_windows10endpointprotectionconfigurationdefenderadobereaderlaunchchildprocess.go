package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess string

const (
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessauditMode     Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "AuditMode"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessenable        Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "Enable"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessnotConfigured Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "NotConfigured"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessuserDefined   Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "UserDefined"
	Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcesswarn          Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess = "Warn"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessauditMode),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessenable),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessnotConfigured),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessuserDefined),
		string(Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcesswarn),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess(input string) (*Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess{
		"auditmode":     Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessauditMode,
		"enable":        Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessenable,
		"notconfigured": Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessnotConfigured,
		"userdefined":   Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcessuserDefined,
		"warn":          Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcesswarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderAdobeReaderLaunchChildProcess(input)
	return &out, nil
}
