package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsNetworkIsolationPolicy struct {
	EnterpriseCloudResources               *[]ProxiedDomain `json:"enterpriseCloudResources,omitempty"`
	EnterpriseIPRanges                     *[]IpRange       `json:"enterpriseIPRanges,omitempty"`
	EnterpriseIPRangesAreAuthoritative     *bool            `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
	EnterpriseInternalProxyServers         *[]string        `json:"enterpriseInternalProxyServers,omitempty"`
	EnterpriseNetworkDomainNames           *[]string        `json:"enterpriseNetworkDomainNames,omitempty"`
	EnterpriseProxyServers                 *[]string        `json:"enterpriseProxyServers,omitempty"`
	EnterpriseProxyServersAreAuthoritative *bool            `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
	NeutralDomainResources                 *[]string        `json:"neutralDomainResources,omitempty"`
	ODataType                              *string          `json:"@odata.type,omitempty"`
}
