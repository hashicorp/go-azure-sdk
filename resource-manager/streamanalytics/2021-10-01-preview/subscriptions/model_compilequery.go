package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompileQuery struct {
	CompatibilityLevel *CompatibilityLevel `json:"compatibilityLevel,omitempty"`
	Functions          *[]QueryFunction    `json:"functions,omitempty"`
	Inputs             *[]QueryInput       `json:"inputs,omitempty"`
	JobType            JobType             `json:"jobType"`
	Query              string              `json:"query"`
}
