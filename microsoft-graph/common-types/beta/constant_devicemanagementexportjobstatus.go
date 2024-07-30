package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJobStatus string

const (
	DeviceManagementExportJobStatus_Completed  DeviceManagementExportJobStatus = "completed"
	DeviceManagementExportJobStatus_Failed     DeviceManagementExportJobStatus = "failed"
	DeviceManagementExportJobStatus_InProgress DeviceManagementExportJobStatus = "inProgress"
	DeviceManagementExportJobStatus_NotStarted DeviceManagementExportJobStatus = "notStarted"
	DeviceManagementExportJobStatus_Unknown    DeviceManagementExportJobStatus = "unknown"
)

func PossibleValuesForDeviceManagementExportJobStatus() []string {
	return []string{
		string(DeviceManagementExportJobStatus_Completed),
		string(DeviceManagementExportJobStatus_Failed),
		string(DeviceManagementExportJobStatus_InProgress),
		string(DeviceManagementExportJobStatus_NotStarted),
		string(DeviceManagementExportJobStatus_Unknown),
	}
}

func (s *DeviceManagementExportJobStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementExportJobStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementExportJobStatus(input string) (*DeviceManagementExportJobStatus, error) {
	vals := map[string]DeviceManagementExportJobStatus{
		"completed":  DeviceManagementExportJobStatus_Completed,
		"failed":     DeviceManagementExportJobStatus_Failed,
		"inprogress": DeviceManagementExportJobStatus_InProgress,
		"notstarted": DeviceManagementExportJobStatus_NotStarted,
		"unknown":    DeviceManagementExportJobStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementExportJobStatus(input)
	return &out, nil
}
