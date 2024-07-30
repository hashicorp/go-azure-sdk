package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedProvider struct {
	DisplayName                     *string                                             `json:"displayName,omitempty"`
	Id                              *string                                             `json:"id,omitempty"`
	IssuerUri                       *string                                             `json:"issuerUri,omitempty"`
	MetadataExchangeUri             *string                                             `json:"metadataExchangeUri,omitempty"`
	ODataType                       *string                                             `json:"@odata.type,omitempty"`
	PassiveSignInUri                *string                                             `json:"passiveSignInUri,omitempty"`
	PreferredAuthenticationProtocol *SamlOrWsFedProviderPreferredAuthenticationProtocol `json:"preferredAuthenticationProtocol,omitempty"`
	SigningCertificate              *string                                             `json:"signingCertificate,omitempty"`
}
