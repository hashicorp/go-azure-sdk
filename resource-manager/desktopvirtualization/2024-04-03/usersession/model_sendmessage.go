package usersession

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendMessage struct {
	MessageBody  *string `json:"messageBody,omitempty"`
	MessageTitle *string `json:"messageTitle,omitempty"`
}
