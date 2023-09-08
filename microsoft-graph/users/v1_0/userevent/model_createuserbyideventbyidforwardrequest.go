package userevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdEventByIdForwardRequest struct {
	Comment      *string      `json:"Comment,omitempty"`
	ToRecipients *[]Recipient `json:"ToRecipients,omitempty"`
}
