package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeviceADAccount struct {
	DomainName *string `json:"domainName,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	Password   *string `json:"password,omitempty"`
	UserName   *string `json:"userName,omitempty"`
}
