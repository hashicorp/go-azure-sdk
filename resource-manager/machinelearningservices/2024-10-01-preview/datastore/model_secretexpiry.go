package datastore

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecretExpiry struct {
	ExpirableSecret  *bool  `json:"expirableSecret,omitempty"`
	ExpireAfterHours *int64 `json:"expireAfterHours,omitempty"`
}
