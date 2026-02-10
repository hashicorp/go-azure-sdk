package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExitConditions struct {
	Default            *ExitOptions            `json:"default,omitempty"`
	ExitCodeRanges     *[]ExitCodeRangeMapping `json:"exitCodeRanges,omitempty"`
	ExitCodes          *[]ExitCodeMapping      `json:"exitCodes,omitempty"`
	FileUploadError    *ExitOptions            `json:"fileUploadError,omitempty"`
	PreProcessingError *ExitOptions            `json:"preProcessingError,omitempty"`
}
