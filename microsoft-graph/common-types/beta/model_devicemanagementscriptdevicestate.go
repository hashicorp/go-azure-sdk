package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptDeviceState struct {
	ErrorCode               *int64                                     `json:"errorCode,omitempty"`
	ErrorDescription        *string                                    `json:"errorDescription,omitempty"`
	Id                      *string                                    `json:"id,omitempty"`
	LastStateUpdateDateTime *string                                    `json:"lastStateUpdateDateTime,omitempty"`
	ManagedDevice           *ManagedDevice                             `json:"managedDevice,omitempty"`
	ODataType               *string                                    `json:"@odata.type,omitempty"`
	ResultMessage           *string                                    `json:"resultMessage,omitempty"`
	RunState                *DeviceManagementScriptDeviceStateRunState `json:"runState,omitempty"`
}
