package containerappssourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureCredentials struct {
	ClientId       *string `json:"clientId,omitempty"`
	ClientSecret   *string `json:"clientSecret,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`
	TenantId       *string `json:"tenantId,omitempty"`
}
