package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Subscription struct {
	DisplayName          *string               `json:"displayName,omitempty"`
	Id                   *string               `json:"id,omitempty"`
	State                *string               `json:"state,omitempty"`
	SubscriptionId       *string               `json:"subscriptionId,omitempty"`
	SubscriptionPolicies *SubscriptionPolicies `json:"subscriptionPolicies,omitempty"`
}
