package billingsubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionMergeRequest struct {
	Quantity                      *int64  `json:"quantity,omitempty"`
	TargetBillingSubscriptionName *string `json:"targetBillingSubscriptionName,omitempty"`
}
