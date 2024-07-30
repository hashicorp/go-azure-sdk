package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScanDirection string

const (
	Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorAllFiles          Windows10EndpointProtectionConfigurationDefenderScanDirection = "monitorAllFiles"
	Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorIncomingFilesOnly Windows10EndpointProtectionConfigurationDefenderScanDirection = "monitorIncomingFilesOnly"
	Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorOutgoingFilesOnly Windows10EndpointProtectionConfigurationDefenderScanDirection = "monitorOutgoingFilesOnly"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScanDirection() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorAllFiles),
		string(Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorIncomingFilesOnly),
		string(Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorOutgoingFilesOnly),
	}
}

func (s *Windows10EndpointProtectionConfigurationDefenderScanDirection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10EndpointProtectionConfigurationDefenderScanDirection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10EndpointProtectionConfigurationDefenderScanDirection(input string) (*Windows10EndpointProtectionConfigurationDefenderScanDirection, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScanDirection{
		"monitorallfiles":          Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorAllFiles,
		"monitorincomingfilesonly": Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorIncomingFilesOnly,
		"monitoroutgoingfilesonly": Windows10EndpointProtectionConfigurationDefenderScanDirection_MonitorOutgoingFilesOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScanDirection(input)
	return &out, nil
}
