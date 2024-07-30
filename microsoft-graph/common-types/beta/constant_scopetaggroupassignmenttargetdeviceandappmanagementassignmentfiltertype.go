package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None    ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude),
		string(ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include),
		string(ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None),
	}
}

func (s *ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Exclude,
		"include": ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_Include,
		"none":    ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScopeTagGroupAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
