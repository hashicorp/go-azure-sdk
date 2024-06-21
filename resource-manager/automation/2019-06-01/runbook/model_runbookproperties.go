package runbook

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunbookProperties struct {
	CreationTime       *string                      `json:"creationTime,omitempty"`
	Description        *string                      `json:"description,omitempty"`
	Draft              *RunbookDraft                `json:"draft,omitempty"`
	JobCount           *int64                       `json:"jobCount,omitempty"`
	LastModifiedBy     *string                      `json:"lastModifiedBy,omitempty"`
	LastModifiedTime   *string                      `json:"lastModifiedTime,omitempty"`
	LogActivityTrace   *int64                       `json:"logActivityTrace,omitempty"`
	LogProgress        *bool                        `json:"logProgress,omitempty"`
	LogVerbose         *bool                        `json:"logVerbose,omitempty"`
	OutputTypes        *[]string                    `json:"outputTypes,omitempty"`
	Parameters         *map[string]RunbookParameter `json:"parameters,omitempty"`
	ProvisioningState  *RunbookProvisioningState    `json:"provisioningState,omitempty"`
	PublishContentLink *ContentLink                 `json:"publishContentLink,omitempty"`
	RunbookType        *RunbookTypeEnum             `json:"runbookType,omitempty"`
	State              *RunbookState                `json:"state,omitempty"`
}
