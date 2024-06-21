package pipelineruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRunResponse struct {
	CatalogDigest           *string                         `json:"catalogDigest,omitempty"`
	FinishTime              *string                         `json:"finishTime,omitempty"`
	ImportedArtifacts       *[]string                       `json:"importedArtifacts,omitempty"`
	PipelineRunErrorMessage *string                         `json:"pipelineRunErrorMessage,omitempty"`
	Progress                *ProgressProperties             `json:"progress,omitempty"`
	Source                  *ImportPipelineSourceProperties `json:"source,omitempty"`
	StartTime               *string                         `json:"startTime,omitempty"`
	Status                  *string                         `json:"status,omitempty"`
	Target                  *ExportPipelineTargetProperties `json:"target,omitempty"`
	Trigger                 *PipelineTriggerDescriptor      `json:"trigger,omitempty"`
}
