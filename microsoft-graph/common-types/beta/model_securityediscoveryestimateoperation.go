package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoveryEstimateOperation struct {
	Action             *SecurityEdiscoveryEstimateOperationAction `json:"action,omitempty"`
	CompletedDateTime  *string                                    `json:"completedDateTime,omitempty"`
	CreatedBy          *IdentitySet                               `json:"createdBy,omitempty"`
	CreatedDateTime    *string                                    `json:"createdDateTime,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	IndexedItemCount   *int64                                     `json:"indexedItemCount,omitempty"`
	IndexedItemsSize   *int64                                     `json:"indexedItemsSize,omitempty"`
	MailboxCount       *int64                                     `json:"mailboxCount,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	PercentProgress    *int64                                     `json:"percentProgress,omitempty"`
	ResultInfo         *ResultInfo                                `json:"resultInfo,omitempty"`
	Search             *SecurityEdiscoverySearch                  `json:"search,omitempty"`
	SiteCount          *int64                                     `json:"siteCount,omitempty"`
	Status             *SecurityEdiscoveryEstimateOperationStatus `json:"status,omitempty"`
	UnindexedItemCount *int64                                     `json:"unindexedItemCount,omitempty"`
	UnindexedItemsSize *int64                                     `json:"unindexedItemsSize,omitempty"`
}
