package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAggregation struct {
	Buckets   *[]SearchBucket `json:"buckets,omitempty"`
	Field     *string         `json:"field,omitempty"`
	ODataType *string         `json:"@odata.type,omitempty"`
}
