package user

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserMailTipRequest struct {
	EmailAddresses  *[]string                             `json:"EmailAddresses,omitempty"`
	MailTipsOptions *GetUserMailTipRequestMailTipsOptions `json:"MailTipsOptions,omitempty"`
}
