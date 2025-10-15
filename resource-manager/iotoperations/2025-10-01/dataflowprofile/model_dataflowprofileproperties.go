package dataflowprofile

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowProfileProperties struct {
	Diagnostics       *ProfileDiagnostics  `json:"diagnostics,omitempty"`
	HealthState       *ResourceHealthState `json:"healthState,omitempty"`
	InstanceCount     *int64               `json:"instanceCount,omitempty"`
	ProvisioningState *ProvisioningState   `json:"provisioningState,omitempty"`
}
