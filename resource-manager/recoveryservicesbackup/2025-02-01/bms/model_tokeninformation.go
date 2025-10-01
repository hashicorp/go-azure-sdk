package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TokenInformation struct {
	ExpiryTimeInUtcTicks *int64  `json:"expiryTimeInUtcTicks,omitempty"`
	SecurityPIN          *string `json:"securityPIN,omitempty"`
	Token                *string `json:"token,omitempty"`
}
