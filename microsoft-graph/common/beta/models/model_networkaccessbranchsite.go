package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchSite struct {
	BandwidthCapacity         *int64                                        `json:"bandwidthCapacity,omitempty"`
	ConnectivityConfiguration *NetworkaccessBranchConnectivityConfiguration `json:"connectivityConfiguration,omitempty"`
	ConnectivityState         *NetworkaccessBranchSiteConnectivityState     `json:"connectivityState,omitempty"`
	Country                   *string                                       `json:"country,omitempty"`
	DeviceLinks               *[]NetworkaccessDeviceLink                    `json:"deviceLinks,omitempty"`
	ForwardingProfiles        *[]NetworkaccessForwardingProfile             `json:"forwardingProfiles,omitempty"`
	Id                        *string                                       `json:"id,omitempty"`
	LastModifiedDateTime      *string                                       `json:"lastModifiedDateTime,omitempty"`
	Name                      *string                                       `json:"name,omitempty"`
	ODataType                 *string                                       `json:"@odata.type,omitempty"`
	Region                    *NetworkaccessBranchSiteRegion                `json:"region,omitempty"`
	Version                   *string                                       `json:"version,omitempty"`
}
