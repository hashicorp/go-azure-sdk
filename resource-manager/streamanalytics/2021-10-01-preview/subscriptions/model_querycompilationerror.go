package subscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryCompilationError struct {
	EndColumn   *int64  `json:"endColumn,omitempty"`
	EndLine     *int64  `json:"endLine,omitempty"`
	IsGlobal    *bool   `json:"isGlobal,omitempty"`
	Message     *string `json:"message,omitempty"`
	StartColumn *int64  `json:"startColumn,omitempty"`
	StartLine   *int64  `json:"startLine,omitempty"`
}
