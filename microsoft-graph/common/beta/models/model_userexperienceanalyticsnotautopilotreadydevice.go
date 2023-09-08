package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsNotAutopilotReadyDevice struct {
	AutoPilotProfileAssigned *bool   `json:"autoPilotProfileAssigned,omitempty"`
	AutoPilotRegistered      *bool   `json:"autoPilotRegistered,omitempty"`
	AzureAdJoinType          *string `json:"azureAdJoinType,omitempty"`
	AzureAdRegistered        *bool   `json:"azureAdRegistered,omitempty"`
	DeviceName               *string `json:"deviceName,omitempty"`
	Id                       *string `json:"id,omitempty"`
	ManagedBy                *string `json:"managedBy,omitempty"`
	Manufacturer             *string `json:"manufacturer,omitempty"`
	Model                    *string `json:"model,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	SerialNumber             *string `json:"serialNumber,omitempty"`
}
