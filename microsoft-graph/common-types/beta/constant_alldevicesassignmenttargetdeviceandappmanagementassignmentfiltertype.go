package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForAllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllDevicesAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
