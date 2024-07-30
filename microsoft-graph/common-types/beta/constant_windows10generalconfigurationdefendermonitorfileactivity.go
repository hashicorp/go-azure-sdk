package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderMonitorFileActivity string

const (
	Windows10GeneralConfigurationDefenderMonitorFileActivity_Disable                  Windows10GeneralConfigurationDefenderMonitorFileActivity = "disable"
	Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorAllFiles          Windows10GeneralConfigurationDefenderMonitorFileActivity = "monitorAllFiles"
	Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorIncomingFilesOnly Windows10GeneralConfigurationDefenderMonitorFileActivity = "monitorIncomingFilesOnly"
	Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorOutgoingFilesOnly Windows10GeneralConfigurationDefenderMonitorFileActivity = "monitorOutgoingFilesOnly"
	Windows10GeneralConfigurationDefenderMonitorFileActivity_UserDefined              Windows10GeneralConfigurationDefenderMonitorFileActivity = "userDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderMonitorFileActivity() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderMonitorFileActivity_Disable),
		string(Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorAllFiles),
		string(Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorIncomingFilesOnly),
		string(Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorOutgoingFilesOnly),
		string(Windows10GeneralConfigurationDefenderMonitorFileActivity_UserDefined),
	}
}

func (s *Windows10GeneralConfigurationDefenderMonitorFileActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderMonitorFileActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderMonitorFileActivity(input string) (*Windows10GeneralConfigurationDefenderMonitorFileActivity, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderMonitorFileActivity{
		"disable":                  Windows10GeneralConfigurationDefenderMonitorFileActivity_Disable,
		"monitorallfiles":          Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorAllFiles,
		"monitorincomingfilesonly": Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorIncomingFilesOnly,
		"monitoroutgoingfilesonly": Windows10GeneralConfigurationDefenderMonitorFileActivity_MonitorOutgoingFilesOnly,
		"userdefined":              Windows10GeneralConfigurationDefenderMonitorFileActivity_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderMonitorFileActivity(input)
	return &out, nil
}
