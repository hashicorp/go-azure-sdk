package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TelecomExpenseManagementPartner struct {
	AppAuthorized          *bool   `json:"appAuthorized,omitempty"`
	DisplayName            *string `json:"displayName,omitempty"`
	Enabled                *bool   `json:"enabled,omitempty"`
	Id                     *string `json:"id,omitempty"`
	LastConnectionDateTime *string `json:"lastConnectionDateTime,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	Url                    *string `json:"url,omitempty"`
}
