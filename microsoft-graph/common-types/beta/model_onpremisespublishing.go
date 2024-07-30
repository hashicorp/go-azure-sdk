package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishing struct {
	AlternateUrl                             *string                                         `json:"alternateUrl,omitempty"`
	ApplicationServerTimeout                 *string                                         `json:"applicationServerTimeout,omitempty"`
	ApplicationType                          *string                                         `json:"applicationType,omitempty"`
	ExternalAuthenticationType               *OnPremisesPublishingExternalAuthenticationType `json:"externalAuthenticationType,omitempty"`
	ExternalUrl                              *string                                         `json:"externalUrl,omitempty"`
	InternalUrl                              *string                                         `json:"internalUrl,omitempty"`
	IsAccessibleViaZTNAClient                *bool                                           `json:"isAccessibleViaZTNAClient,omitempty"`
	IsBackendCertificateValidationEnabled    *bool                                           `json:"isBackendCertificateValidationEnabled,omitempty"`
	IsDnsResolutionEnabled                   *bool                                           `json:"isDnsResolutionEnabled,omitempty"`
	IsHttpOnlyCookieEnabled                  *bool                                           `json:"isHttpOnlyCookieEnabled,omitempty"`
	IsOnPremPublishingEnabled                *bool                                           `json:"isOnPremPublishingEnabled,omitempty"`
	IsPersistentCookieEnabled                *bool                                           `json:"isPersistentCookieEnabled,omitempty"`
	IsSecureCookieEnabled                    *bool                                           `json:"isSecureCookieEnabled,omitempty"`
	IsStateSessionEnabled                    *bool                                           `json:"isStateSessionEnabled,omitempty"`
	IsTranslateHostHeaderEnabled             *bool                                           `json:"isTranslateHostHeaderEnabled,omitempty"`
	IsTranslateLinksInBodyEnabled            *bool                                           `json:"isTranslateLinksInBodyEnabled,omitempty"`
	ODataType                                *string                                         `json:"@odata.type,omitempty"`
	OnPremisesApplicationSegments            *[]OnPremisesApplicationSegment                 `json:"onPremisesApplicationSegments,omitempty"`
	SegmentsConfiguration                    *SegmentConfiguration                           `json:"segmentsConfiguration,omitempty"`
	SingleSignOnSettings                     *OnPremisesPublishingSingleSignOn               `json:"singleSignOnSettings,omitempty"`
	UseAlternateUrlForTranslationAndRedirect *bool                                           `json:"useAlternateUrlForTranslationAndRedirect,omitempty"`
	VerifiedCustomDomainCertificatesMetadata *VerifiedCustomDomainCertificatesMetadata       `json:"verifiedCustomDomainCertificatesMetadata,omitempty"`
	VerifiedCustomDomainKeyCredential        *KeyCredential                                  `json:"verifiedCustomDomainKeyCredential,omitempty"`
	VerifiedCustomDomainPasswordCredential   *PasswordCredential                             `json:"verifiedCustomDomainPasswordCredential,omitempty"`
}
