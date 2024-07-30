package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureSource struct {
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	SubscriptionId       *string `json:"subscriptionId,omitempty"`
}
