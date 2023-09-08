package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignIn struct {
	AppDisplayName                           *string                               `json:"appDisplayName,omitempty"`
	AppId                                    *string                               `json:"appId,omitempty"`
	AppTokenProtectionStatus                 *SignInAppTokenProtectionStatus       `json:"appTokenProtectionStatus,omitempty"`
	AppliedConditionalAccessPolicies         *[]AppliedConditionalAccessPolicy     `json:"appliedConditionalAccessPolicies,omitempty"`
	AppliedEventListeners                    *[]AppliedAuthenticationEventListener `json:"appliedEventListeners,omitempty"`
	AuthenticationAppDeviceDetails           *AuthenticationAppDeviceDetails       `json:"authenticationAppDeviceDetails,omitempty"`
	AuthenticationAppPolicyEvaluationDetails *[]AuthenticationAppPolicyDetails     `json:"authenticationAppPolicyEvaluationDetails,omitempty"`
	AuthenticationContextClassReferences     *[]AuthenticationContext              `json:"authenticationContextClassReferences,omitempty"`
	AuthenticationDetails                    *[]AuthenticationDetail               `json:"authenticationDetails,omitempty"`
	AuthenticationMethodsUsed                *[]string                             `json:"authenticationMethodsUsed,omitempty"`
	AuthenticationProcessingDetails          *[]KeyValue                           `json:"authenticationProcessingDetails,omitempty"`
	AuthenticationProtocol                   *SignInAuthenticationProtocol         `json:"authenticationProtocol,omitempty"`
	AuthenticationRequirement                *string                               `json:"authenticationRequirement,omitempty"`
	AuthenticationRequirementPolicies        *[]AuthenticationRequirementPolicy    `json:"authenticationRequirementPolicies,omitempty"`
	AutonomousSystemNumber                   *int64                                `json:"autonomousSystemNumber,omitempty"`
	AzureResourceId                          *string                               `json:"azureResourceId,omitempty"`
	ClientAppUsed                            *string                               `json:"clientAppUsed,omitempty"`
	ClientCredentialType                     *SignInClientCredentialType           `json:"clientCredentialType,omitempty"`
	ConditionalAccessStatus                  *SignInConditionalAccessStatus        `json:"conditionalAccessStatus,omitempty"`
	CorrelationId                            *string                               `json:"correlationId,omitempty"`
	CreatedDateTime                          *string                               `json:"createdDateTime,omitempty"`
	CrossTenantAccessType                    *SignInCrossTenantAccessType          `json:"crossTenantAccessType,omitempty"`
	DeviceDetail                             *DeviceDetail                         `json:"deviceDetail,omitempty"`
	FederatedCredentialId                    *string                               `json:"federatedCredentialId,omitempty"`
	FlaggedForReview                         *bool                                 `json:"flaggedForReview,omitempty"`
	HomeTenantId                             *string                               `json:"homeTenantId,omitempty"`
	HomeTenantName                           *string                               `json:"homeTenantName,omitempty"`
	Id                                       *string                               `json:"id,omitempty"`
	IncomingTokenType                        *SignInIncomingTokenType              `json:"incomingTokenType,omitempty"`
	IpAddress                                *string                               `json:"ipAddress,omitempty"`
	IpAddressFromResourceProvider            *string                               `json:"ipAddressFromResourceProvider,omitempty"`
	IsInteractive                            *bool                                 `json:"isInteractive,omitempty"`
	IsTenantRestricted                       *bool                                 `json:"isTenantRestricted,omitempty"`
	Location                                 *SignInLocation                       `json:"location,omitempty"`
	ManagedServiceIdentity                   *ManagedIdentity                      `json:"managedServiceIdentity,omitempty"`
	MfaDetail                                *MfaDetail                            `json:"mfaDetail,omitempty"`
	NetworkLocationDetails                   *[]NetworkLocationDetail              `json:"networkLocationDetails,omitempty"`
	ODataType                                *string                               `json:"@odata.type,omitempty"`
	OriginalRequestId                        *string                               `json:"originalRequestId,omitempty"`
	OriginalTransferMethod                   *SignInOriginalTransferMethod         `json:"originalTransferMethod,omitempty"`
	PrivateLinkDetails                       *PrivateLinkDetails                   `json:"privateLinkDetails,omitempty"`
	ProcessingTimeInMilliseconds             *int64                                `json:"processingTimeInMilliseconds,omitempty"`
	ResourceDisplayName                      *string                               `json:"resourceDisplayName,omitempty"`
	ResourceId                               *string                               `json:"resourceId,omitempty"`
	ResourceServicePrincipalId               *string                               `json:"resourceServicePrincipalId,omitempty"`
	ResourceTenantId                         *string                               `json:"resourceTenantId,omitempty"`
	RiskDetail                               *SignInRiskDetail                     `json:"riskDetail,omitempty"`
	RiskEventTypesv2                         *[]string                             `json:"riskEventTypes_v2,omitempty"`
	RiskLevelAggregated                      *SignInRiskLevelAggregated            `json:"riskLevelAggregated,omitempty"`
	RiskLevelDuringSignIn                    *SignInRiskLevelDuringSignIn          `json:"riskLevelDuringSignIn,omitempty"`
	RiskState                                *SignInRiskState                      `json:"riskState,omitempty"`
	ServicePrincipalCredentialKeyId          *string                               `json:"servicePrincipalCredentialKeyId,omitempty"`
	ServicePrincipalCredentialThumbprint     *string                               `json:"servicePrincipalCredentialThumbprint,omitempty"`
	ServicePrincipalId                       *string                               `json:"servicePrincipalId,omitempty"`
	ServicePrincipalName                     *string                               `json:"servicePrincipalName,omitempty"`
	SessionLifetimePolicies                  *[]SessionLifetimePolicy              `json:"sessionLifetimePolicies,omitempty"`
	SignInEventTypes                         *[]string                             `json:"signInEventTypes,omitempty"`
	SignInIdentifier                         *string                               `json:"signInIdentifier,omitempty"`
	SignInIdentifierType                     *SignInSignInIdentifierType           `json:"signInIdentifierType,omitempty"`
	SignInTokenProtectionStatus              *SignInSignInTokenProtectionStatus    `json:"signInTokenProtectionStatus,omitempty"`
	Status                                   *SignInStatus                         `json:"status,omitempty"`
	TokenIssuerName                          *string                               `json:"tokenIssuerName,omitempty"`
	TokenIssuerType                          *SignInTokenIssuerType                `json:"tokenIssuerType,omitempty"`
	UniqueTokenIdentifier                    *string                               `json:"uniqueTokenIdentifier,omitempty"`
	UserAgent                                *string                               `json:"userAgent,omitempty"`
	UserDisplayName                          *string                               `json:"userDisplayName,omitempty"`
	UserId                                   *string                               `json:"userId,omitempty"`
	UserPrincipalName                        *string                               `json:"userPrincipalName,omitempty"`
	UserType                                 *SignInUserType                       `json:"userType,omitempty"`
}
