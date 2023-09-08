package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewInstanceDecisionItemCollectionResponse struct {
	ODataCount    *int64                              `json:"@odata.count,omitempty"`
	ODataNextLink *string                             `json:"@odata.nextLink,omitempty"`
	Value         *[]AccessReviewInstanceDecisionItem `json:"value,omitempty"`
}
