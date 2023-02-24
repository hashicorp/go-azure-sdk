package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Pipeline struct {
	Activities    *[]Activity                        `json:"activities,omitempty"`
	Annotations   *[]interface{}                     `json:"annotations,omitempty"`
	Concurrency   *int64                             `json:"concurrency,omitempty"`
	Description   *string                            `json:"description,omitempty"`
	Folder        *PipelineFolder                    `json:"folder,omitempty"`
	Parameters    *map[string]ParameterSpecification `json:"parameters,omitempty"`
	Policy        *PipelinePolicy                    `json:"policy,omitempty"`
	RunDimensions *map[string]interface{}            `json:"runDimensions,omitempty"`
	Variables     *map[string]VariableSpecification  `json:"variables,omitempty"`
}
