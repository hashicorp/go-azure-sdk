package officeconsents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeConsentProperties struct {
	ConsentId *string `json:"consentId,omitempty"`
	TenantId  *string `json:"tenantId,omitempty"`
}
