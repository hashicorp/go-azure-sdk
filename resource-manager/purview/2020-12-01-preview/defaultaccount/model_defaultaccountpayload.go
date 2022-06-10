package defaultaccount

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultAccountPayload struct {
	AccountName       *string    `json:"accountName,omitempty"`
	ResourceGroupName *string    `json:"resourceGroupName,omitempty"`
	Scope             *string    `json:"scope,omitempty"`
	ScopeTenantId     *string    `json:"scopeTenantId,omitempty"`
	ScopeType         *ScopeType `json:"scopeType,omitempty"`
	SubscriptionId    *string    `json:"subscriptionId,omitempty"`
}
