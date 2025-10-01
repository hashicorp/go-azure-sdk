package bms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityBasedRestoreDetails struct {
	ObjectType             *string `json:"objectType,omitempty"`
	TargetStorageAccountId *string `json:"targetStorageAccountId,omitempty"`
}
