package apisbytag

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductTagResourceContractProperties struct {
	ApprovalRequired     *bool         `json:"approvalRequired,omitempty"`
	Description          *string       `json:"description,omitempty"`
	Id                   *string       `json:"id,omitempty"`
	Name                 string        `json:"name"`
	State                *ProductState `json:"state,omitempty"`
	SubscriptionRequired *bool         `json:"subscriptionRequired,omitempty"`
	SubscriptionsLimit   *int64        `json:"subscriptionsLimit,omitempty"`
	Terms                *string       `json:"terms,omitempty"`
}
