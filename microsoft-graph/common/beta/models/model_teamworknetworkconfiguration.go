package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkNetworkConfiguration struct {
	DefaultGateway  *string `json:"defaultGateway,omitempty"`
	DomainName      *string `json:"domainName,omitempty"`
	HostName        *string `json:"hostName,omitempty"`
	IpAddress       *string `json:"ipAddress,omitempty"`
	IsDhcpEnabled   *bool   `json:"isDhcpEnabled,omitempty"`
	IsPCPortEnabled *bool   `json:"isPCPortEnabled,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	PrimaryDns      *string `json:"primaryDns,omitempty"`
	SecondaryDns    *string `json:"secondaryDns,omitempty"`
	SubnetMask      *string `json:"subnetMask,omitempty"`
}
