package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicePrincipal struct {
	AccountEnabled                         *bool                                `json:"accountEnabled,omitempty"`
	AddIns                                 *[]AddIn                             `json:"addIns,omitempty"`
	AlternativeNames                       *[]string                            `json:"alternativeNames,omitempty"`
	AppDescription                         *string                              `json:"appDescription,omitempty"`
	AppDisplayName                         *string                              `json:"appDisplayName,omitempty"`
	AppId                                  *string                              `json:"appId,omitempty"`
	AppManagementPolicies                  *[]AppManagementPolicy               `json:"appManagementPolicies,omitempty"`
	AppOwnerOrganizationId                 *string                              `json:"appOwnerOrganizationId,omitempty"`
	AppRoleAssignedTo                      *[]AppRoleAssignment                 `json:"appRoleAssignedTo,omitempty"`
	AppRoleAssignmentRequired              *bool                                `json:"appRoleAssignmentRequired,omitempty"`
	AppRoleAssignments                     *[]AppRoleAssignment                 `json:"appRoleAssignments,omitempty"`
	AppRoles                               *[]AppRole                           `json:"appRoles,omitempty"`
	ApplicationTemplateId                  *string                              `json:"applicationTemplateId,omitempty"`
	ClaimsMappingPolicies                  *[]ClaimsMappingPolicy               `json:"claimsMappingPolicies,omitempty"`
	CreatedObjects                         *[]DirectoryObject                   `json:"createdObjects,omitempty"`
	CustomSecurityAttributes               *CustomSecurityAttributeValue        `json:"customSecurityAttributes,omitempty"`
	DelegatedPermissionClassifications     *[]DelegatedPermissionClassification `json:"delegatedPermissionClassifications,omitempty"`
	DeletedDateTime                        *string                              `json:"deletedDateTime,omitempty"`
	Description                            *string                              `json:"description,omitempty"`
	DisabledByMicrosoftStatus              *string                              `json:"disabledByMicrosoftStatus,omitempty"`
	DisplayName                            *string                              `json:"displayName,omitempty"`
	Endpoints                              *[]Endpoint                          `json:"endpoints,omitempty"`
	FederatedIdentityCredentials           *[]FederatedIdentityCredential       `json:"federatedIdentityCredentials,omitempty"`
	HomeRealmDiscoveryPolicies             *[]HomeRealmDiscoveryPolicy          `json:"homeRealmDiscoveryPolicies,omitempty"`
	Homepage                               *string                              `json:"homepage,omitempty"`
	Id                                     *string                              `json:"id,omitempty"`
	Info                                   *InformationalUrl                    `json:"info,omitempty"`
	KeyCredentials                         *[]KeyCredential                     `json:"keyCredentials,omitempty"`
	LoginUrl                               *string                              `json:"loginUrl,omitempty"`
	LogoutUrl                              *string                              `json:"logoutUrl,omitempty"`
	MemberOf                               *[]DirectoryObject                   `json:"memberOf,omitempty"`
	Notes                                  *string                              `json:"notes,omitempty"`
	NotificationEmailAddresses             *[]string                            `json:"notificationEmailAddresses,omitempty"`
	ODataType                              *string                              `json:"@odata.type,omitempty"`
	Oauth2PermissionGrants                 *[]OAuth2PermissionGrant             `json:"oauth2PermissionGrants,omitempty"`
	Oauth2PermissionScopes                 *[]PermissionScope                   `json:"oauth2PermissionScopes,omitempty"`
	OwnedObjects                           *[]DirectoryObject                   `json:"ownedObjects,omitempty"`
	Owners                                 *[]DirectoryObject                   `json:"owners,omitempty"`
	PasswordCredentials                    *[]PasswordCredential                `json:"passwordCredentials,omitempty"`
	PreferredSingleSignOnMode              *string                              `json:"preferredSingleSignOnMode,omitempty"`
	PreferredTokenSigningKeyThumbprint     *string                              `json:"preferredTokenSigningKeyThumbprint,omitempty"`
	ReplyUrls                              *[]string                            `json:"replyUrls,omitempty"`
	ResourceSpecificApplicationPermissions *[]ResourceSpecificPermission        `json:"resourceSpecificApplicationPermissions,omitempty"`
	SamlSingleSignOnSettings               *SamlSingleSignOnSettings            `json:"samlSingleSignOnSettings,omitempty"`
	ServicePrincipalNames                  *[]string                            `json:"servicePrincipalNames,omitempty"`
	ServicePrincipalType                   *string                              `json:"servicePrincipalType,omitempty"`
	SignInAudience                         *string                              `json:"signInAudience,omitempty"`
	Synchronization                        *Synchronization                     `json:"synchronization,omitempty"`
	Tags                                   *[]string                            `json:"tags,omitempty"`
	TokenEncryptionKeyId                   *string                              `json:"tokenEncryptionKeyId,omitempty"`
	TokenIssuancePolicies                  *[]TokenIssuancePolicy               `json:"tokenIssuancePolicies,omitempty"`
	TokenLifetimePolicies                  *[]TokenLifetimePolicy               `json:"tokenLifetimePolicies,omitempty"`
	TransitiveMemberOf                     *[]DirectoryObject                   `json:"transitiveMemberOf,omitempty"`
	VerifiedPublisher                      *VerifiedPublisher                   `json:"verifiedPublisher,omitempty"`
}
