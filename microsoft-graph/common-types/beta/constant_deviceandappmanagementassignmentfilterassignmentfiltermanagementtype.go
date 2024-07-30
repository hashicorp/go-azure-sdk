package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType string

const (
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Apps    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "apps"
	DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Devices DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType = "devices"
)

func PossibleValuesForDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Apps),
		string(DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Devices),
	}
}

func (s *DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input string) (*DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType{
		"apps":    DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Apps,
		"devices": DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType_Devices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentFilterAssignmentFilterManagementType(input)
	return &out, nil
}
