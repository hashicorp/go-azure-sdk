package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BmcKeySetProperties struct {
	AzureGroupId          string                      `json:"azureGroupId"`
	DetailedStatus        *BmcKeySetDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                     `json:"detailedStatusMessage,omitempty"`
	Expiration            string                      `json:"expiration"`
	LastValidation        *string                     `json:"lastValidation,omitempty"`
	PrivilegeLevel        BmcKeySetPrivilegeLevel     `json:"privilegeLevel"`
	ProvisioningState     *BmcKeySetProvisioningState `json:"provisioningState,omitempty"`
	UserList              []KeySetUser                `json:"userList"`
	UserListStatus        *[]KeySetUserStatus         `json:"userListStatus,omitempty"`
}
