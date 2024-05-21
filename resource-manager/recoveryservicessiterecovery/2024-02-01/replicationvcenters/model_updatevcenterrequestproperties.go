package replicationvcenters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVCenterRequestProperties struct {
	FriendlyName    *string `json:"friendlyName,omitempty"`
	IPAddress       *string `json:"ipAddress,omitempty"`
	Port            *string `json:"port,omitempty"`
	ProcessServerId *string `json:"processServerId,omitempty"`
	RunAsAccountId  *string `json:"runAsAccountId,omitempty"`
}
