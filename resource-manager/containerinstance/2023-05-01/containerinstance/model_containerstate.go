package containerinstance

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContainerState struct {
	DetailStatus *string `json:"detailStatus,omitempty"`
	ExitCode     *int64  `json:"exitCode,omitempty"`
	FinishTime   *string `json:"finishTime,omitempty"`
	StartTime    *string `json:"startTime,omitempty"`
	State        *string `json:"state,omitempty"`
}
