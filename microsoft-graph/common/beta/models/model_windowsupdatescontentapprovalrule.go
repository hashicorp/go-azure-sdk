package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesContentApprovalRule struct {
	ContentFilter                 *WindowsUpdatesContentFilter `json:"contentFilter,omitempty"`
	CreatedDateTime               *string                      `json:"createdDateTime,omitempty"`
	DurationBeforeDeploymentStart *string                      `json:"durationBeforeDeploymentStart,omitempty"`
	LastEvaluatedDateTime         *string                      `json:"lastEvaluatedDateTime,omitempty"`
	LastModifiedDateTime          *string                      `json:"lastModifiedDateTime,omitempty"`
	ODataType                     *string                      `json:"@odata.type,omitempty"`
}
