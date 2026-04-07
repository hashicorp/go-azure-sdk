package managednetwork

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedNetworkSettingsProperties struct {
	ManagedNetwork    *ManagedNetworkSettingsEx        `json:"managedNetwork,omitempty"`
	ProvisioningState *ManagedNetworkProvisioningState `json:"provisioningState,omitempty"`
}
