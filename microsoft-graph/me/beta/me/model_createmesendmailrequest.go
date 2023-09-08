package me

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeSendMailRequest struct {
	Message         *Message `json:"Message,omitempty"`
	SaveToSentItems *bool    `json:"SaveToSentItems,omitempty"`
}
