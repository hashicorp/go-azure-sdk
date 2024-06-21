package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BmcKeySetPatchProperties struct {
	Expiration *string       `json:"expiration,omitempty"`
	UserList   *[]KeySetUser `json:"userList,omitempty"`
}
