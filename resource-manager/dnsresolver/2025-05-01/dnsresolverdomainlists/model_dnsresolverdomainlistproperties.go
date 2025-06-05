package dnsresolverdomainlists

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsResolverDomainListProperties struct {
	Domains           *[]string          `json:"domains,omitempty"`
	DomainsURL        *string            `json:"domainsUrl,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
	ResourceGuid      *string            `json:"resourceGuid,omitempty"`
}
