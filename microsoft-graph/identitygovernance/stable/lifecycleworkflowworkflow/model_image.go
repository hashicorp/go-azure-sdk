package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Image struct {
	Height    *int64  `json:"height,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Width     *int64  `json:"width,omitempty"`
}
