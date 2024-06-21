package serverendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudTieringLowDiskMode struct {
	LastUpdatedTimestamp *string                       `json:"lastUpdatedTimestamp,omitempty"`
	State                *CloudTieringLowDiskModeState `json:"state,omitempty"`
}
