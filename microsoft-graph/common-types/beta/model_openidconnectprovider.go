package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectProvider struct {
	ClaimsMapping *ClaimsMapping                     `json:"claimsMapping,omitempty"`
	ClientId      *string                            `json:"clientId,omitempty"`
	ClientSecret  *string                            `json:"clientSecret,omitempty"`
	DomainHint    *string                            `json:"domainHint,omitempty"`
	Id            *string                            `json:"id,omitempty"`
	MetadataUrl   *string                            `json:"metadataUrl,omitempty"`
	Name          *string                            `json:"name,omitempty"`
	ODataType     *string                            `json:"@odata.type,omitempty"`
	ResponseMode  *OpenIdConnectProviderResponseMode `json:"responseMode,omitempty"`
	ResponseType  *OpenIdConnectProviderResponseType `json:"responseType,omitempty"`
	Scope         *string                            `json:"scope,omitempty"`
	Type          *string                            `json:"type,omitempty"`
}
