package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MountPointMap struct {
	MountPoint *string    `json:"mountPoint,omitempty"`
	MountType  *MountType `json:"mountType,omitempty"`
	RoleId     *string    `json:"roleId,omitempty"`
	RoleType   *RoleTypes `json:"roleType,omitempty"`
	ShareId    string     `json:"shareId"`
}
