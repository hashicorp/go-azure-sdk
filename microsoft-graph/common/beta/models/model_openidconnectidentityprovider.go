package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenIdConnectIdentityProvider struct {
	ClaimsMapping *ClaimsMapping                             `json:"claimsMapping,omitempty"`
	ClientId      *string                                    `json:"clientId,omitempty"`
	ClientSecret  *string                                    `json:"clientSecret,omitempty"`
	DisplayName   *string                                    `json:"displayName,omitempty"`
	DomainHint    *string                                    `json:"domainHint,omitempty"`
	Id            *string                                    `json:"id,omitempty"`
	MetadataUrl   *string                                    `json:"metadataUrl,omitempty"`
	ODataType     *string                                    `json:"@odata.type,omitempty"`
	ResponseMode  *OpenIdConnectIdentityProviderResponseMode `json:"responseMode,omitempty"`
	ResponseType  *OpenIdConnectIdentityProviderResponseType `json:"responseType,omitempty"`
	Scope         *string                                    `json:"scope,omitempty"`
}
