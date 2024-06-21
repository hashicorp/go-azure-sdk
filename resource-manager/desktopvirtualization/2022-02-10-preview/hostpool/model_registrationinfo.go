package hostpool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistrationInfo struct {
	ExpirationTime             *string                     `json:"expirationTime,omitempty"`
	RegistrationTokenOperation *RegistrationTokenOperation `json:"registrationTokenOperation,omitempty"`
	Token                      *string                     `json:"token,omitempty"`
}
