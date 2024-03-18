package cloudlinks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudLinkProperties struct {
	LinkedCloud       *string                     `json:"linkedCloud,omitempty"`
	ProvisioningState *CloudLinkProvisioningState `json:"provisioningState,omitempty"`
	Status            *CloudLinkStatus            `json:"status,omitempty"`
}
