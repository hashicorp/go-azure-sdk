package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogonUser struct {
	AccountDomain     *string                `json:"accountDomain,omitempty"`
	AccountName       *string                `json:"accountName,omitempty"`
	AccountType       *LogonUserAccountType  `json:"accountType,omitempty"`
	FirstSeenDateTime *string                `json:"firstSeenDateTime,omitempty"`
	LastSeenDateTime  *string                `json:"lastSeenDateTime,omitempty"`
	LogonId           *string                `json:"logonId,omitempty"`
	LogonTypes        *[]LogonUserLogonTypes `json:"logonTypes,omitempty"`
	ODataType         *string                `json:"@odata.type,omitempty"`
}
