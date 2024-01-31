package entitytypes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsEntityProperties struct {
	AdditionalData        *map[string]interface{} `json:"additionalData,omitempty"`
	DnsServerIPEntityId   *string                 `json:"dnsServerIpEntityId,omitempty"`
	DomainName            *string                 `json:"domainName,omitempty"`
	FriendlyName          *string                 `json:"friendlyName,omitempty"`
	HostIPAddressEntityId *string                 `json:"hostIpAddressEntityId,omitempty"`
	IPAddressEntityIds    *[]string               `json:"ipAddressEntityIds,omitempty"`
}
