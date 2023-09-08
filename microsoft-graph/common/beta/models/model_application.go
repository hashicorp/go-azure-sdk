package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Application struct {
	Api                               *ApiApplication                    `json:"api,omitempty"`
	AppId                             *string                            `json:"appId,omitempty"`
	AppManagementPolicies             *[]AppManagementPolicy             `json:"appManagementPolicies,omitempty"`
	AppRoles                          *[]AppRole                         `json:"appRoles,omitempty"`
	AuthenticationBehaviors           *AuthenticationBehaviors           `json:"authenticationBehaviors,omitempty"`
	Certification                     *Certification                     `json:"certification,omitempty"`
	ConnectorGroup                    *ConnectorGroup                    `json:"connectorGroup,omitempty"`
	CreatedDateTime                   *string                            `json:"createdDateTime,omitempty"`
	CreatedOnBehalfOf                 *DirectoryObject                   `json:"createdOnBehalfOf,omitempty"`
	DefaultRedirectUri                *string                            `json:"defaultRedirectUri,omitempty"`
	DeletedDateTime                   *string                            `json:"deletedDateTime,omitempty"`
	Description                       *string                            `json:"description,omitempty"`
	DisabledByMicrosoftStatus         *string                            `json:"disabledByMicrosoftStatus,omitempty"`
	DisplayName                       *string                            `json:"displayName,omitempty"`
	ExtensionProperties               *[]ExtensionProperty               `json:"extensionProperties,omitempty"`
	FederatedIdentityCredentials      *[]FederatedIdentityCredential     `json:"federatedIdentityCredentials,omitempty"`
	GroupMembershipClaims             *string                            `json:"groupMembershipClaims,omitempty"`
	HomeRealmDiscoveryPolicies        *[]HomeRealmDiscoveryPolicy        `json:"homeRealmDiscoveryPolicies,omitempty"`
	Id                                *string                            `json:"id,omitempty"`
	IdentifierUris                    *[]string                          `json:"identifierUris,omitempty"`
	Info                              *InformationalUrl                  `json:"info,omitempty"`
	IsDeviceOnlyAuthSupported         *bool                              `json:"isDeviceOnlyAuthSupported,omitempty"`
	IsFallbackPublicClient            *bool                              `json:"isFallbackPublicClient,omitempty"`
	KeyCredentials                    *[]KeyCredential                   `json:"keyCredentials,omitempty"`
	Logo                              *string                            `json:"logo,omitempty"`
	Notes                             *string                            `json:"notes,omitempty"`
	ODataType                         *string                            `json:"@odata.type,omitempty"`
	OnPremisesPublishing              *OnPremisesPublishing              `json:"onPremisesPublishing,omitempty"`
	OptionalClaims                    *OptionalClaims                    `json:"optionalClaims,omitempty"`
	Owners                            *[]DirectoryObject                 `json:"owners,omitempty"`
	ParentalControlSettings           *ParentalControlSettings           `json:"parentalControlSettings,omitempty"`
	PasswordCredentials               *[]PasswordCredential              `json:"passwordCredentials,omitempty"`
	PublicClient                      *PublicClientApplication           `json:"publicClient,omitempty"`
	PublisherDomain                   *string                            `json:"publisherDomain,omitempty"`
	RequestSignatureVerification      *RequestSignatureVerification      `json:"requestSignatureVerification,omitempty"`
	RequiredResourceAccess            *[]RequiredResourceAccess          `json:"requiredResourceAccess,omitempty"`
	SamlMetadataUrl                   *string                            `json:"samlMetadataUrl,omitempty"`
	ServiceManagementReference        *string                            `json:"serviceManagementReference,omitempty"`
	ServicePrincipalLockConfiguration *ServicePrincipalLockConfiguration `json:"servicePrincipalLockConfiguration,omitempty"`
	SignInAudience                    *string                            `json:"signInAudience,omitempty"`
	Spa                               *SpaApplication                    `json:"spa,omitempty"`
	Synchronization                   *Synchronization                   `json:"synchronization,omitempty"`
	Tags                              *[]string                          `json:"tags,omitempty"`
	TokenEncryptionKeyId              *string                            `json:"tokenEncryptionKeyId,omitempty"`
	TokenIssuancePolicies             *[]TokenIssuancePolicy             `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies             *[]TokenLifetimePolicy             `json:"tokenLifetimePolicies,omitempty"`
	UniqueName                        *string                            `json:"uniqueName,omitempty"`
	VerifiedPublisher                 *VerifiedPublisher                 `json:"verifiedPublisher,omitempty"`
	Web                               *WebApplication                    `json:"web,omitempty"`
	Windows                           *WindowsApplication                `json:"windows,omitempty"`
}
