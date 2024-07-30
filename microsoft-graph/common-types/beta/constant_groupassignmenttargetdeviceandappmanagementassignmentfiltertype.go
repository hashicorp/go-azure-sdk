package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
