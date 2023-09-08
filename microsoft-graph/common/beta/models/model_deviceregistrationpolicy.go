package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceRegistrationPolicy struct {
	AzureADJoin                  *AzureAdJoinPolicy                                    `json:"azureADJoin,omitempty"`
	AzureADRegistration          *AzureADRegistrationPolicy                            `json:"azureADRegistration,omitempty"`
	Description                  *string                                               `json:"description,omitempty"`
	DisplayName                  *string                                               `json:"displayName,omitempty"`
	Id                           *string                                               `json:"id,omitempty"`
	LocalAdminPassword           *LocalAdminPasswordSettings                           `json:"localAdminPassword,omitempty"`
	MultiFactorAuthConfiguration *DeviceRegistrationPolicyMultiFactorAuthConfiguration `json:"multiFactorAuthConfiguration,omitempty"`
	ODataType                    *string                                               `json:"@odata.type,omitempty"`
	UserDeviceQuota              *int64                                                `json:"userDeviceQuota,omitempty"`
}
