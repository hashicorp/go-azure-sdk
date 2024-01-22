package agreements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingProfileInfo struct {
	BillingProfileDisplayName            *string `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string `json:"billingProfileId,omitempty"`
	IndirectRelationshipOrganizationName *string `json:"indirectRelationshipOrganizationName,omitempty"`
}
