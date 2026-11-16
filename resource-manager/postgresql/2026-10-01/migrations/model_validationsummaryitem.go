package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationSummaryItem struct {
	Messages *[]ValidationMessage `json:"messages,omitempty"`
	State    *ValidationState     `json:"state,omitempty"`
	Type     *string              `json:"type,omitempty"`
}
