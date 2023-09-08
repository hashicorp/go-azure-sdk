package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnumeratedDomains struct {
	DomainNames *[]string                     `json:"domainNames,omitempty"`
	ODataType   *string                       `json:"@odata.type,omitempty"`
	RootDomains *EnumeratedDomainsRootDomains `json:"rootDomains,omitempty"`
}
