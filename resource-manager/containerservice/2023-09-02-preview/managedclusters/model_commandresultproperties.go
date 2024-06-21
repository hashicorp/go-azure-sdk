package managedclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandResultProperties struct {
	ExitCode          *int64  `json:"exitCode,omitempty"`
	FinishedAt        *string `json:"finishedAt,omitempty"`
	Logs              *string `json:"logs,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Reason            *string `json:"reason,omitempty"`
	StartedAt         *string `json:"startedAt,omitempty"`
}
