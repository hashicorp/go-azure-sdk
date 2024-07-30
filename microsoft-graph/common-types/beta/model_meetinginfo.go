package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingInfo struct {
	AllowConversationWithoutHost *bool   `json:"allowConversationWithoutHost,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
}
