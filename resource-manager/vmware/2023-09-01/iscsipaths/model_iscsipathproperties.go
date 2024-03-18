package iscsipaths

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IscsiPathProperties struct {
	NetworkBlock      string                      `json:"networkBlock"`
	ProvisioningState *IscsiPathProvisioningState `json:"provisioningState,omitempty"`
}
