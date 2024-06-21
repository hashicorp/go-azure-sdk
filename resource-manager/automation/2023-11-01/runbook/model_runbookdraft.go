package runbook

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunbookDraft struct {
	CreationTime     *string                      `json:"creationTime,omitempty"`
	DraftContentLink *ContentLink                 `json:"draftContentLink,omitempty"`
	InEdit           *bool                        `json:"inEdit,omitempty"`
	LastModifiedTime *string                      `json:"lastModifiedTime,omitempty"`
	OutputTypes      *[]string                    `json:"outputTypes,omitempty"`
	Parameters       *map[string]RunbookParameter `json:"parameters,omitempty"`
}
