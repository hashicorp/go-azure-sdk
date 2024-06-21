package subscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionContractProperties struct {
	AllowTracing     *bool             `json:"allowTracing,omitempty"`
	CreatedDate      *string           `json:"createdDate,omitempty"`
	DisplayName      *string           `json:"displayName,omitempty"`
	EndDate          *string           `json:"endDate,omitempty"`
	ExpirationDate   *string           `json:"expirationDate,omitempty"`
	NotificationDate *string           `json:"notificationDate,omitempty"`
	OwnerId          *string           `json:"ownerId,omitempty"`
	PrimaryKey       *string           `json:"primaryKey,omitempty"`
	Scope            string            `json:"scope"`
	SecondaryKey     *string           `json:"secondaryKey,omitempty"`
	StartDate        *string           `json:"startDate,omitempty"`
	State            SubscriptionState `json:"state"`
	StateComment     *string           `json:"stateComment,omitempty"`
}
