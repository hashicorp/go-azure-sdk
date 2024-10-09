package trigger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerPipelineReference struct {
	Parameters        *map[string]interface{} `json:"parameters,omitempty"`
	PipelineReference *PipelineReference      `json:"pipelineReference,omitempty"`
}
