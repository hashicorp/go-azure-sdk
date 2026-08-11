package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CacheMountTargetProperties struct {
	IPAddress     *string `json:"ipAddress,omitempty"`
	MountTargetId *string `json:"mountTargetId,omitempty"`
	SmbServerFqdn *string `json:"smbServerFqdn,omitempty"`
}
