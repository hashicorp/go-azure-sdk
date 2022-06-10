package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActiveDirectoryProperties struct {
	AzureStorageSid   string `json:"azureStorageSid"`
	DomainGuid        string `json:"domainGuid"`
	DomainName        string `json:"domainName"`
	DomainSid         string `json:"domainSid"`
	ForestName        string `json:"forestName"`
	NetBiosDomainName string `json:"netBiosDomainName"`
}
