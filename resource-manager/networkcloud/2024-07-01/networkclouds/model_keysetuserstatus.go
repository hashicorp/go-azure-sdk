package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeySetUserStatus struct {
	AzureUserName *string                                `json:"azureUserName,omitempty"`
	Status        *BareMetalMachineKeySetUserSetupStatus `json:"status,omitempty"`
	StatusMessage *string                                `json:"statusMessage,omitempty"`
}
