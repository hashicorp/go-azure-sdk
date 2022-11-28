package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentSettings struct {
	AddonConfigs                  *map[string]map[string]interface{} `json:"addonConfigs,omitempty"`
	ContainerProbeSettings        *ContainerProbeSettings            `json:"containerProbeSettings"`
	EnvironmentVariables          *map[string]string                 `json:"environmentVariables,omitempty"`
	LivenessProbe                 *Probe                             `json:"livenessProbe"`
	ReadinessProbe                *Probe                             `json:"readinessProbe"`
	ResourceRequests              *ResourceRequests                  `json:"resourceRequests"`
	StartupProbe                  *Probe                             `json:"startupProbe"`
	TerminationGracePeriodSeconds *int64                             `json:"terminationGracePeriodSeconds,omitempty"`
}
