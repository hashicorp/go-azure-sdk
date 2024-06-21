package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetProperties struct {
	AzureGroupId          string                                   `json:"azureGroupId"`
	DetailedStatus        *BareMetalMachineKeySetDetailedStatus    `json:"detailedStatus,omitempty"`
	DetailedStatusMessage *string                                  `json:"detailedStatusMessage,omitempty"`
	Expiration            string                                   `json:"expiration"`
	JumpHostsAllowed      []string                                 `json:"jumpHostsAllowed"`
	LastValidation        *string                                  `json:"lastValidation,omitempty"`
	OsGroupName           *string                                  `json:"osGroupName,omitempty"`
	PrivilegeLevel        BareMetalMachineKeySetPrivilegeLevel     `json:"privilegeLevel"`
	ProvisioningState     *BareMetalMachineKeySetProvisioningState `json:"provisioningState,omitempty"`
	UserList              []KeySetUser                             `json:"userList"`
	UserListStatus        *[]KeySetUserStatus                      `json:"userListStatus,omitempty"`
}
