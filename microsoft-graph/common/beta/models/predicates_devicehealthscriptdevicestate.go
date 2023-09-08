package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceStateOperationPredicate struct {
	ExpectedStateUpdateDateTime          *string
	Id                                   *string
	LastStateUpdateDateTime              *string
	LastSyncDateTime                     *string
	ODataType                            *string
	PostRemediationDetectionScriptError  *string
	PostRemediationDetectionScriptOutput *string
	PreRemediationDetectionScriptError   *string
	PreRemediationDetectionScriptOutput  *string
	RemediationScriptError               *string
}

func (p DeviceHealthScriptDeviceStateOperationPredicate) Matches(input DeviceHealthScriptDeviceState) bool {

	if p.ExpectedStateUpdateDateTime != nil && (input.ExpectedStateUpdateDateTime == nil || *p.ExpectedStateUpdateDateTime != *input.ExpectedStateUpdateDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastStateUpdateDateTime != nil && (input.LastStateUpdateDateTime == nil || *p.LastStateUpdateDateTime != *input.LastStateUpdateDateTime) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PostRemediationDetectionScriptError != nil && (input.PostRemediationDetectionScriptError == nil || *p.PostRemediationDetectionScriptError != *input.PostRemediationDetectionScriptError) {
		return false
	}

	if p.PostRemediationDetectionScriptOutput != nil && (input.PostRemediationDetectionScriptOutput == nil || *p.PostRemediationDetectionScriptOutput != *input.PostRemediationDetectionScriptOutput) {
		return false
	}

	if p.PreRemediationDetectionScriptError != nil && (input.PreRemediationDetectionScriptError == nil || *p.PreRemediationDetectionScriptError != *input.PreRemediationDetectionScriptError) {
		return false
	}

	if p.PreRemediationDetectionScriptOutput != nil && (input.PreRemediationDetectionScriptOutput == nil || *p.PreRemediationDetectionScriptOutput != *input.PreRemediationDetectionScriptOutput) {
		return false
	}

	if p.RemediationScriptError != nil && (input.RemediationScriptError == nil || *p.RemediationScriptError != *input.RemediationScriptError) {
		return false
	}

	return true
}
