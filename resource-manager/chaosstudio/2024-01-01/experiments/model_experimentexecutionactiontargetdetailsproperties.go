package experiments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExperimentExecutionActionTargetDetailsProperties struct {
	Error               *ExperimentExecutionActionTargetDetailsError `json:"error,omitempty"`
	Status              *string                                      `json:"status,omitempty"`
	Target              *string                                      `json:"target,omitempty"`
	TargetCompletedTime *string                                      `json:"targetCompletedTime,omitempty"`
	TargetFailedTime    *string                                      `json:"targetFailedTime,omitempty"`
}
