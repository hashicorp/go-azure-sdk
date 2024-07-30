package devicehealthscript

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptAssignment struct {
	Id                   *string                                 `json:"id,omitempty"`
	ODataType            *string                                 `json:"@odata.type,omitempty"`
	RunRemediationScript *bool                                   `json:"runRemediationScript,omitempty"`
	RunSchedule          *DeviceHealthScriptRunSchedule          `json:"runSchedule,omitempty"`
	Target               *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
