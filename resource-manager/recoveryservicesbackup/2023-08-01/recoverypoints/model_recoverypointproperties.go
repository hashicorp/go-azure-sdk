package recoverypoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecoveryPointProperties struct {
	ExpiryTime    *string `json:"expiryTime,omitempty"`
	IsSoftDeleted *bool   `json:"isSoftDeleted,omitempty"`
	RuleName      *string `json:"ruleName,omitempty"`
}
