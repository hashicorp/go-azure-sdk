package fileshares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPolicy struct {
	ExpiryTime *string `json:"expiryTime,omitempty"`
	Permission *string `json:"permission,omitempty"`
	StartTime  *string `json:"startTime,omitempty"`
}
