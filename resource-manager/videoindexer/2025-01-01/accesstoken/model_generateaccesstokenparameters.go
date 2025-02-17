package accesstoken

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GenerateAccessTokenParameters struct {
	PermissionType PermissionType `json:"permissionType"`
	ProjectId      *string        `json:"projectId,omitempty"`
	Scope          Scope          `json:"scope"`
	VideoId        *string        `json:"videoId,omitempty"`
}
