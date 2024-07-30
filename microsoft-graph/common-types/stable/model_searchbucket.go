package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBucket struct {
	AggregationFilterToken *string `json:"aggregationFilterToken,omitempty"`
	Count                  *int64  `json:"count,omitempty"`
	Key                    *string `json:"key,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
}
