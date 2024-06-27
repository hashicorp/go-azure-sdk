package agreement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileInfo struct {
	BillingAccountId                     *string `json:"billingAccountId,omitempty"`
	BillingProfileDisplayName            *string `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string `json:"billingProfileId,omitempty"`
	BillingProfileSystemId               *string `json:"billingProfileSystemId,omitempty"`
	IndirectRelationshipOrganizationName *string `json:"indirectRelationshipOrganizationName,omitempty"`
}
