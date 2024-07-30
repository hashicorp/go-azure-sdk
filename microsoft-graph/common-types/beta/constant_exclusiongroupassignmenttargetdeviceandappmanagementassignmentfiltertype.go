package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExclusionGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
