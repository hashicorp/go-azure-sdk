package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskProfile struct {
	AppConfiguration          *WindowsKioskAppConfiguration `json:"appConfiguration,omitempty"`
	ODataType                 *string                       `json:"@odata.type,omitempty"`
	ProfileId                 *string                       `json:"profileId,omitempty"`
	ProfileName               *string                       `json:"profileName,omitempty"`
	UserAccountsConfiguration *[]WindowsKioskUser           `json:"userAccountsConfiguration,omitempty"`
}
