package subscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionUpdateParameterProperties struct {
	AllowTracing   *bool              `json:"allowTracing,omitempty"`
	DisplayName    *string            `json:"displayName,omitempty"`
	ExpirationDate *string            `json:"expirationDate,omitempty"`
	OwnerId        *string            `json:"ownerId,omitempty"`
	PrimaryKey     *string            `json:"primaryKey,omitempty"`
	Scope          *string            `json:"scope,omitempty"`
	SecondaryKey   *string            `json:"secondaryKey,omitempty"`
	State          *SubscriptionState `json:"state,omitempty"`
	StateComment   *string            `json:"stateComment,omitempty"`
}
