package usersubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionState string

const (
	SubscriptionStateActive    SubscriptionState = "active"
	SubscriptionStateCancelled SubscriptionState = "cancelled"
	SubscriptionStateExpired   SubscriptionState = "expired"
	SubscriptionStateRejected  SubscriptionState = "rejected"
	SubscriptionStateSubmitted SubscriptionState = "submitted"
	SubscriptionStateSuspended SubscriptionState = "suspended"
)

func PossibleValuesForSubscriptionState() []string {
	return []string{
		string(SubscriptionStateActive),
		string(SubscriptionStateCancelled),
		string(SubscriptionStateExpired),
		string(SubscriptionStateRejected),
		string(SubscriptionStateSubmitted),
		string(SubscriptionStateSuspended),
	}
}
