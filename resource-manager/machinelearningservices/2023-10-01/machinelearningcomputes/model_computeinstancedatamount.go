package machinelearningcomputes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeInstanceDataMount struct {
	CreatedBy   *string      `json:"createdBy,omitempty"`
	Error       *string      `json:"error,omitempty"`
	MountAction *MountAction `json:"mountAction,omitempty"`
	MountName   *string      `json:"mountName,omitempty"`
	MountPath   *string      `json:"mountPath,omitempty"`
	MountState  *MountState  `json:"mountState,omitempty"`
	MountedOn   *string      `json:"mountedOn,omitempty"`
	Source      *string      `json:"source,omitempty"`
	SourceType  *SourceType  `json:"sourceType,omitempty"`
}
