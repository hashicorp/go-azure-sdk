package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLoggedOnUser struct {
	AccountName *string `json:"accountName,omitempty"`
	DomainName  *string `json:"domainName,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
