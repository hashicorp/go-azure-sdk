package machinelearningcomputes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeInstanceProperties struct {
	ApplicationSharingPolicy         *ApplicationSharingPolicy             `json:"applicationSharingPolicy,omitempty"`
	Applications                     *[]ComputeInstanceApplication         `json:"applications,omitempty"`
	ComputeInstanceAuthorizationType *ComputeInstanceAuthorizationType     `json:"computeInstanceAuthorizationType,omitempty"`
	ConnectivityEndpoints            *ComputeInstanceConnectivityEndpoints `json:"connectivityEndpoints,omitempty"`
	CreatedBy                        *ComputeInstanceCreatedBy             `json:"createdBy,omitempty"`
	Errors                           *[]ErrorResponse                      `json:"errors,omitempty"`
	LastOperation                    *ComputeInstanceLastOperation         `json:"lastOperation,omitempty"`
	PersonalComputeInstanceSettings  *PersonalComputeInstanceSettings      `json:"personalComputeInstanceSettings,omitempty"`
	SetupScripts                     *SetupScripts                         `json:"setupScripts,omitempty"`
	SshSettings                      *ComputeInstanceSshSettings           `json:"sshSettings,omitempty"`
	State                            *ComputeInstanceState                 `json:"state,omitempty"`
	Subnet                           *ResourceId                           `json:"subnet,omitempty"`
	VmSize                           *string                               `json:"vmSize,omitempty"`
}
