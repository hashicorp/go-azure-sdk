package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Directory struct {
	AdministrativeUnits                *[]AdministrativeUnit                 `json:"administrativeUnits,omitempty"`
	AttributeSets                      *[]AttributeSet                       `json:"attributeSets,omitempty"`
	CertificateAuthorities             *CertificateAuthorityPath             `json:"certificateAuthorities,omitempty"`
	CustomSecurityAttributeDefinitions *[]CustomSecurityAttributeDefinition  `json:"customSecurityAttributeDefinitions,omitempty"`
	DeletedItems                       *[]DirectoryObject                    `json:"deletedItems,omitempty"`
	FeatureRolloutPolicies             *[]FeatureRolloutPolicy               `json:"featureRolloutPolicies,omitempty"`
	FederationConfigurations           *[]IdentityProviderBase               `json:"federationConfigurations,omitempty"`
	Id                                 *string                               `json:"id,omitempty"`
	ImpactedResources                  *[]ImpactedResource                   `json:"impactedResources,omitempty"`
	InboundSharedUserProfiles          *[]InboundSharedUserProfile           `json:"inboundSharedUserProfiles,omitempty"`
	ODataType                          *string                               `json:"@odata.type,omitempty"`
	OnPremisesSynchronization          *[]OnPremisesDirectorySynchronization `json:"onPremisesSynchronization,omitempty"`
	OutboundSharedUserProfiles         *[]OutboundSharedUserProfile          `json:"outboundSharedUserProfiles,omitempty"`
	Recommendations                    *[]Recommendation                     `json:"recommendations,omitempty"`
	SharedEmailDomains                 *[]SharedEmailDomain                  `json:"sharedEmailDomains,omitempty"`
	Subscriptions                      *[]CompanySubscription                `json:"subscriptions,omitempty"`
}
