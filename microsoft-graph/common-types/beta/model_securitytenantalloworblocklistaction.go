package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityTenantAllowOrBlockListAction struct {
	Action             *SecurityTenantAllowOrBlockListActionAction `json:"action,omitempty"`
	ExpirationDateTime *string                                     `json:"expirationDateTime,omitempty"`
	Note               *string                                     `json:"note,omitempty"`
	ODataType          *string                                     `json:"@odata.type,omitempty"`
	Results            *[]SecurityTenantAllowBlockListEntryResult  `json:"results,omitempty"`
}
