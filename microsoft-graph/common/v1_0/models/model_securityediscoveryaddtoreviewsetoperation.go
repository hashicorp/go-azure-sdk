package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryAddToReviewSetOperation struct {
	Action            *SecurityEdiscoveryAddToReviewSetOperationAction `json:"action,omitempty"`
	CompletedDateTime *string                                          `json:"completedDateTime,omitempty"`
	CreatedBy         *IdentitySet                                     `json:"createdBy,omitempty"`
	CreatedDateTime   *string                                          `json:"createdDateTime,omitempty"`
	Id                *string                                          `json:"id,omitempty"`
	ODataType         *string                                          `json:"@odata.type,omitempty"`
	PercentProgress   *int64                                           `json:"percentProgress,omitempty"`
	ResultInfo        *ResultInfo                                      `json:"resultInfo,omitempty"`
	ReviewSet         *SecurityEdiscoveryReviewSet                     `json:"reviewSet,omitempty"`
	Search            *SecurityEdiscoverySearch                        `json:"search,omitempty"`
	Status            *SecurityEdiscoveryAddToReviewSetOperationStatus `json:"status,omitempty"`
}
