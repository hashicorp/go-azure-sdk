package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScanDirection string

const (
	Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorAllFiles          Windows10EndpointProtectionConfigurationDefenderScanDirection = "MonitorAllFiles"
	Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorIncomingFilesOnly Windows10EndpointProtectionConfigurationDefenderScanDirection = "MonitorIncomingFilesOnly"
	Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorOutgoingFilesOnly Windows10EndpointProtectionConfigurationDefenderScanDirection = "MonitorOutgoingFilesOnly"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScanDirection() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorAllFiles),
		string(Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorIncomingFilesOnly),
		string(Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorOutgoingFilesOnly),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderScanDirection(input string) (*Windows10EndpointProtectionConfigurationDefenderScanDirection, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScanDirection{
		"monitorallfiles":          Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorAllFiles,
		"monitorincomingfilesonly": Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorIncomingFilesOnly,
		"monitoroutgoingfilesonly": Windows10EndpointProtectionConfigurationDefenderScanDirectionmonitorOutgoingFilesOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScanDirection(input)
	return &out, nil
}
