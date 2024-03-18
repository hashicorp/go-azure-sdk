package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccountProperties struct {
	AccountId      *string `json:"accountId,omitempty"`
	AccountName    *string `json:"accountName,omitempty"`
	OrganizationId *string `json:"organizationId,omitempty"`
	Region         *string `json:"region,omitempty"`
}
