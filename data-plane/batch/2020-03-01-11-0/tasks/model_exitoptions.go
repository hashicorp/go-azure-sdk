package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExitOptions struct {
	DependencyAction *DependencyAction `json:"dependencyAction,omitempty"`
	JobAction        *JobAction        `json:"jobAction,omitempty"`
}
