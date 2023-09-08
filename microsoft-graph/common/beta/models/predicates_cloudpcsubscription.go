package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCSubscriptionOperationPredicate struct {
	ODataType        *string
	SubscriptionId   *string
	SubscriptionName *string
}

func (p CloudPCSubscriptionOperationPredicate) Matches(input CloudPCSubscription) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SubscriptionId != nil && (input.SubscriptionId == nil || *p.SubscriptionId != *input.SubscriptionId) {
		return false
	}

	if p.SubscriptionName != nil && (input.SubscriptionName == nil || *p.SubscriptionName != *input.SubscriptionName) {
		return false
	}

	return true
}
