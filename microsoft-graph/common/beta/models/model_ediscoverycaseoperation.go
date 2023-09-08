package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryCaseOperation struct {
	Action            *EdiscoveryCaseOperationAction `json:"action,omitempty"`
	CompletedDateTime *string                        `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet                   `json:"createdBy,omitempty"`
	CreatedDateTime   *string                        `json:"createdDateTime,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	ODataType         *string                        `json:"@odata.type,omitempty"`
	PercentProgress   *int64                         `json:"percentProgress,omitempty"`
	ResultInfo        *ResultInfo                    `json:"resultInfo,omitempty"`
	Status            *EdiscoveryCaseOperationStatus `json:"status,omitempty"`
}
