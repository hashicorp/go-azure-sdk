package sqlpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetadataSyncConfigProperties struct {
	Enabled               *bool  `json:"enabled,omitempty"`
	SyncIntervalInMinutes *int64 `json:"syncIntervalInMinutes,omitempty"`
}
