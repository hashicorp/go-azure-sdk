package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AwsSource struct {
	AccountId            *string `json:"accountId,omitempty"`
	IdentityProviderType *string `json:"identityProviderType,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
}
