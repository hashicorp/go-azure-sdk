package connector

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidateResult struct {
	AuthType              *AuthType               `json:"authType,omitempty"`
	IsConnectionAvailable *bool                   `json:"isConnectionAvailable,omitempty"`
	LinkerName            *string                 `json:"linkerName,omitempty"`
	ReportEndTimeUtc      *string                 `json:"reportEndTimeUtc,omitempty"`
	ReportStartTimeUtc    *string                 `json:"reportStartTimeUtc,omitempty"`
	SourceId              *string                 `json:"sourceId,omitempty"`
	TargetId              *string                 `json:"targetId,omitempty"`
	ValidationDetail      *[]ValidationResultItem `json:"validationDetail,omitempty"`
}
