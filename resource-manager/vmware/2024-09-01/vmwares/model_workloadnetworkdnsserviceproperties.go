package vmwares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDnsServiceProperties struct {
	DefaultDnsZone    *string                                     `json:"defaultDnsZone,omitempty"`
	DisplayName       *string                                     `json:"displayName,omitempty"`
	DnsServiceIP      *string                                     `json:"dnsServiceIp,omitempty"`
	FqdnZones         *[]string                                   `json:"fqdnZones,omitempty"`
	LogLevel          *DnsServiceLogLevelEnum                     `json:"logLevel,omitempty"`
	ProvisioningState *WorkloadNetworkDnsServiceProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                      `json:"revision,omitempty"`
	Status            *DnsServiceStatusEnum                       `json:"status,omitempty"`
}
