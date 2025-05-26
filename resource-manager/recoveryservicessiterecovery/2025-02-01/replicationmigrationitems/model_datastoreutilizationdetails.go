package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataStoreUtilizationDetails struct {
	DataStoreName           *string `json:"dataStoreName,omitempty"`
	TotalSnapshotsCreated   *int64  `json:"totalSnapshotsCreated,omitempty"`
	TotalSnapshotsSupported *int64  `json:"totalSnapshotsSupported,omitempty"`
}
