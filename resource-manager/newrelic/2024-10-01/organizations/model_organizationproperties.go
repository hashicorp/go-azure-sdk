package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationProperties struct {
	BillingSource    *BillingSource `json:"billingSource,omitempty"`
	OrganizationId   *string        `json:"organizationId,omitempty"`
	OrganizationName *string        `json:"organizationName,omitempty"`
}
