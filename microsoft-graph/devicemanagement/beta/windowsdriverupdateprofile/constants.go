package windowsdriverupdateprofile

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsDriverUpdateProfileExecuteActionRequestActionName string

const (
	CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameApprove CreateWindowsDriverUpdateProfileExecuteActionRequestActionName = "approve"
	CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameDecline CreateWindowsDriverUpdateProfileExecuteActionRequestActionName = "decline"
	CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameSuspend CreateWindowsDriverUpdateProfileExecuteActionRequestActionName = "suspend"
)

func PossibleValuesForCreateWindowsDriverUpdateProfileExecuteActionRequestActionName() []string {
	return []string{
		string(CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameApprove),
		string(CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameDecline),
		string(CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameSuspend),
	}
}

func (s *CreateWindowsDriverUpdateProfileExecuteActionRequestActionName) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateWindowsDriverUpdateProfileExecuteActionRequestActionName(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateWindowsDriverUpdateProfileExecuteActionRequestActionName(input string) (*CreateWindowsDriverUpdateProfileExecuteActionRequestActionName, error) {
	vals := map[string]CreateWindowsDriverUpdateProfileExecuteActionRequestActionName{
		"approve": CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameApprove,
		"decline": CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameDecline,
		"suspend": CreateWindowsDriverUpdateProfileExecuteActionRequestActionNameSuspend,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateWindowsDriverUpdateProfileExecuteActionRequestActionName(input)
	return &out, nil
}

type DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType string

const (
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "exclude"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "include"
	DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType = "none"
)

func PossibleValuesForDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType() []string {
	return []string{
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude),
		string(DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone),
	}
}

func (s *DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input string) (*DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType, error) {
	vals := map[string]DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType{
		"exclude": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeExclude,
		"include": DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeInclude,
		"none":    DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementAssignmentTargetDeviceAndAppManagementAssignmentFilterType(input)
	return &out, nil
}
