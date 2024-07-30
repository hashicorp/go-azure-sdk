package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BulkDriverActionResult struct {
	FailedDriverIds     *[]string `json:"failedDriverIds,omitempty"`
	NotFoundDriverIds   *[]string `json:"notFoundDriverIds,omitempty"`
	ODataType           *string   `json:"@odata.type,omitempty"`
	SuccessfulDriverIds *[]string `json:"successfulDriverIds,omitempty"`
}
