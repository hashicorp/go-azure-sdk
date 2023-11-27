package domains

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostName struct {
	AzureResourceName           *string                      `json:"azureResourceName,omitempty"`
	AzureResourceType           *AzureResourceType           `json:"azureResourceType,omitempty"`
	CustomHostNameDnsRecordType *CustomHostNameDnsRecordType `json:"customHostNameDnsRecordType,omitempty"`
	HostNameType                *HostNameType                `json:"hostNameType,omitempty"`
	Name                        *string                      `json:"name,omitempty"`
	SiteNames                   *[]string                    `json:"siteNames,omitempty"`
}
