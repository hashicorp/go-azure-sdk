package workflowversions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowSku struct {
	Name WorkflowSkuName    `json:"name"`
	Plan *ResourceReference `json:"plan,omitempty"`
}
