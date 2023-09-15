package netappresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NicInfo struct {
	IPAddress         *string   `json:"ipAddress,omitempty"`
	VolumeResourceIds *[]string `json:"volumeResourceIds,omitempty"`
}
