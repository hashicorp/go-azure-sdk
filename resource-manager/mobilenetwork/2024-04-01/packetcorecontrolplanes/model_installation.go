package packetcorecontrolplanes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Installation struct {
	DesiredState      *DesiredInstallationState `json:"desiredState,omitempty"`
	Operation         *AsyncOperationId         `json:"operation,omitempty"`
	Reasons           *[]InstallationReason     `json:"reasons,omitempty"`
	ReinstallRequired *ReinstallRequired        `json:"reinstallRequired,omitempty"`
	State             *InstallationState        `json:"state,omitempty"`
}
