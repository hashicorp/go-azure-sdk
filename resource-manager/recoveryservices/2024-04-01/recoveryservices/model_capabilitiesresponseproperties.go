package recoveryservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapabilitiesResponseProperties struct {
	DnsZones *[]DNSZoneResponse `json:"dnsZones,omitempty"`
}
