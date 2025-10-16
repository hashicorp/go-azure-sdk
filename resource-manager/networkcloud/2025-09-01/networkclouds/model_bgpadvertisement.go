package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BgpAdvertisement struct {
	AdvertiseToFabric *AdvertiseToFabric `json:"advertiseToFabric,omitempty"`
	Communities       *[]string          `json:"communities,omitempty"`
	IPAddressPools    []string           `json:"ipAddressPools"`
	Peers             *[]string          `json:"peers,omitempty"`
}
