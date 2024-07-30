package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptDeviceState struct {
	DetectionState              *DeviceComplianceScriptDeviceStateDetectionState `json:"detectionState,omitempty"`
	ExpectedStateUpdateDateTime *string                                          `json:"expectedStateUpdateDateTime,omitempty"`
	Id                          *string                                          `json:"id,omitempty"`
	LastStateUpdateDateTime     *string                                          `json:"lastStateUpdateDateTime,omitempty"`
	LastSyncDateTime            *string                                          `json:"lastSyncDateTime,omitempty"`
	ManagedDevice               *ManagedDevice                                   `json:"managedDevice,omitempty"`
	ODataType                   *string                                          `json:"@odata.type,omitempty"`
	ScriptError                 *string                                          `json:"scriptError,omitempty"`
	ScriptOutput                *string                                          `json:"scriptOutput,omitempty"`
}
