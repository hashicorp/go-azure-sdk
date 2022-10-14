package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkDnsZoneProperties struct {
	DisplayName       *string                                  `json:"displayName,omitempty"`
	DnsServerIPs      *[]string                                `json:"dnsServerIps,omitempty"`
	DnsServices       *int64                                   `json:"dnsServices,omitempty"`
	Domain            *[]string                                `json:"domain,omitempty"`
	ProvisioningState *WorkloadNetworkDnsZoneProvisioningState `json:"provisioningState,omitempty"`
	Revision          *int64                                   `json:"revision,omitempty"`
	SourceIP          *string                                  `json:"sourceIp,omitempty"`
}
