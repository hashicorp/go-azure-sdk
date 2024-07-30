package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallOptions struct {
	HideBotAfterEscalation              *bool   `json:"hideBotAfterEscalation,omitempty"`
	IsContentSharingNotificationEnabled *bool   `json:"isContentSharingNotificationEnabled,omitempty"`
	ODataType                           *string `json:"@odata.type,omitempty"`
}
