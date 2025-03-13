package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstructionStep struct {
	Description  *string                   `json:"description,omitempty"`
	InnerSteps   *[]InstructionStep        `json:"innerSteps,omitempty"`
	Instructions *[]InstructionStepDetails `json:"instructions,omitempty"`
	Title        *string                   `json:"title,omitempty"`
}
