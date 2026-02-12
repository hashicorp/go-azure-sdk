package trigger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChainingTriggerTypeProperties struct {
	DependsOn    []PipelineReference `json:"dependsOn"`
	RunDimension string              `json:"runDimension"`
}
