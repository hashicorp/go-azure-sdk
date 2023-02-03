package sites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareSiteUsage struct {
	MachineCount      *int64 `json:"machineCount,omitempty"`
	RunAsAccountCount *int64 `json:"runAsAccountCount,omitempty"`
	VCenterCount      *int64 `json:"vCenterCount,omitempty"`
}
