package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCRemoteActionResult struct {
	ActionName          *string                               `json:"actionName,omitempty"`
	ActionState         *CloudPCRemoteActionResultActionState `json:"actionState,omitempty"`
	CloudPCId           *string                               `json:"cloudPcId,omitempty"`
	LastUpdatedDateTime *string                               `json:"lastUpdatedDateTime,omitempty"`
	ManagedDeviceId     *string                               `json:"managedDeviceId,omitempty"`
	ODataType           *string                               `json:"@odata.type,omitempty"`
	StartDateTime       *string                               `json:"startDateTime,omitempty"`
	StatusDetails       *CloudPCStatusDetails                 `json:"statusDetails,omitempty"`
}
