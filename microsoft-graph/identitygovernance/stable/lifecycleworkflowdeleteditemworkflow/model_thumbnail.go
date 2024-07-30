package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Thumbnail struct {
	Content      *string `json:"content,omitempty"`
	Height       *int64  `json:"height,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	SourceItemId *string `json:"sourceItemId,omitempty"`
	Url          *string `json:"url,omitempty"`
	Width        *int64  `json:"width,omitempty"`
}
