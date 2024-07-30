package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitedUserMessageInfo struct {
	CcRecipients          *[]Recipient `json:"ccRecipients,omitempty"`
	CustomizedMessageBody *string      `json:"customizedMessageBody,omitempty"`
	MessageLanguage       *string      `json:"messageLanguage,omitempty"`
	ODataType             *string      `json:"@odata.type,omitempty"`
}
