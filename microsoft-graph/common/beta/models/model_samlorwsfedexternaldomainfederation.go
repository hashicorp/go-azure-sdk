package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamlOrWsFedExternalDomainFederation struct {
	DisplayName                     *string                                                             `json:"displayName,omitempty"`
	Domains                         *[]ExternalDomainName                                               `json:"domains,omitempty"`
	Id                              *string                                                             `json:"id,omitempty"`
	IssuerUri                       *string                                                             `json:"issuerUri,omitempty"`
	MetadataExchangeUri             *string                                                             `json:"metadataExchangeUri,omitempty"`
	ODataType                       *string                                                             `json:"@odata.type,omitempty"`
	PassiveSignInUri                *string                                                             `json:"passiveSignInUri,omitempty"`
	PreferredAuthenticationProtocol *SamlOrWsFedExternalDomainFederationPreferredAuthenticationProtocol `json:"preferredAuthenticationProtocol,omitempty"`
	SigningCertificate              *string                                                             `json:"signingCertificate,omitempty"`
}
