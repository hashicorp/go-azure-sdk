package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceLink struct {
	BandwidthCapacityInMbps *NetworkaccessDeviceLinkBandwidthCapacityInMbps `json:"bandwidthCapacityInMbps,omitempty"`
	BgpConfiguration        *BgpConfiguration                               `json:"bgpConfiguration,omitempty"`
	DeviceVendor            *NetworkaccessDeviceLinkDeviceVendor            `json:"deviceVendor,omitempty"`
	Id                      *string                                         `json:"id,omitempty"`
	IpAddress               *string                                         `json:"ipAddress,omitempty"`
	LastModifiedDateTime    *string                                         `json:"lastModifiedDateTime,omitempty"`
	Name                    *string                                         `json:"name,omitempty"`
	ODataType               *string                                         `json:"@odata.type,omitempty"`
	RedundancyConfiguration *NetworkaccessRedundancyConfiguration           `json:"redundancyConfiguration,omitempty"`
	TunnelConfiguration     *TunnelConfiguration                            `json:"tunnelConfiguration,omitempty"`
}
