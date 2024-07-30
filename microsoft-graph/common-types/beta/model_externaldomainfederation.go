package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalDomainFederation struct {
	DisplayName *string `json:"displayName,omitempty"`
	DomainName  *string `json:"domainName,omitempty"`
	IssuerUri   *string `json:"issuerUri,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
