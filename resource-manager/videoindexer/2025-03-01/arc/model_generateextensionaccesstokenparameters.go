package arc

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateExtensionAccessTokenParameters struct {
	ExtensionId            string         `json:"extensionId"`
	PermissionType         PermissionType `json:"permissionType"`
	Scope                  Scope          `json:"scope"`
	TokenLifetimeInSeconds *int64         `json:"tokenLifetimeInSeconds,omitempty"`
	VideoId                *string        `json:"videoId,omitempty"`
}
