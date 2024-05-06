package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfaceUpdate struct {
	DeviceKey   *int64             `json:"deviceKey,omitempty"`
	Name        *string            `json:"name,omitempty"`
	NetworkId   *string            `json:"networkId,omitempty"`
	NicType     *NICType           `json:"nicType,omitempty"`
	PowerOnBoot *PowerOnBootOption `json:"powerOnBoot,omitempty"`
}
