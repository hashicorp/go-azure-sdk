package billingsubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Beneficiary struct {
	ObjectId *string `json:"objectId,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}
