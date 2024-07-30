package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionProxiedDomainCollection struct {
	DisplayName    *string          `json:"displayName,omitempty"`
	ODataType      *string          `json:"@odata.type,omitempty"`
	ProxiedDomains *[]ProxiedDomain `json:"proxiedDomains,omitempty"`
}
