package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Settings struct {
	HasGraphMailbox *bool   `json:"hasGraphMailbox,omitempty"`
	HasLicense      *bool   `json:"hasLicense,omitempty"`
	HasOptedOut     *bool   `json:"hasOptedOut,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
