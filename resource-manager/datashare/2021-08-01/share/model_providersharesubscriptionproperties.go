package share

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderShareSubscriptionProperties struct {
	ConsumerEmail             *string                  `json:"consumerEmail,omitempty"`
	ConsumerName              *string                  `json:"consumerName,omitempty"`
	ConsumerTenantName        *string                  `json:"consumerTenantName,omitempty"`
	CreatedAt                 *string                  `json:"createdAt,omitempty"`
	ExpirationDate            *string                  `json:"expirationDate,omitempty"`
	ProviderEmail             *string                  `json:"providerEmail,omitempty"`
	ProviderName              *string                  `json:"providerName,omitempty"`
	ShareSubscriptionObjectId *string                  `json:"shareSubscriptionObjectId,omitempty"`
	ShareSubscriptionStatus   *ShareSubscriptionStatus `json:"shareSubscriptionStatus,omitempty"`
	SharedAt                  *string                  `json:"sharedAt,omitempty"`
}
