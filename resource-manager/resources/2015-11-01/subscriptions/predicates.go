package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionOperationPredicate struct {
	DisplayName    *string
	Id             *string
	State          *string
	SubscriptionId *string
}

func (p SubscriptionOperationPredicate) Matches(input Subscription) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.State != nil && (input.State == nil || *p.State != *input.State) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil || *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	return true
}
