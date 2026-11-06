package sandboxgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SandboxGroupProperties struct {
	DefaultDomain     *string                        `json:"defaultDomain,omitempty"`
	EnvironmentId     *string                        `json:"environmentId,omitempty"`
	ProvisioningState *SandboxGroupProvisioningState `json:"provisioningState,omitempty"`
}
