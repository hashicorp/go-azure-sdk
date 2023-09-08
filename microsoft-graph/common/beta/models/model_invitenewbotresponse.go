package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InviteNewBotResponse struct {
	InviteUri *string `json:"inviteUri,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
