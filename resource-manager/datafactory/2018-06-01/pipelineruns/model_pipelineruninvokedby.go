package pipelineruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunInvokedBy struct {
	Id            *string `json:"id,omitempty"`
	InvokedByType *string `json:"invokedByType,omitempty"`
	Name          *string `json:"name,omitempty"`
	PipelineName  *string `json:"pipelineName,omitempty"`
	PipelineRunId *string `json:"pipelineRunId,omitempty"`
}
