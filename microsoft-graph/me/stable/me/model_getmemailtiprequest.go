package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMeMailTipRequest struct {
	EmailAddresses  *[]string                           `json:"EmailAddresses,omitempty"`
	MailTipsOptions *GetMeMailTipRequestMailTipsOptions `json:"MailTipsOptions,omitempty"`
}
