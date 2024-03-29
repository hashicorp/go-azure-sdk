package actions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecallActionParameters struct {
	Pattern    *string `json:"pattern,omitempty"`
	RecallPath *string `json:"recallPath,omitempty"`
}
