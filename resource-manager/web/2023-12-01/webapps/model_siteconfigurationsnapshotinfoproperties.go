package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteConfigurationSnapshotInfoProperties struct {
	SnapshotId *int64  `json:"snapshotId,omitempty"`
	Time       *string `json:"time,omitempty"`
}
