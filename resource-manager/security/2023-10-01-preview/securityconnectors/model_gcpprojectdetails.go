package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GcpProjectDetails struct {
	ProjectId              *string `json:"projectId,omitempty"`
	ProjectName            *string `json:"projectName,omitempty"`
	ProjectNumber          *string `json:"projectNumber,omitempty"`
	WorkloadIdentityPoolId *string `json:"workloadIdentityPoolId,omitempty"`
}
