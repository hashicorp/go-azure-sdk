package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMDeviceState struct {
	CreatedDateTime                          *string                      `json:"createdDateTime,omitempty"`
	DeviceName                               *string                      `json:"deviceName,omitempty"`
	Id                                       *string                      `json:"id,omitempty"`
	LastSyncDateTime                         *string                      `json:"lastSyncDateTime,omitempty"`
	ModifiedDateTime                         *string                      `json:"modifiedDateTime,omitempty"`
	ODataType                                *string                      `json:"@odata.type,omitempty"`
	State                                    *EmbeddedSIMDeviceStateState `json:"state,omitempty"`
	StateDetails                             *string                      `json:"stateDetails,omitempty"`
	UniversalIntegratedCircuitCardIdentifier *string                      `json:"universalIntegratedCircuitCardIdentifier,omitempty"`
	UserName                                 *string                      `json:"userName,omitempty"`
}
