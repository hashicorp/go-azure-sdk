package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InternalDomainFederation struct {
	ActiveSignInUri                       *string                                                  `json:"activeSignInUri,omitempty"`
	DisplayName                           *string                                                  `json:"displayName,omitempty"`
	FederatedIdpMfaBehavior               *InternalDomainFederationFederatedIdpMfaBehavior         `json:"federatedIdpMfaBehavior,omitempty"`
	Id                                    *string                                                  `json:"id,omitempty"`
	IsSignedAuthenticationRequestRequired *bool                                                    `json:"isSignedAuthenticationRequestRequired,omitempty"`
	IssuerUri                             *string                                                  `json:"issuerUri,omitempty"`
	MetadataExchangeUri                   *string                                                  `json:"metadataExchangeUri,omitempty"`
	NextSigningCertificate                *string                                                  `json:"nextSigningCertificate,omitempty"`
	ODataType                             *string                                                  `json:"@odata.type,omitempty"`
	PassiveSignInUri                      *string                                                  `json:"passiveSignInUri,omitempty"`
	PreferredAuthenticationProtocol       *InternalDomainFederationPreferredAuthenticationProtocol `json:"preferredAuthenticationProtocol,omitempty"`
	PromptLoginBehavior                   *InternalDomainFederationPromptLoginBehavior             `json:"promptLoginBehavior,omitempty"`
	SignOutUri                            *string                                                  `json:"signOutUri,omitempty"`
	SigningCertificate                    *string                                                  `json:"signingCertificate,omitempty"`
	SigningCertificateUpdateStatus        *SigningCertificateUpdateStatus                          `json:"signingCertificateUpdateStatus,omitempty"`
}
