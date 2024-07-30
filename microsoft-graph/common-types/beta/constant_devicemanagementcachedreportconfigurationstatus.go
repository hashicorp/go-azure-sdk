package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCachedReportConfigurationStatus string

const (
	DeviceManagementCachedReportConfigurationStatus_Completed  DeviceManagementCachedReportConfigurationStatus = "completed"
	DeviceManagementCachedReportConfigurationStatus_Failed     DeviceManagementCachedReportConfigurationStatus = "failed"
	DeviceManagementCachedReportConfigurationStatus_InProgress DeviceManagementCachedReportConfigurationStatus = "inProgress"
	DeviceManagementCachedReportConfigurationStatus_NotStarted DeviceManagementCachedReportConfigurationStatus = "notStarted"
	DeviceManagementCachedReportConfigurationStatus_Unknown    DeviceManagementCachedReportConfigurationStatus = "unknown"
)

func PossibleValuesForDeviceManagementCachedReportConfigurationStatus() []string {
	return []string{
		string(DeviceManagementCachedReportConfigurationStatus_Completed),
		string(DeviceManagementCachedReportConfigurationStatus_Failed),
		string(DeviceManagementCachedReportConfigurationStatus_InProgress),
		string(DeviceManagementCachedReportConfigurationStatus_NotStarted),
		string(DeviceManagementCachedReportConfigurationStatus_Unknown),
	}
}

func (s *DeviceManagementCachedReportConfigurationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementCachedReportConfigurationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementCachedReportConfigurationStatus(input string) (*DeviceManagementCachedReportConfigurationStatus, error) {
	vals := map[string]DeviceManagementCachedReportConfigurationStatus{
		"completed":  DeviceManagementCachedReportConfigurationStatus_Completed,
		"failed":     DeviceManagementCachedReportConfigurationStatus_Failed,
		"inprogress": DeviceManagementCachedReportConfigurationStatus_InProgress,
		"notstarted": DeviceManagementCachedReportConfigurationStatus_NotStarted,
		"unknown":    DeviceManagementCachedReportConfigurationStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCachedReportConfigurationStatus(input)
	return &out, nil
}
