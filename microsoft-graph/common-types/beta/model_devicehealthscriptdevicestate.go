package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceState struct {
	AssignmentFilterIds                  *[]string                                      `json:"assignmentFilterIds,omitempty"`
	DetectionState                       *DeviceHealthScriptDeviceStateDetectionState   `json:"detectionState,omitempty"`
	ExpectedStateUpdateDateTime          *string                                        `json:"expectedStateUpdateDateTime,omitempty"`
	Id                                   *string                                        `json:"id,omitempty"`
	LastStateUpdateDateTime              *string                                        `json:"lastStateUpdateDateTime,omitempty"`
	LastSyncDateTime                     *string                                        `json:"lastSyncDateTime,omitempty"`
	ManagedDevice                        *ManagedDevice                                 `json:"managedDevice,omitempty"`
	ODataType                            *string                                        `json:"@odata.type,omitempty"`
	PostRemediationDetectionScriptError  *string                                        `json:"postRemediationDetectionScriptError,omitempty"`
	PostRemediationDetectionScriptOutput *string                                        `json:"postRemediationDetectionScriptOutput,omitempty"`
	PreRemediationDetectionScriptError   *string                                        `json:"preRemediationDetectionScriptError,omitempty"`
	PreRemediationDetectionScriptOutput  *string                                        `json:"preRemediationDetectionScriptOutput,omitempty"`
	RemediationScriptError               *string                                        `json:"remediationScriptError,omitempty"`
	RemediationState                     *DeviceHealthScriptDeviceStateRemediationState `json:"remediationState,omitempty"`
}
