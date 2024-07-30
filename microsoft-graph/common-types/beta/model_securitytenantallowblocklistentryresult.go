package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowBlockListEntryResult struct {
	EntryType          *SecurityTenantAllowBlockListEntryResultEntryType `json:"entryType,omitempty"`
	ExpirationDateTime *string                                           `json:"expirationDateTime,omitempty"`
	Identity           *string                                           `json:"identity,omitempty"`
	ODataType          *string                                           `json:"@odata.type,omitempty"`
	Status             *SecurityTenantAllowBlockListEntryResultStatus    `json:"status,omitempty"`
	Value              *string                                           `json:"value,omitempty"`
}
