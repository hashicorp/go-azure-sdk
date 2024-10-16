package appplatform

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobExecutionTemplate struct {
	Args                 *[]string            `json:"args,omitempty"`
	EnvironmentVariables *[]EnvVar            `json:"environmentVariables,omitempty"`
	ResourceRequests     *JobResourceRequests `json:"resourceRequests,omitempty"`
}
