package sessionhostmanagements

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostPoolUpdateConfigurationProperties struct {
	DeleteOriginalVM   *bool   `json:"deleteOriginalVm,omitempty"`
	LogOffDelayMinutes int64   `json:"logOffDelayMinutes"`
	LogOffMessage      *string `json:"logOffMessage,omitempty"`
	MaxVMsRemoved      int64   `json:"maxVmsRemoved"`
}
