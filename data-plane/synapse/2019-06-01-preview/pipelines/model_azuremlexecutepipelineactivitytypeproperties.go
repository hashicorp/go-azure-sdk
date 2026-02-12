package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureMLExecutePipelineActivityTypeProperties struct {
	ContinueOnStepFailure *interface{} `json:"continueOnStepFailure,omitempty"`
	ExperimentName        *interface{} `json:"experimentName,omitempty"`
	MlParentRunId         *interface{} `json:"mlParentRunId,omitempty"`
	MlPipelineId          interface{}  `json:"mlPipelineId"`
	MlPipelineParameters  *interface{} `json:"mlPipelineParameters,omitempty"`
}
