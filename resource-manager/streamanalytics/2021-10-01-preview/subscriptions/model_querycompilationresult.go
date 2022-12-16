package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryCompilationResult struct {
	Errors    *[]QueryCompilationError `json:"errors,omitempty"`
	Functions *[]string                `json:"functions,omitempty"`
	Inputs    *[]string                `json:"inputs,omitempty"`
	Outputs   *[]string                `json:"outputs,omitempty"`
	Warnings  *[]string                `json:"warnings,omitempty"`
}
