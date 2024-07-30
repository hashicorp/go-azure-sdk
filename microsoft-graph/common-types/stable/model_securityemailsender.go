package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEmailSender struct {
	DisplayName  *string `json:"displayName,omitempty"`
	DomainName   *string `json:"domainName,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
