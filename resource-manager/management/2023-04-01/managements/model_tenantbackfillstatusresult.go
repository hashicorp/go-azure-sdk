package managements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TenantBackfillStatusResult struct {
	Status   *Status `json:"status,omitempty"`
	TenantId *string `json:"tenantId,omitempty"`
}
