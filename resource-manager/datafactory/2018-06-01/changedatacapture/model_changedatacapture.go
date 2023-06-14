package changedatacapture

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCapture struct {
	AllowVNetOverride     *bool                         `json:"allowVNetOverride,omitempty"`
	Description           *string                       `json:"description,omitempty"`
	Folder                *ChangeDataCaptureFolder      `json:"folder,omitempty"`
	Policy                MapperPolicy                  `json:"policy"`
	SourceConnectionsInfo []MapperSourceConnectionsInfo `json:"sourceConnectionsInfo"`
	Status                *string                       `json:"status,omitempty"`
	TargetConnectionsInfo []MapperTargetConnectionsInfo `json:"targetConnectionsInfo"`
}
