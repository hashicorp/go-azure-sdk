package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDestination struct {
	DeviceCount        *int64                                      `json:"deviceCount,omitempty"`
	Fqdn               *string                                     `json:"fqdn,omitempty"`
	Ip                 *string                                     `json:"ip,omitempty"`
	LastAccessDateTime *string                                     `json:"lastAccessDateTime,omitempty"`
	NetworkingProtocol *NetworkaccessDestinationNetworkingProtocol `json:"networkingProtocol,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	Port               *int64                                      `json:"port,omitempty"`
	TrafficType        *NetworkaccessDestinationTrafficType        `json:"trafficType,omitempty"`
	TransactionCount   *int64                                      `json:"transactionCount,omitempty"`
	UserCount          *int64                                      `json:"userCount,omitempty"`
}
