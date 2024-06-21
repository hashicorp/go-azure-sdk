package sharesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShareSubscriptionProperties struct {
	CreatedAt               *string                  `json:"createdAt,omitempty"`
	ExpirationDate          *string                  `json:"expirationDate,omitempty"`
	InvitationId            string                   `json:"invitationId"`
	ProviderEmail           *string                  `json:"providerEmail,omitempty"`
	ProviderName            *string                  `json:"providerName,omitempty"`
	ProviderTenantName      *string                  `json:"providerTenantName,omitempty"`
	ProvisioningState       *ProvisioningState       `json:"provisioningState,omitempty"`
	ShareDescription        *string                  `json:"shareDescription,omitempty"`
	ShareKind               *ShareKind               `json:"shareKind,omitempty"`
	ShareName               *string                  `json:"shareName,omitempty"`
	ShareSubscriptionStatus *ShareSubscriptionStatus `json:"shareSubscriptionStatus,omitempty"`
	ShareTerms              *string                  `json:"shareTerms,omitempty"`
	SourceShareLocation     string                   `json:"sourceShareLocation"`
	UserEmail               *string                  `json:"userEmail,omitempty"`
	UserName                *string                  `json:"userName,omitempty"`
}
