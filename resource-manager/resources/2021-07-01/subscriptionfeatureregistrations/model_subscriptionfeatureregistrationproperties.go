package subscriptionfeatureregistrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionFeatureRegistrationProperties struct {
	ApprovalType                 *SubscriptionFeatureRegistrationApprovalType `json:"approvalType,omitempty"`
	AuthorizationProfile         *AuthorizationProfile                        `json:"authorizationProfile,omitempty"`
	Description                  *string                                      `json:"description,omitempty"`
	DisplayName                  *string                                      `json:"displayName,omitempty"`
	DocumentationLink            *string                                      `json:"documentationLink,omitempty"`
	FeatureName                  *string                                      `json:"featureName,omitempty"`
	Metadata                     *map[string]string                           `json:"metadata,omitempty"`
	ProviderNamespace            *string                                      `json:"providerNamespace,omitempty"`
	RegistrationDate             *string                                      `json:"registrationDate,omitempty"`
	ReleaseDate                  *string                                      `json:"releaseDate,omitempty"`
	ShouldFeatureDisplayInPortal *bool                                        `json:"shouldFeatureDisplayInPortal,omitempty"`
	State                        *SubscriptionFeatureRegistrationState        `json:"state,omitempty"`
	SubscriptionId               *string                                      `json:"subscriptionId,omitempty"`
	TenantId                     *string                                      `json:"tenantId,omitempty"`
}
