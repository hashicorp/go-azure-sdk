package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesExpediteSettings struct {
	IsExpedited     *bool   `json:"isExpedited,omitempty"`
	IsReadinessTest *bool   `json:"isReadinessTest,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
