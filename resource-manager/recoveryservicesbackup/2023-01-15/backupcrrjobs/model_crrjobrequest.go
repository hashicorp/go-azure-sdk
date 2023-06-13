package backupcrrjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrrJobRequest struct {
	JobName    *string `json:"jobName,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}
