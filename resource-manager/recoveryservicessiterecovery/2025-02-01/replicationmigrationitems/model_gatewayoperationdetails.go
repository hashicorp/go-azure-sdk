package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewayOperationDetails struct {
	DataStores           *[]string `json:"dataStores,omitempty"`
	HostName             *string   `json:"hostName,omitempty"`
	ProgressPercentage   *int64    `json:"progressPercentage,omitempty"`
	State                *string   `json:"state,omitempty"`
	TimeElapsed          *int64    `json:"timeElapsed,omitempty"`
	TimeRemaining        *int64    `json:"timeRemaining,omitempty"`
	UploadSpeed          *int64    `json:"uploadSpeed,omitempty"`
	VMwareReadThroughput *int64    `json:"vmwareReadThroughput,omitempty"`
}
