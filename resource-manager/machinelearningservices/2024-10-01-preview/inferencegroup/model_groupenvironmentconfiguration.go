package inferencegroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupEnvironmentConfiguration struct {
	EnvironmentId        *string                     `json:"environmentId,omitempty"`
	EnvironmentVariables *[]StringStringKeyValuePair `json:"environmentVariables,omitempty"`
	LivenessProbe        *ProbeSettings              `json:"livenessProbe,omitempty"`
	ReadinessProbe       *ProbeSettings              `json:"readinessProbe,omitempty"`
	StartupProbe         *ProbeSettings              `json:"startupProbe,omitempty"`
}
