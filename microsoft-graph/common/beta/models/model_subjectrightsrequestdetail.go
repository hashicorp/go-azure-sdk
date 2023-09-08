package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestDetail struct {
	ExcludedItemCount  *int64          `json:"excludedItemCount,omitempty"`
	InsightCounts      *[]KeyValuePair `json:"insightCounts,omitempty"`
	ItemCount          *int64          `json:"itemCount,omitempty"`
	ItemNeedReview     *int64          `json:"itemNeedReview,omitempty"`
	ODataType          *string         `json:"@odata.type,omitempty"`
	ProductItemCounts  *[]KeyValuePair `json:"productItemCounts,omitempty"`
	SignedOffItemCount *int64          `json:"signedOffItemCount,omitempty"`
	TotalItemSize      *int64          `json:"totalItemSize,omitempty"`
}
