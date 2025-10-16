package actions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionResponseProperties struct {
	LogicAppResourceId string  `json:"logicAppResourceId"`
	WorkflowId         *string `json:"workflowId,omitempty"`
}
