package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyState struct {
	AssignmentFilterIds                  *[]string                                      `json:"assignmentFilterIds,omitempty"`
	DetectionState                       *DeviceHealthScriptPolicyStateDetectionState   `json:"detectionState,omitempty"`
	DeviceId                             *string                                        `json:"deviceId,omitempty"`
	DeviceName                           *string                                        `json:"deviceName,omitempty"`
	ExpectedStateUpdateDateTime          *string                                        `json:"expectedStateUpdateDateTime,omitempty"`
	Id                                   *string                                        `json:"id,omitempty"`
	LastStateUpdateDateTime              *string                                        `json:"lastStateUpdateDateTime,omitempty"`
	LastSyncDateTime                     *string                                        `json:"lastSyncDateTime,omitempty"`
	ODataType                            *string                                        `json:"@odata.type,omitempty"`
	OsVersion                            *string                                        `json:"osVersion,omitempty"`
	PolicyId                             *string                                        `json:"policyId,omitempty"`
	PolicyName                           *string                                        `json:"policyName,omitempty"`
	PostRemediationDetectionScriptError  *string                                        `json:"postRemediationDetectionScriptError,omitempty"`
	PostRemediationDetectionScriptOutput *string                                        `json:"postRemediationDetectionScriptOutput,omitempty"`
	PreRemediationDetectionScriptError   *string                                        `json:"preRemediationDetectionScriptError,omitempty"`
	PreRemediationDetectionScriptOutput  *string                                        `json:"preRemediationDetectionScriptOutput,omitempty"`
	RemediationScriptError               *string                                        `json:"remediationScriptError,omitempty"`
	RemediationState                     *DeviceHealthScriptPolicyStateRemediationState `json:"remediationState,omitempty"`
	UserName                             *string                                        `json:"userName,omitempty"`
}
