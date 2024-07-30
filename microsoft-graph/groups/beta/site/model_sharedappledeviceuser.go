package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedAppleDeviceUser struct {
	DataQuota         *int64  `json:"dataQuota,omitempty"`
	DataToSync        *bool   `json:"dataToSync,omitempty"`
	DataUsed          *int64  `json:"dataUsed,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}
