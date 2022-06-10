package deploymentinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentInfoResponse struct {
	DiskCapacity   *string                  `json:"diskCapacity,omitempty"`
	MemoryCapacity *string                  `json:"memoryCapacity,omitempty"`
	Status         *ElasticDeploymentStatus `json:"status,omitempty"`
	Version        *string                  `json:"version,omitempty"`
}
