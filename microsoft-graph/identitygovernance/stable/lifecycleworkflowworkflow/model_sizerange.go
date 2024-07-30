package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SizeRange struct {
	MaximumSize *int64  `json:"maximumSize,omitempty"`
	MinimumSize *int64  `json:"minimumSize,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
}
